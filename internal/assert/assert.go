// Package assert has various functions useful for testing rosed.
package assert

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

// Asserter performs various tasks and fails a provided testing.T when a test
// fails. The zero-value of an Asserter is not suitable for use, and should be
// created with a call to New()
type Asserter struct {
	t *testing.T

	// NonFatal sets how a test fails. When true, t.Errorf is called on
	// assertion failure instead of t.Fatalf
	NonFatal bool

	// VarName is the what the variable is called in test error messages. It
	// can be modified between assertions, or by calling Var() on the Asserter.
	//
	// If VarName is set to the empty string, "value" will be used as the name
	// of the value being tested.
	VarName string

	// VarNamePrefix is the string that is printed before every varName. It can
	// be used to assign a prefix when a long series of subpaths must be tested.
	VarNamePrefix string

	// Format is a custom format function used on failure. If set to a non-empty
	// value, that output is used as the error message on test failure instead
	// of the default.
	//
	// Variables:
	// varName - the name of the variable being tested. It will include any
	// prefix specified on the Asserter at the time that Format is called.
	// expected - what the caller expected a value to be.
	// actual - what the test actually resulted in.
	//
	// If set to nil, default behavior for formatting failure messages is
	// performed, the specifics of which vary depending on the assertion.
	Format func(varName string, expected interface{}, actual interface{}) string
}

// New creates a new Asserter that fails the provided testing.T when an
// assertion fails.
func New(t *testing.T) Asserter {
	return Asserter{t: t}
}

// Var sets the variable name for future tests. Can be chained as Var().Equal,
// etc. This is shorthand for just setting VarName in the Asserter.
func (a *Asserter) Var(name string) *Asserter {
	a.VarName = name
	return a
}

// Equal checks that the actual and expected values are equal.
func (a Asserter) Equal(expected, actual interface{}) {
	argsEqual, err := checkEqual(expected, actual, nil)
	if err != nil {
		a.fail("comparison for %s failed; expected and actual are not comparable types", a.fullVar(), skipArg{expected}, skipArg{actual})
	}

	if argsEqual {
		eVerb := fmtVerbForArg(expected)
		aVerb := fmtVerbForArg(actual)
		a.fail("expected %s to be "+eVerb+" but was "+aVerb, a.fullVar(), expected, actual)
	}
}

// DeepEqual checks that the actual and expected values are deeply-equal by
// calling reflect.DeepEqual on the two arguments.
func (a Asserter) DeepEqual(expected, actual interface{}) {
	if reflect.DeepEqual(actual, expected) {
		eVerb := fmtVerbForArg(expected)
		aVerb := fmtVerbForArg(actual)
		a.fail("expected %s to be "+eVerb+" but was "+aVerb, a.fullVar(), expected, actual)
	}
}

// EqualContentsString checks that the actual and expected values are either
// both nil or both point to the same contents.
func (a Asserter) EqualContentsString(expected, actual *string) {
	if expected == nil && actual != nil {
		a.fail("expected %s to be <nil> but was %q", a.fullVar(), skipArg{expected}, *actual)
	}

	if expected != nil && actual == nil {
		a.fail("expected %s to be %q but was <nil>", a.fullVar(), *expected, skipArg{actual})
	}

	// at this point both expected and actual are either both nil or both non-nil.
	if expected != nil {
		if *expected != *actual {
			a.fail("expected %s to be %q but was %q", a.fullVar(), *expected, *actual)
		}
	}
}

// EqualSlices checks whether expected and actual are two slice-like (slice
// or array) objects with equal size, same type of element, and equal elements.
func (a Asserter) EqualSlices(expected, actual interface{}) {
	a.EqualSlicesFunc(expected, actual, nil)
}

// EqualSlicesFunc checks whether expected and actual are two slice-like (slice
// or array) objects with equal size, same type of element, and equal elements.
//
// The elements are compared using the provided comp function which returns
// whether the two elements passed to it are equal.
func (a Asserter) EqualSlicesFunc(expected, actual interface{}, elemComp func(expected interface{}, actual interface{}) bool) {
	eType := reflect.TypeOf(expected)
	aType := reflect.TypeOf(actual)

	// assert that both are slicey
	if eType.Kind() != reflect.Slice && eType.Kind() != reflect.Array {
		a.fail("%s: expected is not a slice", a.fullVar(), skipArg{expected}, skipArg{actual})
	}
	if aType.Kind() != reflect.Slice && aType.Kind() != reflect.Array {
		a.fail("expected %s to be %v but actual value was not a slice", a.fullVar(), expected, skipArg{actual})
	}

	// assert that both are of the same type
	if eType.Elem() != aType.Elem() {
		a.fail("expected %s to have type %q but was %q", a.fullVar(), eType.Elem().Name(), aType.Elem().Name)
	}

	var eVal, aVal = reflect.ValueOf(expected), reflect.ValueOf(actual)
	// Do nil check
	aIsNil := aType.Kind() == reflect.Slice && aVal.IsNil()
	eIsNil := eType.Kind() == reflect.Slice && eVal.IsNil()
	if aIsNil && !eIsNil {
		a.fail("expected %s to be %v but was a nil slice", a.fullVar(), expected, skipArg{actual})
	}
	if !aIsNil && eIsNil {
		a.fail("expected %s to be a nil slice but was %v", a.fullVar(), skipArg{expected}, actual)
	}
	if aIsNil && eIsNil {
		// nothing else to do, they are both nil slices of the same type so they
		// are equal
		return
	}

	// now we know for sure that both actual and expected are non-nil slicy
	// vals of the same element.

	if eVal.Len() != aVal.Len() {
		a.fail("expected %s to have len of %d but was %d", a.fullVar(), eVal.Len(), aVal.Len())
	}
	for i := 0; i < eVal.Len(); i++ {
		eItem := eVal.Index(i).Interface()
		aItem := aVal.Index(i).Interface()

		varName := fmt.Sprintf("%s[%d]", a.fullVar(), i)
		eq, err := checkEqual(eItem, aItem, elemComp)
		if err != nil {
			a.fail("comparison for %s failed; expected and actual are not comparable types", varName, skipArg{expected}, skipArg{actual})
		}

		if !eq {
			eVerb := fmtVerbForArg(expected)
			aVerb := fmtVerbForArg(actual)
			a.fail("expected %s to be "+eVerb+" but was "+aVerb, varName, eVal, aVal)
		}
	}
}

// formats s with a and adds caller info if it is available.
//
// must always be called only from a function in Asserter
// must always be called only from a top-level exported function in this module.
func addCallerInfo(s string, b ...interface{}) string {
	// try getting caller info first bc everything could break
	// unlikely we are at a call depth of 10, but if we are we can
	// always get more stack frames
	framePCs := make([]uintptr, 10)
	n := runtime.Callers(1, framePCs)
	if n == 0 {
		// cant get any call frames for some reason; just default to not adding
		// the info
		return fmt.Sprintf(s, b...)
	}
	frames := runtime.CallersFrames(framePCs[:n])
	for f, more := frames.Next(); more; f, more = frames.Next() {
		fmt.Printf("FUNC: %v\n", f)
	}

	full := fmt.Sprintf(s, b...)
	_, callingFile, callingLine, ok := runtime.Caller(2)
	if ok {
		return fmt.Sprintf("\n%s:%d: %s", callingFile, callingLine, full)
	}
	return full
}

// gets the format verb to use for a particular arg passed in to an Assert
// function.
func fmtVerbForArg(value interface{}) string {
	if reflect.ValueOf(value).Kind() == reflect.String {
		return "%q"
	}
	return "%v"
}

// varName is full varName as it will be shown.
//
// expected or actual can be set to skipArg in which case they will not be
// included to the arguments to format iff the user has not defined a custom
// Format func. If the user HAS defined a custom format func, expected and
// actual args that have been set to skipArg are converted to nil and the user
// is expected to be able to deal with further formatting.
func (a Asserter) fail(format string, varName string, expected interface{}, actual interface{}, moreArgs ...interface{}) {
	var failureMsg string
	if a.Format != nil {
		if exp, isSkip := expected.(skipArg); isSkip {
			expected = exp.arg
		}
		if act, isSkip := actual.(skipArg); isSkip {
			actual = act.arg
		}
		failureMsg = addCallerInfo(a.Format(varName, expected, actual))
	} else {
		// build up args only out of non-skip elements
		fmtArgs := []interface{}{varName}
		if _, isSkip := expected.(skipArg); !isSkip {
			fmtArgs = append(fmtArgs, expected)
		}
		if _, isSkip := actual.(skipArg); !isSkip {
			fmtArgs = append(fmtArgs, actual)
		}
		for _, m := range moreArgs {
			if _, isSkip := m.(skipArg); !isSkip {
				fmtArgs = append(fmtArgs, m)
			}
		}
		failureMsg = addCallerInfo(fmt.Sprintf(format, fmtArgs))
	}

	if a.NonFatal {
		a.t.Errorf(failureMsg)
	} else {
		a.t.Fatalf(failureMsg)
	}
}

// returns the currently defined VarName or else "value" if it is undefined. If
// VarNamePrefix is defined and VarName is also defined, VarNamePrefix is
// prepended to the result.
func (a Asserter) fullVar() string {
	var name string
	if a.VarName != "" {
		name = a.VarNamePrefix + a.VarName
	} else if a.VarNamePrefix != "" {
		name = a.VarNamePrefix + "'s value"
	} else {
		name = "value"
	}
	return name
}

// skipArg is used in calls to fail() to indicate that an argument should be
// skipped if not being passed to a custom Format function.
type skipArg struct {
	arg interface{}
}
