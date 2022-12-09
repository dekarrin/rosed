package rosed

import (
	"fmt"
	"strings"
)

func ExampleEdit() {
	ed := Edit("sample text")

	fmt.Println(ed.Text)
	// Output: sample text
}

// This example shows the use of Apply to add a numbered prefix to every line.
func ExampleEditor_Apply() {
	namerFunc := func(lineIdx int, line string) []string {
		newStr := fmt.Sprintf("Alpha Kid #%d: %s", lineIdx+1, line)
		return []string{newStr}
	}

	ed := Edit("Jane\nJake\nDirk\nRoxy").Apply(namerFunc)

	fmt.Println(ed.String())
	// Output:
	// Alpha Kid #1: Jane
	// Alpha Kid #2: Jake
	// Alpha Kid #3: Dirk
	// Alpha Kid #4: Roxy
}

// This example uses options to tell the Editor to use a custom LineSeparator of
// the HTML tag "<br/>", and it tells it that any trailing line ending is in
// fact the start of a new, empty line, which should be processed by the
// function.
func ExampleEditor_ApplyOpts() {
	alphaKids := 0
	namerFunc := func(lineIdx int, line string) []string {
		if line == "" {
			return []string{"Nobody here!"}
		}

		alphaKids++
		newStr := fmt.Sprintf("Alpha Kid #%d: %s", alphaKids, line)
		return []string{newStr}
	}

	opts := Options{
		NoTrailingLineSeparators: true,
		LineSeparator:            "<br/>",
	}

	// have a trailing line separator so we get the extra call of our func
	// due to NoTrailingLineSeparators
	ed := Edit("Jane<br/><br/>Dirk<br/>Roxy<br/>").ApplyOpts(namerFunc, opts)

	fmt.Println(ed.String())
	// Output:
	// Alpha Kid #1: Jane<br/>Nobody here!<br/>Alpha Kid #2: Dirk<br/>Alpha Kid #3: Roxy<br/>Nobody here!
}

// This example names the people in each paragraph in the input text based on
// who is present. Since no options are specified, it uses the default
// ParagraphSeparator of "\n\n".
func ExampleEditor_ApplyParagraphs() {
	text := ""
	text += "John\n" // paragraph 1
	text += "Rose\n"
	text += "Dave\n"
	text += "Jade\n"
	text += "\n"
	text += "Jane\n" // paragraph 2
	text += "Jake\n"
	text += "Dirk\n"
	text += "Roxy\n"
	text += "\n"
	text += "Aradia\n" // paragraph 3
	text += "Tavros\n"
	text += "Sollux\n"
	text += "Karkat"

	paraOp := func(idx int, para, sepPrefix, sepSuffix string) []string {
		var newPara string
		if strings.Index(para, "John") > -1 {
			newPara = "Beta Kids:\n" + para
		} else if strings.Index(para, "Jane") > -1 {
			newPara = "Alpha Kids:\n" + para
		} else {
			newPara = "Someone Else:\n" + para
		}
		return []string{newPara}
	}

	ed := Edit(text).ApplyParagraphs(paraOp)

	fmt.Println(ed.String())
	// Output:
	// Beta Kids:
	// John
	// Rose
	// Dave
	// Jade
	//
	// Alpha Kids:
	// Jane
	// Jake
	// Dirk
	// Roxy
	//
	// Someone Else:
	// Aradia
	// Tavros
	// Sollux
	// Karkat
}

// This example uses options to tell the Editor to use a custom
// ParagraphSeparator for splitting paragrahs, and also shows how sepPrefix and
// sepSuffix are set.
func ExampleEditor_ApplyParagraphsOpts() {
	opts := Options{ParagraphSeparator: "<P1>\n<P2>"}
	ed := Edit("para1<P1>\n<P2>para2<P1>\n<P2>para3<P1>\n<P2>para4")

	paraOp := func(idx int, para, sepPrefix, sepSuffix string) []string {
		newPara := fmt.Sprintf("(PREFIX=%s,PARA=%s,SUFFIX=%s)", sepPrefix, para, sepSuffix)
		return []string{newPara}
	}

	ed = ed.ApplyParagraphsOpts(paraOp, opts)

	fmt.Println(ed.String())
	// Output:
	// (PREFIX=,PARA=para1,SUFFIX=<P1>)<P1>
	// <P2>(PREFIX=<P2>,PARA=para2,SUFFIX=<P1>)<P1>
	// <P2>(PREFIX=<P2>,PARA=para3,SUFFIX=<P1>)<P1>
	// <P2>(PREFIX=<P2>,PARA=para4,SUFFIX=)
}

// This example gets a sub-Editor for the the "ell" part of "Hello!".
func ExampleEditor_Chars() {
	ed := Edit("Hello!")
	subEd := ed.Chars(1, 4)

	// Not doing Editor.String for the example because that would call Commit
	// and get back the starting string.
	fmt.Println(subEd.Text)
	// Output: ell
}

// This example gets a sub-Editor for the the "ello!" part of "Hello!".
func ExampleEditor_CharsFrom() {
	ed := Edit("Hello!")
	subEd := ed.CharsFrom(1)

	// Not doing Editor.String for the example because that would call Commit
	// and get back the starting string.
	fmt.Println(subEd.Text)
	// Output: ello!
}

// This example gets a sub-Editor for the the "He" part of "Hello!".
func ExampleEditor_CharsTo() {
	ed := Edit("Hello!")
	subEd := ed.CharsTo(2)

	// Not doing Editor.String for the example because that would call Commit
	// and get back the starting string.
	fmt.Println(subEd.Text)
	// Output: He
}

// This example shows how CollapseSpace collapses all whitespace sequences into
// a single space, even those that have characters not typically encountered. As
// long as Unicode considers it a whitespace character, it will be collapsed.
func ExampleEditor_CollapseSpace() {
	text := "Some \n\n\n sample text \t\n\t  \t with  \u2002   \v\v whitespace"

	ed := Edit(text).CollapseSpace()

	fmt.Println(ed.String())
	// Output: Some sample text with whitespace
}

// This example shows the use of Options with CollapseSpace to specify a
// LineSeparator that contains no whitespace. It shows that even LineSeparators
// without any Unicode space characters will be collapsed.
func ExampleEditor_CollapseSpaceOpts() {
	text := "Some \n\n\n sample<P>text \u202f <P> \t\n\t  \t with  \u2002   <P> whitespace"

	opts := Options{
		LineSeparator: "<P>",
	}

	ed := Edit(text).CollapseSpaceOpts(opts)

	fmt.Println(ed.String())
	// Output: Some sample text with whitespace
}

// This example shows how Commit can be used to commit the results of a normal
// Editor operation called on a subeditor.
func ExampleEditor_Commit_operation() {
	// get a subeditor on the last 6 chars, "string"
	subEd := Edit("Test string").CharsFrom(-6)

	// any Editor operation may be done, in our case we will call the operation
	// Indent to give it some leading whitespace
	subEd = subEd.Indent(2)

	// now call Commit to get the results merged in
	mergedEd := subEd.Commit()

	fmt.Println(mergedEd.String())
	// Output: Test 		string
}

// This example shows how Commit can be used to commit the results of manually
// assigning a subeditor's Text field.
func ExampleEditor_Commit_manualReplacement() {
	// get a subeditor on the first 7 chars, "Initial"
	subEd := Edit("Initial string").Chars(0, 7)

	// any Editor operation may be done, in our case we will manually update
	// the Text of the subeditor to replace the word entirely
	subEd.Text = "Test"

	// now call Commit to send it back up and get our original editor
	mergedEd := subEd.Commit()

	fmt.Println(mergedEd.String())
	// Output: Test string
}

// This example shows the use of CommitAll to commit all outstanding subeditors.
// It edits a simple essay outline.
func ExampleEditor_CommitAll() {
	text := ""
	text += "\t\tLine 1: An intro\n"
	text += "Line 2: A body\n"
	text += "\t\tLine 3: A conclusion"

	startingEd := Edit(text)

	// get a sub-editor for the second line, and fix the missing indent
	firstLineSubEd := startingEd.Lines(1, 2)
	firstLineSubEd = firstLineSubEd.Indent(2)

	// 'body' is a boring word, let's get a subeditor on that part and do
	// something about that.

	// -5 to -1 because that is the last 4 chars other than the trailing newline
	bodySubSubEd := firstLineSubEd.Chars(-5, -1)
	bodySubSubEd = bodySubSubEd.Overtype(0, "rationale")

	// changes are done, so commit all to get all changes merged
	mergedEd := bodySubSubEd.CommitAll()

	fmt.Println(mergedEd.String())
	// Output:
	// 		Line 1: An intro
	//		Line 2: A rationale
	//		Line 3: A conclusion
}

// This example shows the deletion of unwanted text in the editor.
func ExampleEditor_Delete() {
	ed := Edit("Here is some EXTRA text")

	ed = ed.Delete(13, 19)

	fmt.Println(ed.String())
	// Output: Here is some text
}

// This example shows a typical indent being applied to a list of people.
func ExampleEditor_Indent() {
	text := ""
	text += "* Aradia\n"
	text += "* Tavros\n"
	text += "* Sollux\n"
	text += "* Karkat\n"

	ed := Edit(text)

	ed = ed.Indent(1)

	fmt.Println(ed.String())
	// Output:
	// 	* Aradia
	//	* Tavros
	//	* Sollux
	//	* Karkat
}

// This example shows using options to specify a custom string to use for the
// indent.
func ExampleEditor_IndentOpts_indentStr() {
	text := ""
	text += "* Nepeta\n"
	text += "* Kanaya\n"
	text += "* Terezi\n"
	text += "* Vriska\n"

	opts := Options{
		IndentStr: "==>",
	}

	ed := Edit(text)

	ed = ed.IndentOpts(2, opts)

	fmt.Println(ed.String())
	// Output:
	// ==>==>* Nepeta
	// ==>==>* Kanaya
	// ==>==>* Terezi
	// ==>==>* Vriska
}

// This example shows using options to respect paragraph breaks.
func ExampleEditor_IndentOpts_preserveParagraphs() {
	text := ""
	text += "Beta Kids:\n"
	text += "* John\n"
	text += "* Rose\n"
	text += "* Dave\n"
	text += "* Jade\n"
	text += "\n"
	text += "Alpha Kids:\n"
	text += "* Jane\n"
	text += "* Jake\n"
	text += "* Dirk\n"
	text += "* Roxy\n"

	opts := Options{
		PreserveParagraphs: true,
	}

	ed := Edit(text)

	ed = ed.IndentOpts(1, opts)

	fmt.Println(ed.String())
	// Output:
	//  Beta Kids:
	// 	* John
	// 	* Rose
	// 	* Dave
	// 	* Jade
	//
	// 	Alpha Kids:
	// 	* Jane
	// 	* Jake
	// 	* Dirk
	// 	* Roxy
}
