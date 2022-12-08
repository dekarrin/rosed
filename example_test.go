package rosed

import "fmt"

func ExampleEdit() {
	ed := Edit("sample text")

	fmt.Println(ed.Text)
	// Output: sample text
}

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
