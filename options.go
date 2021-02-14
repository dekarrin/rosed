package rosed

const (
	// DefaultParagraphSeparator is the sequence that separates paragraphs.
	DefaultParagraphSeparator = "\n\n"

	// DefaultLineSeparator is the default line separator sequence.
	DefaultLineSeparator = "\n"

	// DefaultIndentString is the string used for a single horizontal indent
	// by default.
	DefaultIndentString = "\t"
)

// Options are the options to an Editor. The Zero-value is an Options with all
// items set to defaults.
type Options struct {

	// LineSeparator is the string that the Editor considers to signify the
	// end of a line. If this is set to "", the Editor will use the default
	// string of DefaultLineSeparator.
	LineSeparator string

	// IndentStr is the string that is used for a single horizontal indent. If
	// this is set to "", the Editor will use the default string of "\t".
	IndentStr string

	// NoTrailingLineSeparators is whether the Editor considers lines to not end
	// with the separator, and thus would assume that a properly formatted line
	// does not include a line separator at the end even if it is the last line.
	//
	// If this is set to false (the default), line separator chars are assumed
	// to signify the end of the line.
	NoTrailingLineSeparators bool
}

// WithLineSeparator returns a new Options identical to this one but with the
// LineSeparator set to sep. If sep is the empty string, the line separator is
// interpreted as DefaultLineSeparator.
func (opts Options) WithLineSeparator(sep string) Options {
	opts.LineSeparator = sep
	return opts
}

// WithIndentStr returns a new Options identical to this one but with the
// IndentStr set to str. If str is the empty string, the indent str is
// interpreted as the default indent string ("\t").
func (opts Options) WithIndentStr(str string) Options {
	opts.IndentStr = str
	return opts
}

// WithNoTrailingLineSeparators returns a new Options identical to this one but
// with NoTrailingLineSeparators set to noTrailingLineSeps.
func (opts Options) WithNoTrailingLineSeparators(noTrailingLineSeps bool) Options {
	opts.NoTrailingLineSeparators = noTrailingLineSeps
	return opts
}

// WithDefaults returns a copy of the options with all blank members filled with
// their defaults. Internally, this function is used on user-provided Options
// objects in order to get ready-to-use copies.
func (opts Options) WithDefaults() Options {
	if opts.LineSeparator == "" {
		opts.LineSeparator = DefaultLineSeparator
	}
	if opts.IndentStr == "" {
		opts.IndentStr = DefaultIndentString
	}
	return opts
}
