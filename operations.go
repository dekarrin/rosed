package rosed

// this file contains operations performed by Editors.

import (
	"fmt"
	"strings"
)

// LineOperation is a function that takes a zero-indexed line number and the
// contents of that line and performs some operation on the string to get a
// new string to replace the contents of the line with.
//
// The return value for a LineOperation is a slice of lines to insert at the
// old line position. This can be used to delete the line or insert additional
// new ones; to insert, just include the new lines in the returned slice in the
// proper position relative to the old line in the slice, and to delete the
// original line, a slice with len < 1 can be returned.
//
// The idx will always be the index of the line before any transformations were
// applied; i.e. if used in ForEachLine, a call to a LineOperation with idx = 4
// will always be after a call with idx = 3, regardless of the size of the
// returned slice in the prior call.
type LineOperation func(idx int, line string) []string

// Apply applies the given LineOperation for each line in the text. Line
// termination at the last line is transparently handled as per the options
// currently set on the Editor.
func (ed Editor) Apply(op LineOperation) Editor {
	return ed.ApplyOpts(op, ed.Options)
}

// ApplyOpts applies the given LineOperation for each line in the text. Line
// termination at the last line is transparently handled as per the provided
// options.
func (ed Editor) ApplyOpts(op LineOperation, opts Options) Editor {
	opts = opts.WithDefaults()

	lines := ed.linesSep(opts.LineSeparator)
	applied := make([]string, 0, len(lines))

	for idx, line := range lines {
		newLines := op(idx, line)
		if len(newLines) > 0 {
			applied = append(applied, newLines...)
		}
	}

	// make sure to preserve the last line sep if it exists; it will have been
	// clobbered in call to lines() if it was.
	if strings.HasSuffix(ed.Text, opts.LineSeparator) {
		applied = append(applied, "")
	}

	ed.Text = strings.Join(applied, opts.LineSeparator)
	return ed
}

// Indent adds the currently configured indent string level times at the start
// of each line in the Editor. If level is 0 or less, the text is unchanged.
func (ed Editor) Indent(level int) Editor {
	return ed.IndentOpts(level, ed.Options)
}

// IndentOpts adds an indent string level times at the start of each line in the
// Editor. If level is 0 or less, the text is unchanged.
//
// The provided Options object is used to override the options currently set on
// the Editor for the indent. LineSeparator, IndentStr, and
// NoTrailingLineSeparators are read from the provided Options obejct.
func (ed Editor) IndentOpts(level int, opts Options) Editor {
	if level < 1 {
		// caller wants fewer than 1 indent. Okay, that is zero; return
		// unchanged
		return ed
	}

	indent := opts.WithDefaults().IndentStr
	doIndent := func(_ int, line string) []string {
		newLine := indent + line

		// only have the one line, returne that
		return []string{newLine}
	}

	return ed.ApplyOpts(doIndent, opts)
}

// WrapOpts performs a wrap of all text to the given width. The provided options
// are used instead of the Editor's built-in options. If width is less than 2,
// it is assumed to be 2 because no meaningful wrap algorithm can be applied to
// anything smaller.
func (ed Editor) WrapOpts(width int, opts Options) []string {
	opts = opts.WithDefaults()

	if width < 2 {
		width = 2
	}

	// get the correct 'complete' paragraph separator based on our line ending
	// mode

	// split the paragraph separator about its line separators so we can see any
	// extra chars that will be chopped off while in a preserve-mode wrap that
	// messes with line separators
	var paraSepPrevSuffix, paraSepNextPrefix string
	var paraSepLines []string
	parts := strings.Split(opts.ParagraphSeparator, opts.LineSeparator)
	paraSepPrevSuffix = parts[0]
	if len(parts) > 1 {
		paraSepLines = parts[1 : len(parts)-1]
		paraSepNextPrefix = parts[len(parts)-1]
	}

	var totalLines int
	var wrappedParagraphs [][]string
	if options.PreserveParagraphs {
		var paragraphs = strings.Split(text, options.ParagraphSeparator)
		wrappedParagraphs = make([][]string, len(paragraphs))
		for idx, para := range paragraphs {
			trimParaSuffix := false
			trimParaPrefix := false
			// to get proper width, make sure we add the prefix and suffix to the para
			if idx > 0 {
				// if there are prev paragraphs, add the next line prefix
				para = paraSepNextPrefix + para
				trimParaPrefix = true
			}
			if idx+1 < len(paragraphs) {
				// if there are more paragraphs, add the prev line suffix
				para += paraSepPrevSuffix
				trimParaSuffix = true

				// update totalLines to include the "between paragraphs" lines added by the separator.
				totalLines += len(paraSepLines)
			}

			para = options.ParagraphPrefix + para + options.ParagraphSuffix
			wrappedPara := doPrecalculatedWidthWrap(para, widthWithoutAffixes)
			// trim off paragraph prefixes and suffixes to let the later routine add them outside of the
			// prefix/suffix specified in options (if any)
			wrappedPara[0] = strings.TrimPrefix(wrappedPara[0], options.ParagraphPrefix)
			wrappedPara[len(wrappedPara)-1] = strings.TrimSuffix(wrappedPara[len(wrappedPara)-1], options.ParagraphPrefix)
			totalLines += len(wrappedPara)

			// trim off para break prefixes and suffixes to let the later routine add them outside of the
			// prefix/suffix specified in options (if any)
			if trimParaPrefix {
				wrappedPara[0] = strings.TrimPrefix(wrappedPara[0], paraSepNextPrefix)
			}
			if trimParaSuffix {
				wrappedPara[len(wrappedPara)-1] = strings.TrimSuffix(wrappedPara[len(wrappedPara)-1], paraSepPrevSuffix)
			}
			wrappedParagraphs[idx] = wrappedPara
		}
	} else {
		para := options.ParagraphPrefix + text + options.ParagraphSuffix
		wrapped := doPrecalculatedWidthWrap(para, widthWithoutAffixes)
		// trim off paragraph prefixes and suffixes to let the later routine add them outside of the
		// prefix/suffix specified in options (if any)
		wrapped[0] = strings.TrimPrefix(wrapped[0], options.ParagraphPrefix)
		wrapped[len(wrapped)-1] = strings.TrimSuffix(wrapped[len(wrapped)-1], options.ParagraphPrefix)
		totalLines = len(wrapped)
		wrappedParagraphs = [][]string{wrapped}
	}

	// add affixes and paragraph breaks to the lines
	lines := make([]string, totalLines)
	curParaLines := lines
	for paraIdx, para := range wrappedParagraphs {
		for lineIdx := range para {
			curParaLines[lineIdx] = fmt.Sprintf("%s%s%s", options.Prefix, para[lineIdx], options.Suffix)

			if lineIdx == 0 {
				// if on the first line of the paragraph, add the paragraph prefix
				curParaLines[lineIdx] = options.ParagraphPrefix + curParaLines[lineIdx]

				// if on the first line of any paragraph after the first, add the separator prefix to it
				if paraIdx > 0 {
					curParaLines[lineIdx] = paraSepNextPrefix + curParaLines[lineIdx]
				}
			}

			if lineIdx+1 >= len(para) {
				// insert paragraph suffix after last line
				curParaLines[lineIdx] = curParaLines[lineIdx] + options.ParagraphSuffix
			}
		}
		extraLines := 0
		if paraIdx+1 < len(wrappedParagraphs) {
			curParaLines[len(para)-1] = curParaLines[len(para)-1] + paraSepPrevSuffix

			for paraSepIdx, paraSepLine := range paraSepLines {
				curParaLines[len(para)+paraSepIdx] = paraSepLine
			}
			extraLines += len(paraSepLines)
		}

		// set destination lines to refer to the end of where we have written
		// for next loop
		curParaLines = curParaLines[len(para)+extraLines:]
	}

	return lines
}
