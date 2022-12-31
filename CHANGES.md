v1.2.0 - December 31st, 2022
----------------------------
* Added InsertTable text operation
* Refactored `manip.go` and `block.go` into their own internal packages as they
didn't export any functions
* Updated gem.String to use a pointer to its grapheme cluster indexes so that
changes to it are not lost on every operation.
* All code conforms to go-staticcheck. In fact, the above issue was found with
its help. Thanks, go-staticcheck!
* Changed the name of this file from VERSION.md to CHANGES.md
* Added dates to each version header in this file
* Removed use of `_g` wrapper function in source code
* Examples for Options.WithX functions now show the specific members being
affected instead of the entire output of Options.String.
* Added more glub (mostly in text used in test functions)

v1.1.0 - December 22nd, 2022
----------------------------
* Changed Justify to no longer justify the last line in a paragraph by default
* Added JustifyLastLine option which can be set to true to restore the old
functionality
* Added script to generate grapheme cluster break tests from Unicode's published
recommended tests
* Fixed incorrect implementation of rules GB12 and GB13 from UAX #29 that were
causing some of the newly-added grapheme cluster break tests to fail
* Fixed bug where a line separator followed by a paragraph separator would in
some circumstances be interpreted as a paragraph separator followed by a line
separator
* Added CharCount Editor info function
* Added Align text operation and associated Alignment type
* Added VERSION.md file (the one you're reading now!)
* Added gofmt gate to release process. No more malformatted code making it to
main branch!
* Improved performance of Editor.Chars when the requested end index is beyond
the end of the text
* Improved performance of gem.Strings produced from gem.String.Sub; Sub now uses
the pre-calculated grapheme cache of the String it is called on and no longer
forces the sub-string to recount its graphemes

v1.0.1 - December 14th, 2022
----------------------------
* Fixed bug in implementation of rule GB11 from UAX #29 that was causing emoji
ZWJ sequences to be broken where they shouldn't be. Discovered it by testing the
emoji sequence listed in the article at https://hsivonen.fi/string-length/ by
Henri Sivonen

v1.0.0 - December 13th, 2022
----------------------------
* First stable release
* Added an Example function for every exported function
* Finalized all documentation including rewritten README.md
* Moved unit test cases for Test_Wrap that were testing WrapOpts into
Test_WrapOpts
* Fixed bug in JustifyOpts where sepEnd was calculated from the paragraph
separator prefix rather than correctly from the suffix.

v0.2.0 - December 8th, 2022
---------------------------
* Added Delete text operation
* Added Overtype text operation
* Completed initial README.md
* Added in-depth package level documentation
* Finished first draft of all function doc comments
* Added first example functions

v0.1.2 - December 7th, 2022
---------------------------
* Improved documentation
* Mostly released to test CI process and see how updated in-prog docs rendered

v0.1.1 - December 7th, 2022
---------------------------
* Added Go 1.17, 1.18, and 1.19 to testing matrix
* Added an extra unit test case for ApplyParagraphOpts
* Added more documentation

v0.1.0 - December 5th, 2022
---------------------------
* Initial pre-release
* Initial unit tests and test cases completed
* Minimal/low-effort docs
* Editor text functions available:
  * Apply
  * ApplyParagraphs
  * CollapseSpace
  * Indent
  * Insert
  * InsertDefinitionsTable
  * InsertTwoColumns
  * Justify
  * Wrap
  * (Opts versions for all of the above except Insert)
* Sub-Editor creation functions available:
  * Chars
  * CharsFrom
  * CharsTo
  * Lines
  * LinesFrom
  * LinesTo
* Options available:
  * IndentStr
  * LineSeparator
  * NoTrailingLineSeparators
  * ParagraphSeparator
  * LineSeparator
