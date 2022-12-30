package rosed

import (
	"fmt"

	"github.com/dekarrin/rosed/internal/gem"
)

const (
	// DefaultIndentString is the default string used for a single horizontal
	// indent.
	DefaultIndentString = "\t"

	// DefaultLineSeparator is the default line separator sequence.
	DefaultLineSeparator = "\n"

	// DefaultParagraphSeparator is the default sequence that separates
	// paragraphs.
	DefaultParagraphSeparator = "\n\n"

	// DefaultTableCharSet is the default characters used to draw table borders.
	DefaultTableCharSet = "+|-"
)

// Options control the behavior of an [Editor]. The zero-value is an Options
// with all members set to defaults.
//
// IndentStr, LineSeparator, ParagraphSeparator, and TableChars have special
// behavior if not set manually. In a zero-valued Options, each one will be the
// empty string. When interpreting the options in the course of performing an
// operation, functions that use those values will treat an empty string as
// [DefaultIndentString], [DefaultLineSeparator], [DefaultParagraphSeparator],
// or [DefaultTableChars] respectively.
type Options struct {
	// IndentStr is the string that is used for a single horizontal indent. If
	// this is set to "", it will be interpreted as though it were set to
	// DefaultIndentString.
	IndentStr string

	// LineSeparator is the string that the Editor considers to signify the
	// end of a line. If this is set to "", it will be interpreted as though it
	// were set to DefaultLineSeparator.
	LineSeparator string

	// NoTrailingLineSeparators is whether the Editor considers lines to not end
	// with the separator, and thus would assume that a properly formatted line
	// does not include a line separator at the end even if it is the last line.
	//
	// This has a variety of effects depending on the function that is being
	// called on an Editor; functions that are affected will call it out in
	// their documentation.
	//
	// If this is set to false (the default), line separator chars are assumed
	// to signify the end of the line.
	NoTrailingLineSeparators bool

	// ParagraphSeparator is the sequence that is considered to separate
	// paragraphs in the text. Paragraphs are considered to have separators
	// rather than terminators; i.e. this sequence does not occur at the start
	// of the first paragraph or at the end of the final paragraph. It may or
	// may not include LineSeparator as a substring; every mode of operation of
	// Editor is intended to transparantly handle this case.
	//
	// If this is set to "", it will be interpreted as though it were set to
	// DefaultParagraphSeparator.
	ParagraphSeparator string

	// PreserveParagraphs says whether operations that adjust separator
	// characters (such as wrap) should preserve paragraphs and their
	// separators. If not set, certain operations may modify paragraph
	// separators.
	PreserveParagraphs bool

	// JustifyLastLine gives whether text justify operations should apply to
	// the last line of a block of text (or paragraph if PreserveParagraphs is
	// set to true and is respected). Conventionally, justifications are not
	// applied to the last line of a block of text and this is the default
	// behavior.
	JustifyLastLine bool

	// TableBorders sets whether created tables draw cell and table borders on
	// them.
	TableBorders bool

	// TableHeaders sets whether table creation functions should consider the
	// first row of data to be data headers. If so, they will be layed out
	// separately from the other data and separated by a horizontal rule.
	TableHeaders bool

	// TableCharSet is the set of characters used to draw tables. It can be up
	// to 3 characters long. The first char is used to draw table corners and
	// line crossings, the second char is used to draw vertical lines, and the
	// third char is used to draw horizontal lines.
	//
	// Any characters beyond the first three are ignored.
	//
	// If this is set to "", it will be interpreted as though it were set to
	// DefaultTableCharSet. If it is set to less than three characters, the
	// missing characters up to 3 will be filled in from the remaining chars in
	// DefaultTableCharSet; e.g. setting TableCharSet to "#" will result in an
	// interpreted TableCharSet of "#|-".
	TableCharSet string
}

// String gets the string representation of the Options.
func (opts Options) String() string {
	fmtStr := "Options{ParagraphSeparator: %q,"
	fmtStr += " LineSeparator: %q,"
	fmtStr += " IndentStr: %q,"
	fmtStr += " NoTrailingLineSeparators: %v,"
	fmtStr += " PreserveParagraphs: %v,"
	fmtStr += " JustifyLastLine: %v,"
	fmtStr += " TableBorders: %v,"
	fmtStr += " TableHeaders: %v,"
	fmtStr += " TableCharSet: %q}"
	return fmt.Sprintf(
		fmtStr, opts.ParagraphSeparator, opts.LineSeparator, opts.IndentStr,
		opts.NoTrailingLineSeparators, opts.PreserveParagraphs,
		opts.JustifyLastLine, opts.TableBorders, opts.TableHeaders,
		opts.TableCharSet,
	)
}

// WithDefaults returns a copy of the options with all blank members filled with
// their defaults. Internally, this function is used on user-provided Options
// structs in order to get ready-to-use copies.
//
// This function does not modify the Options it is called on.
func (opts Options) WithDefaults() Options {
	if opts.LineSeparator == "" {
		opts.LineSeparator = DefaultLineSeparator
	}
	if opts.IndentStr == "" {
		opts.IndentStr = DefaultIndentString
	}
	if opts.ParagraphSeparator == "" {
		opts.ParagraphSeparator = DefaultParagraphSeparator
	}

	gemTableCharSet := gem.New(opts.TableCharSet)
	gemDefaultTableCharSet := gem.New(DefaultTableCharSet)
	if gemTableCharSet.Len() != gemDefaultTableCharSet.Len() {
		if gemTableCharSet.Len() < gemDefaultTableCharSet.Len() {
			numNeeded := gemDefaultTableCharSet.Len() - gemTableCharSet.Len()

			// need 1, get (2, 3)
			// need 2, get (1, 3)
			// need 3, get (0, 3)

			end := gemDefaultTableCharSet.Len()
			start := end - numNeeded

			gemTableCharSet = gemTableCharSet.Add(gemDefaultTableCharSet.Sub(start, end))
		} else {
			gemTableCharSet = gemTableCharSet.Sub(0, gemDefaultTableCharSet.Len())
		}
		opts.TableCharSet = gemTableCharSet.String()
	}

	return opts
}

// WithIndentStr returns a new Options identical to this one but with IndentStr
// set to str. If str is the empty string, the indent str is interpreted as
// [DefaultIndentString].
//
// This function does not modify the Options it is called on.
func (opts Options) WithIndentStr(str string) Options {
	opts.IndentStr = str
	return opts
}

// WithJustifyLastLine returns a new Options identical to this one but with
// JustifyLastLine set to justifyLastLine.
//
// This function does not modify the Options it is called on.
func (opts Options) WithJustifyLastLine(justifyLastLine bool) Options {
	opts.JustifyLastLine = justifyLastLine
	return opts
}

// WithLineSeparator returns a new Options identical to this one but with the
// LineSeparator member set to sep. If sep is the empty string, the line
// separator is interpreted as [DefaultLineSeparator].
//
// This function does not modify the Options it is called on.
func (opts Options) WithLineSeparator(sep string) Options {
	opts.LineSeparator = sep
	return opts
}

// WithNoTrailingLineSeparators returns a new Options identical to this one but
// with NoTrailingLineSeparators set to noTrailingLineSeps.
//
// This function does not modify the Options it is called on.
func (opts Options) WithNoTrailingLineSeparators(noTrailingLineSeps bool) Options {
	opts.NoTrailingLineSeparators = noTrailingLineSeps
	return opts
}

// WithParagraphSeparator returns a new Options identical to this one but with
// ParagraphSeparator set to sep. If sep is the empty string, the paragraph
// separator is interpreted as [DefaultParagraphSeparator].
//
// This function does not modify the Options it is called on.
func (opts Options) WithParagraphSeparator(sep string) Options {
	opts.ParagraphSeparator = sep
	return opts
}

// WithPreserveParagraphs returns a new Options identical to this one but
// with PreserveParagraphs set to preserve.
//
// This function does not modify the Options it is called on.
func (opts Options) WithPreserveParagraphs(preserve bool) Options {
	opts.PreserveParagraphs = preserve
	return opts
}

// WithTableBorders returns a new Options identical to this one but with
// TableBorders set to borders.
//
// This function does not modify the Options it is called on.
func (opts Options) WithTableBorders(borders bool) Options {
	opts.TableBorders = borders
	return opts
}

// WithTableHeaders returns a new Options identical to this one but with
// TableHeaders set to headers.
//
// This function does not modify the Options it is called on.
func (opts Options) WithTableHeaders(headers bool) Options {
	opts.TableHeaders = headers
	return opts
}

// WithTableCharSet returns a new Options identical to this one but with
// TableCharSet set to charSet.
//
// This function does not modify the Options it is called on.
func (opts Options) WithTableCharSet(charSet string) Options {
	opts.TableCharSet = charSet
	return opts
}
