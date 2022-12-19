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

// This example shows use of Align to take inconsistently tabbed input and
// normalize each line to no indent before doing further operations.
func ExampleEditor_Align_left() {
	text := "\t\t\tYour name is TEREZI PYROPE.\n"
	text += "\n"
	text += "\t  You are pretty enthusiastic about dragons.\n"
	text += "     But you have a PARTICULAR AFFECTION\n"
	text += "\t for their COLORFUL SCALES, which you gather\n"
	text += "\t\tand use to decorate your hive.\n"
	
	ed := Edit(text)
	
	ed = ed.Align(Left, 45)
	
	// also add a pipe char at start and end of each line to show that space is
	// added to fill out each line to the specified width:
	ed = ed.Apply(func(idx int, line string) []string {
		return []string{"|" + line + "|"}
	})
	
	fmt.Println(ed.String())
	// Output:
	// |Your name is TEREZI PYROPE.                  |
	// |                                             |
	// |You are pretty enthusiastic about dragons.   |
	// |But you have a PARTICULAR AFFECTION          |
	// |for their COLORFUL SCALES, which you gather  |
	// |and use to decorate your hive.               |
}

// This example shows use of Align to make every line align to the right side.
func ExampleEditor_Align_right() {
	text := "Your name is TEREZI PYROPE.\n"
	text += "\n"
	text += "You are pretty enthusiastic about dragons.\n"
	text += "But you have a PARTICULAR AFFECTION\n"
	text += "for their COLORFUL SCALES, which you gather\n"
	text += "and use to decorate your hive.\n"
	
	ed := Edit(text)
	
	ed = ed.Align(Right, 45)
	
	// also add a pipe char at start and end of each line to show that space is
	// added to fill out each line to the specified width:
	ed = ed.Apply(func(idx int, line string) []string {
		return []string{"|" + line + "|"}
	})
	
	fmt.Println(ed.String())
	// Output:
	// |                  Your name is TEREZI PYROPE.|
	// |                                             |
	// |   You are pretty enthusiastic about dragons.|
	// |          But you have a PARTICULAR AFFECTION|
	// |  for their COLORFUL SCALES, which you gather|
	// |               and use to decorate your hive.|
}

// This example shows use of Align to center each line of text.
func ExampleEditor_Align_center() {
	text := "Your name is TEREZI PYROPE.\n"
	text += "\n"
	text += "You are pretty enthusiastic about dragons.\n"
	text += "But you have a PARTICULAR AFFECTION\n"
	text += "for their COLORFUL SCALES, which you gather\n"
	text += "and use to decorate your hive.\n"
	
	ed := Edit(text)
	
	ed = ed.Align(Center, 45)
	
	// also add a pipe char at start and end of each line to show that space is
	// added to fill out each line to the specified width:
	ed = ed.Apply(func(idx int, line string) []string {
		return []string{"|" + line + "|"}
	})
	
	fmt.Println(ed.String())
	// Output:
	// |         Your name is TEREZI PYROPE.         |
	// |                                             |
	// |  You are pretty enthusiastic about dragons. |
	// |     But you have a PARTICULAR AFFECTION     |
	// | for their COLORFUL SCALES, which you gather |
	// |        and use to decorate your hive.       |
}

// This example shows use of a custom line separator with Align to respect text
// that uses the HTML tag "<br/>\n" to separate lines.
func ExampleEditor_AlignOpts() {
	text := "Your name is TEREZI PYROPE.<br/>\n"
	text += "<br/>\n"
	text += "You are pretty enthusiastic about dragons.<br/>\n"
	text += "But you have a PARTICULAR AFFECTION<br/>\n"
	text += "for their COLORFUL SCALES, which you gather<br/>\n"
	text += "and use to decorate your hive.<br/>\n"
	
	opts := Options{
		LineSeparator: "<br/>\n",
	}
	
	ed := Edit(text)
	
	ed = ed.AlignOpts(Center, 45, opts)
	
	// also add a pipe char at start and end of each line to show that space is
	// added to fill out each line to the specified width:
	/*ed = ed.Apply(func(idx int, line string) []string {
		return []string{"|" + line + "|"}
	})*/
	
	fmt.Println(ed.String())
	// Output:
	//          Your name is TEREZI PYROPE.         <br/>
	//                                              <br/>
	//   You are pretty enthusiastic about dragons. <br/>
	//      But you have a PARTICULAR AFFECTION     <br/>
	//  for their COLORFUL SCALES, which you gather <br/>
	//         and use to decorate your hive.       <br/>
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

// This example gets the length of several different strings.
func ExampleEditor_CharCount() {
	emptyCount := Edit("").CharCount()
	fmt.Println(emptyCount)

	testCount := Edit("test").CharCount()
	fmt.Println(testCount)

	emojiCount := Edit("👩‍❤️‍💋‍👩").CharCount()
	fmt.Println(emojiCount)

	// Output:
	// 0
	// 4
	// 1
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

// This example inserts the text "burb" in the middle of the editor text.
func ExampleEditor_Insert() {
	ed := Edit("S world!")

	ed = ed.Insert(1, "burb")

	fmt.Println(ed.String())
	// Output: Sburb world!
}

// This example produces the table seen above.
func ExampleEditor_InsertDefinitionsTable() {
	ed := Edit("")

	var johnDef, roseDef, daveDef, jadeDef string

	johnDef += "Has a passion for REALLY TERRIBLE MOVIES. Likes "
	johnDef += "to program computers but is NOT VERY GOOD AT IT."

	roseDef += "Has a passion for RATHER OBSCURE LITERATURE. "
	roseDef += "Enjoys creative writing and is SOMEWHAT "
	roseDef += "SECRETIVE ABOUT IT."

	daveDef += "Has a penchant for spinning out UNBELIEVABLY ILL "
	daveDef += "JAMS with his TURNTABLES AND MIXING GEAR. Likes "
	daveDef += "to rave about BANDS NO ONE'S EVER HEARD OF BUT HIM."

	jadeDef += "Has so many INTERESTS, she has trouble keeping "
	jadeDef += "track of them all, even with an assortment of "
	jadeDef += "COLORFUL REMINDERS on her fingers to help sort out "
	jadeDef += "everything on her mind."

	defs := [][2]string{
		{"John", johnDef},
		{"Rose", roseDef},
		{"Dave", daveDef},
		{"Jade", jadeDef},
	}

	ed = ed.InsertDefinitionsTable(0, defs, 76)

	fmt.Println("TABLE:")
	fmt.Println(ed.String())
	// Output:
	// TABLE:
	//   John  - Has a passion for REALLY TERRIBLE MOVIES. Likes to program
	//           computers but is NOT VERY GOOD AT IT.
	//
	//   Rose  - Has a passion for RATHER OBSCURE LITERATURE. Enjoys creative
	//           writing and is SOMEWHAT SECRETIVE ABOUT IT.
	//
	//   Dave  - Has a penchant for spinning out UNBELIEVABLY ILL JAMS with his
	//           TURNTABLES AND MIXING GEAR. Likes to rave about BANDS NO ONE'S
	//           EVER HEARD OF BUT HIM.
	//
	//   Jade  - Has so many INTERESTS, she has trouble keeping track of them all,
	//           even with an assortment of COLORFUL REMINDERS on her fingers to
	//           help sort out everything on her mind.
}

// This example uses options to set a custom paragraph separator so that terms
// are separated by dashes instead of blank lines.
func ExampleEditor_InsertDefinitionsTableOpts() {
	ed := Edit("")

	defs := [][2]string{
		{"Apple", "A delicious fruit that can be eaten by pretty much anybody who likes fruit."},
		{"Bottle", "Holds liquids."},
		{"Crow's Egg", "The egg of a crow, who may or may not go CAW-CAW."},
		{"Dog Pinata", "If you hit it, candy will come out."},
	}

	opts := Options{
		ParagraphSeparator: "\n------------------------------------------------------------\n",
	}

	ed = ed.InsertDefinitionsTableOpts(0, defs, 60, opts)

	fmt.Println("TABLE:")
	fmt.Println(ed.String())
	// Output:
	// TABLE:
	//   Apple       - A delicious fruit that can be eaten by
	//                 pretty much anybody who likes fruit.
	// ------------------------------------------------------------
	//   Bottle      - Holds liquids.
	// ------------------------------------------------------------
	//   Crow's Egg  - The egg of a crow, who may or may not go
	//                 CAW-CAW.
	// ------------------------------------------------------------
	//   Dog Pinata  - If you hit it, candy will come out.
	//
}

// This example creates two columns from two runs of text.
func ExampleEditor_InsertTwoColumns() {
	leftText := "Karkalicious, definition: makes Terezi loco. "
	leftText += "She wants to know the secrets that she can't "
	leftText += "taste in my photo."

	rightText := "A young man stands in his bedroom. It just so happens that "
	rightText += "today, the 13th of April, 2009, is this young man's birthday. "
	rightText += "Though it was thirteen years ago he was given life, it is "
	rightText += "only today he will be given a name!"

	// insert it at the start of the editor
	pos := 0

	// minimum 3 spaces between each column at their closest point
	minSpace := 3

	// wrap the entire layout to 50 chars
	width := 50

	// make the left column take up 40% of the available space
	leftPercent := 0.4

	ed := Edit("").InsertTwoColumns(pos, leftText, rightText, minSpace, width, leftPercent)

	fmt.Println(ed.String())
	// Output:
	// Karkalicious,        A young man stands in his
	// definition: makes    bedroom. It just so happens
	// Terezi loco. She     that today, the 13th of
	// wants to know the    April, 2009, is this young
	// secrets that she     man's birthday. Though it was
	// can't taste in my    thirteen years ago he was
	// photo.               given life, it is only today
	//                      he will be given a name!
}

// This example uses options to tell the Editor to include the HTML tag "<br/>"
// followed by a literal newline to separate lines in the output columns.
func ExampleEditor_InsertTwoColumnsOpts() {
	left := "A sample short text run that wraps once."
	right := "This run of text should also take up 2 lines."

	pos := 0
	minSpace := 3
	width := 50
	leftPercent := 0.5

	opts := Options{
		LineSeparator: "<br/>\n",
	}

	ed := Edit("").InsertTwoColumnsOpts(pos, left, right, minSpace, width, leftPercent, opts)

	fmt.Println(ed.String())
	// Output:
	// A sample short text run   This run of text should<br/>
	// that wraps once.          also take up 2 lines.<br/>
}

// This example shows that only Editors created from a sub-editor producing
// function will return true for IsSubEditor.
func ExampleEditor_IsSubEditor() {
	notASubEd := Edit("Hello, world!")
	subEd := Edit("Sub, Sburb?").CharsFrom(5)

	fmt.Println(notASubEd.IsSubEditor())
	fmt.Println(subEd.IsSubEditor())
	// Output:
	// false
	// true
}

// This example shows justification of input text to a length of 60 chars.
func ExampleEditor_Justify() {
	input := "Some words that will have spacing justified.\n"
	input += "This occurs on a per-line basis.\n"
	input += "Lines closer to the justified length have less adjustment."

	ed := Edit(input)

	ed = ed.Justify(60)

	fmt.Println(ed.String())
	// Output:
	// Some    words   that    will    have   spacing    justified.
	// This      occurs       on       a      per-line       basis.
	// Lines  closer to the justified length have less  adjustment.
}

// This example shows the use of options to make the justification respect a
// rather contrived paragraph splitter of "\nPARA SPLIT\n"
func ExampleEditor_JustifyOpts() {
	input := "Some words that will have spacing justified.\n"
	input += "This occurs on a per-line basis.\n"
	input += "Lines closer to the justified length have less adjustment.\n"
	input += "PARA SPLIT\n"
	input += "This is a second paragraph that is used to show how\n"
	input += "paragraphs can be respected with options.\n"

	opts := Options{
		PreserveParagraphs: true,
		ParagraphSeparator: "\nPARA SPLIT\n",
	}

	ed := Edit(input)

	ed = ed.JustifyOpts(60, opts)

	fmt.Println(ed.String())
	// Output:
	// Some    words   that    will    have   spacing    justified.
	// This      occurs       on       a      per-line       basis.
	// Lines  closer to the justified length have less  adjustment.
	// PARA SPLIT
	// This  is a  second  paragraph  that  is  used  to  show  how
	// paragraphs    can     be     respected     with     options.
}

// This example shows querying the number of lines for a variety of text.
func ExampleEditor_LineCount() {
	zeroLinesEd := Edit("")
	fmt.Println(zeroLinesEd.LineCount())

	oneLineEd := Edit("Line 1")
	fmt.Println(oneLineEd.LineCount())

	twoLinesEd := Edit("Line 1\nLine 2")
	fmt.Println(twoLinesEd.LineCount())
	// Output:
	// 0
	// 1
	// 2
}

// This example shows the two different ways currently-set options affect the
// output of LineCount. In the first case, it shows how NoTrailingLineSeparators
// affects the number of lines in an empty editor, and the in the second case,
// it shows the use of a custom LineSeparator to change how lines are counted.
func ExampleEditor_LineCount_options() {
	noTrailingOpts := Options{NoTrailingLineSeparators: true}
	emptyButNoTrailingEd := Edit("").WithOptions(noTrailingOpts)
	fmt.Println(emptyButNoTrailingEd.LineCount())

	customLineOpts := Options{LineSeparator: "<br/>"}
	customLineEndEd := Edit("Line 1<br/>Line 2").WithOptions(customLineOpts)
	fmt.Println(customLineEndEd.LineCount())
	// Output:
	// 1
	// 2
}

// This example gets a subeditor on the middle three lines of a five-line
// string.
func ExampleEditor_Lines() {
	ed := Edit("Act 1\nAct 2\nAct 3\nAct 4\nAct 5")

	ed = ed.Lines(1, 4)

	// Not doing Editor.String for the example because that would call Commit
	// and get back the starting string.
	fmt.Println(ed.Text)
	// Output:
	// Act 2
	// Act 3
	// Act 4
}

// This example gets a subeditor on the last two lines of a five-line string.
func ExampleEditor_LinesFrom() {
	ed := Edit("Act 1\nAct 2\nAct 3\nAct 4\nAct 5")

	ed = ed.LinesFrom(3)

	// Not doing Editor.String for the example because that would call Commit
	// and get back the starting string.
	fmt.Println(ed.Text)
	// Output:
	// Act 4
	// Act 5
}

// This example gets a subeditor on the first three lines of a five-line string.
func ExampleEditor_LinesTo() {
	ed := Edit("Act 1\nAct 2\nAct 3\nAct 4\nAct 5")

	ed = ed.LinesTo(3)

	// Not doing Editor.String for the example because that would call Commit
	// and get back the starting string.
	fmt.Println(ed.Text)
	// Output:
	// Act 1
	// Act 2
	// Act 3
}

// This example uses Overtype to replace a part of a greeting message. This
// works so nicely in the example because the replacement is the exact same
// length as the replaced text. If it were of a longer length, it would end up
// overwriting the text that comes after the replaced text; if it were shorter,
// it would end up not replacing the entire intended section.
func ExampleEditor_Overtype() {
	ed := Edit("How are you, Miss Lalonde?")

	ed = ed.Overtype(4, "goes it")

	fmt.Println(ed.String())
	// Output: How goes it, Miss Lalonde?
}

// This example uses String on a normal Editor to get its text.
func ExampleEditor_String() {
	ed := Edit("Some text")

	text := ed.String()

	fmt.Println(text)
	// Output: Some text
}

// This example uses String on a sub-editor to avoid having to explicitly call
// [Editor.Commit] or [Editor.CommitAll] after performing an operation on the
// sub-editor.
func ExampleEditor_String_subEditor() {
	ed := Edit("Act 1\nAct 2\nAct BLAH\nAct 4\nAct 5")
	subEd := ed.Lines(2, 3)

	subEd = subEd.Delete(4, 8).Insert(4, "3")

	fullText := subEd.String()
	fmt.Println(fullText)
	// Output:
	// Act 1
	// Act 2
	// Act 3
	// Act 4
	// Act 5
}

// This example sets the IndentStr property of the Options on the Editor.
func ExampleEditor_WithOptions() {
	ed := Edit("Vriska Serket")

	ed = ed.WithOptions(ed.Options.WithIndentStr("-->"))

	fmt.Println(ed.Options.IndentStr)
	// Output: -->
}

// This example shows wrapping applied to a long string.
func ExampleEditor_Wrap() {
	ed := Edit("Your name is VRISKA SERKET. You are a master of EXTREME ROLEPLAYING.")

	ed = ed.Wrap(25)

	fmt.Println(ed.String())
	// Output:
	// Your name is VRISKA
	// SERKET. You are a master
	// of EXTREME ROLEPLAYING.
}

// This example uses options to tell the wrap to use a line ending consisting of
// the HTML tag "<br/>" followed by a new-line, and to respect paragraphs
// separated by a double "<br/>\n". It also shows how pre-wrapped text will have
// the hard wraps removed prior to the new wrap.
func ExampleEditor_WrapOpts() {
	text := "Your name is VRISKA SERKET.<br/>\n"
	text += "<br/>\n"
	text += "You are a master of EXTREME ROLEPLAYING. You can't get enough<br/>\n"
	text += "of it, or really any game of high stakes and chance. You have<br/>\n"
	text += "persisted with the habit even in spite of your ACCIDENT. But<br/>\n"
	text += "then again, you don't have much choice.<br/>\n"
	text += "<br/>\n"
	text += "You are something of an APOCALYPSE BUFF, which is something you<br/>\n"
	text += "can be on Alternia. You are fascinated by end of the world<br/>\n"
	text += "scenarios."

	opts := Options{
		LineSeparator:      "<br/>\n",
		ParagraphSeparator: "<br/>\n<br/>\n",
		PreserveParagraphs: true,
	}

	ed := Edit(text)

	ed = ed.WrapOpts(50, opts)

	fmt.Println(ed.String())
	// Output:
	// Your name is VRISKA SERKET.<br/>
	// <br/>
	// You are a master of EXTREME ROLEPLAYING. You can't<br/>
	// get enough of it, or really any game of high<br/>
	// stakes and chance. You have persisted with the<br/>
	// habit even in spite of your ACCIDENT. But then<br/>
	// again, you don't have much choice.<br/>
	// <br/>
	// You are something of an APOCALYPSE BUFF, which is<br/>
	// something you can be on Alternia. You are<br/>
	// fascinated by end of the world scenarios.
}

func ExampleOptions_String() {
	opts := Options{IndentStr: "-->"}

	str := opts.String()

	fmt.Println(str)
	// Output:
	// Options{ParagraphSeparator: "", LineSeparator: "", IndentStr: "-->", NoTrailingLineSeparators: false, PreserveParagraphs: false}
}

// This example shows how WithDefaults can be called to set all currently unset
// properties to their default values while leaving the set values alone.
func ExampleOptions_WithDefaults() {
	opts := Options{
		LineSeparator:      "<br/>",
		PreserveParagraphs: true,
	}
	fmt.Println(opts)

	optsDefault := opts.WithDefaults()
	fmt.Println(optsDefault)
	// Output:
	// Options{ParagraphSeparator: "", LineSeparator: "<br/>", IndentStr: "", NoTrailingLineSeparators: false, PreserveParagraphs: true}
	// Options{ParagraphSeparator: "\n\n", LineSeparator: "<br/>", IndentStr: "\t", NoTrailingLineSeparators: false, PreserveParagraphs: true}
}

func ExampleOptions_WithIndentStr() {
	opts := Options{}

	opts = opts.WithIndentStr("-->")

	fmt.Println(opts.String())
	// Output:
	// Options{ParagraphSeparator: "", LineSeparator: "", IndentStr: "-->", NoTrailingLineSeparators: false, PreserveParagraphs: false}
}

func ExampleOptions_WithLineSeparator() {
	opts := Options{}

	opts = opts.WithLineSeparator("<br/>")

	fmt.Println(opts.String())
	// Output:
	// Options{ParagraphSeparator: "", LineSeparator: "<br/>", IndentStr: "", NoTrailingLineSeparators: false, PreserveParagraphs: false}
}

func ExampleOptions_WithNoTrailingLineSeparators() {
	opts := Options{}

	opts = opts.WithNoTrailingLineSeparators(true)

	fmt.Println(opts.String())
	// Output:
	// Options{ParagraphSeparator: "", LineSeparator: "", IndentStr: "", NoTrailingLineSeparators: true, PreserveParagraphs: false}
}

func ExampleOptions_WithParagraphSeparator() {
	opts := Options{}

	opts = opts.WithParagraphSeparator("<P>")

	fmt.Println(opts.String())
	// Output:
	// Options{ParagraphSeparator: "<P>", LineSeparator: "", IndentStr: "", NoTrailingLineSeparators: false, PreserveParagraphs: false}
}

func ExampleOptions_WithPreserveParagraphs() {
	opts := Options{}

	opts = opts.WithPreserveParagraphs(true)

	fmt.Println(opts.String())
	// Output:
	// Options{ParagraphSeparator: "", LineSeparator: "", IndentStr: "", NoTrailingLineSeparators: false, PreserveParagraphs: true}
}