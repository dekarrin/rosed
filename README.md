rosed
=====

![Tests Status Badge](https://github.com/dekarrin/rosed/actions/workflows/tests.yml/badge.svg?branch=main&event=push)
[![Go Reference](https://pkg.go.dev/badge/github.com/dekarrin/rosed.svg)](https://pkg.go.dev/github.com/dekarrin/rosed)

Pronounced "ROSE-edd". Fluent text editing and layout library for Go.

This library treats text as a sequence of "grapheme clusters" (user-visible
characters) as opposed to actual runes or bytes. This allows wrapping and other
character-counting/index operations to function in an expected way regardless of
whether visible characters are made up of Unicode decomposed sequences or
pre-composed sequences.

### Graphemes? What Are Those?

Take something like "Ý", which appears to be one single character but behind the
scenes could be a single precomposed character rune or multiple runes that
compose into a single apparent character using Unicode combining marks, such as
the Unicode codepoint for uppercase "Y" (U+0079) followed by the codepoint for
the combining acute accent (U+0301) (not to be confused with the *non-combining*
codepoint for the acute accent symbol by itself, U+00B4. Naturally. Isn't
Unicode *fun?*).

Regardless of its representation in runes, we as humans see the Y with an acute
accent and will usually say "ah, that is one character", and that is a grapheme
cluster (or just grapheme for short when referring to the thing being displayed
as opposed to the runes that make up that thing). This library takes that
opinion as well, and will operate on text in a consistent way regardless of how
each grapheme is represented.

You can find out more information on the Unicode definition of "grapheme
cluster" under [Chapter 2 of the Unicode Standard](https://www.unicode.org/versions/Unicode15.0.0/ch02.pdf),
in particular see the heading '"Characters" and Grapheme Clusters" at the end of
section 2.11, "Combining Characters". If you want, you can also dig into section
2.12, "Equivalant Sequences", to see how different grapheme clusters (codepoint
sequences) can translate into the same visual grapheme.

### What This Library Is Not

* Performant - Especially for initial release the idea is to create a library
that is easy to use. Better performance may come in later versions but that is
not a goal at the time.
* Collating - This libarary does not support Unicode collation. It may come at
a later date.

## Example Usages

This library aims to make it easy to work with text in a CLI environment, or
really any that has fixed-width text in it. This section shows some of the
things you can do with `rosed`, but by no means is it all of it. Check the
[Full Reference Docs](https://pkg.go.dev/github.com/dekarrin/rosed) for a
listing of all available functions as well as examples of use.

In general, to use rosed, first create an Editor, use it to modify text, then
get it back by converting it to a string. As easy as plucking a rose from the
bed of thorns that grows in the midst of the most eldritch of gardens. Perhaps
easier, in fact.

### Wrap Text

You can use rosed to wrap text to a certain width:

*[Try it on the Go Playground](https://go.dev/play/p/C-OzpH-HENV)*

```golang
text := "It's hard, being a kid and growing up. It's hard and nobody understands."

wrapped := rosed.Edit(text).Wrap(30).String()

fmt.Println(wrapped)
```

Output:

```
It's hard, being a kid and
growing up. It's hard and
nobody understands.
```

You can do this even if it's already been hard-wrapped:

*[Try it on the Go Playground](https://go.dev/play/p/NsyWZLiXuU0)*

```golang
text := "It's hard, being\na kid and growing up.\nIt's hard and nobody\nunderstands."

wrapped := rosed.Edit(text).Wrap(30).String()

fmt.Println(wrapped)
```

### Justify Text

You can use rosed to ensure that all lines in the text take up the same width,
and add spacing equally between words on each line to ensure this is the case.
This is also known as justifying a block of text:

*[Try it on the Go Playground](https://go.dev/play/p/pcI7hQIJH_R)*

```golang
text := "I WARNED YOU ABOUT STAIRS\n"
text += "BRO! I TOLD YOU DOG!\n"
text += "I TOLD YOU MAN!"

justified := rosed.Edit(text).Justify(30).String()

fmt.Println(justified)
```

Output:

```
I   WARNED  YOU  ABOUT  STAIRS
BRO!    I   TOLD   YOU    DOG!
I TOLD YOU MAN!
```

You could combine this with wrapping to account for the lines that didn't quite
end up the same length as the others:

*[Try it on the Go Playground](https://go.dev/play/p/PCW9dcubxSX)*

```golang
text := "Your name is KANAYA MARYAM. You are one of the few of your "
text += "kind who can withstand the BLISTERING ALTERNIAN SUN, and "
text += "perhaps the only who enjoys the feel of its rays. As such, "
text += "you are one of the few of your kind who has taken a shining "
text += "to LANDSCAPING. You have cultivated a lush oasis."

wrappedAndJustified := rosed.Edit(text).Wrap(50).Justify(50).String()

fmt.Println(wrappedAndJustified)
```

Output:

```
Your name is KANAYA MARYAM. You are one of the few
of  your kind  who  can  withstand  the BLISTERING
ALTERNIAN SUN, and perhaps the only who enjoys the
feel  of its rays. As such, you are one of the few
of   your  kind   who  has  taken  a  shining   to
LANDSCAPING. You have cultivated a lush oasis.
```

### Definitions Table

You can use this library to build a table of terms and their definitions:

*[Try it on the Go Playground](https://go.dev/play/p/A8lHKVDDv2Q)*

```golang
aradiaDef := "Once had a fondness for ARCHEOLOGY, though now has trouble "
aradiaDef += "recalling this passion."

tavDef := "Known to be heavily arrested by FAIRY TALES AND FANTASY STORIES. "
tavDef += "Has an acute ability to COMMUNE WITH THE MANY CREATURES OF "
tavDef += "ALTERNIA."

solluxDef := "Is good at computers, and knows ALL THE CODES. All of them."

karkatDef := "Has a passion for RIDICULOUSLY TERRIBLE ROMANTIC MOVIES AND "
karkatDef += "ROMCOMS. Should really be EMBARRASSED for liking this DREADFUL "
karkatDef += "CINEMA, but for some reason is not."

defs := [][2]string{
	{"Aradia", aradiaDef},
	{"Tavros", tavDef},
	{"Sollux", solluxDef},
	{"Karkat", karkatDef},
}

const (
	position = 0
	width    = 80
)

defTable := rosed.Edit("").InsertDefinitionsTable(position, defs, width).String()

fmt.Println(defTable)
```

*Output:*

```
  Aradia  - Once had a fondness for ARCHEOLOGY, though now has trouble recalling
            this passion.

  Tavros  - Known to be heavily arrested by FAIRY TALES AND FANTASY STORIES. Has
            an acute ability to COMMUNE WITH THE MANY CREATURES OF ALTERNIA.

  Sollux  - Is good at computers, and knows ALL THE CODES. All of them.

  Karkat  - Has a passion for RIDICULOUSLY TERRIBLE ROMANTIC MOVIES AND ROMCOMS.
            Should really be EMBARRASSED for liking this DREADFUL CINEMA, but
            for some reason is not.
```

### Two-Column Layout

You can use this library to create two columns of text:

*[Try it on the Go Playround](https://go.dev/play/p/uITIF04nz22)*

```golang
left := "You have a feeling it's going to be a long day."

right := "A familiar note is produced. It's the one Desolation plays to "
right += "keep its instrument in tune."

const (
	position    = 0
	middleSpace = 2
	width       = 60
	leftPercent = 0.33
)

columns := rosed.Edit("").
	InsertTwoColumns(position, left, right, middleSpace, width, leftPercent).
	String()

fmt.Println(columns)
```

Output:

```
You have a feeling   A familiar not is produced. It's the
it's going to be a   one Desolation plays to keep its
long day.            instrument in tune.
```

### Custom Functionality
Do you like the idea of operating on lines with customizability but rosed's
built-in Editor functions aren't enough for you? Use Apply to give your own
function to run on each line, or ApplyParagraphs to run it on each paragraph.

For instance, to upper-case every line while giving each a line number:

*[Try it on the Go Playground](https://go.dev/play/p/DDxzVyvBxph)*

```golang
text := "She went too far this time and she knows it.\n"
text += "She's got to pay.\n"
text += "Justice is long overdue.\n"

lineNumberMaker := func(idx int, line string) []string {
	upperLine := strings.ToUpper(line)
	line = fmt.Sprintf("LINE %d: %s", idx+1, upperLine)
	return []string{
		line,
	}
}

customized := rosed.Edit(text).Apply(lineNumberMaker).String()

fmt.Println(customized)
```

Output:

```
LINE 1: SHE WENT TOO FAR THIS TIME AND SHE KNOWS IT.
LINE 2: SHE'S GOT TO PAY.
LINE 3: JUSTICE IS LONG OVERDUE.
```

### Sub-Section Selection
You can open a sub-editor on a sub-section of text in an Editor to only affect
that part of the text.

What if we forgot to tab some lines of an indented paragraph that comes after a
non-indented paragraph? Oh no! We better fix that, quick:

*[Try it on the Go Playground](https://go.dev/play/p/zxhp_t6ghTn)*
```golang
text := "Your name is NEPETA LEIJON.\n"
text += "\n"
text += "\tYou live in a CAVE that is also a HIVE.\n"
text += "But still mostly just a cave.\n"                 // this line should have an indent in it
text += "You like to engage in FRIENDLY ROLE PLAYING.\n"  // so should this one
text += "\tBut not the DANGEROUS KIND.\n"
text += "\tIt's TOO DANGEROUS!"

corrected := rosed.Edit(text).Lines(3, 5).Indent(1).String()

fmt.Println(corrected)
```

Output:

```
Your name is NEPETA LEIJON.

	You live in a CAVE that is also a HIVE.
	But still mostly just a cave.
	You like to engage in FRIENDLY ROLE PLAYING.
	But not the DANGEROUS KIND.
	It's TOO DANGEROUS!
```

## Contributing
This library uses its [GitHub Issues Page](https://github.com/dekarrin/rosed/issues)
for coordination of work. If you'd like to assist with an open issue, feel free
to drop a comment on the issue mentioning your interest. Then, fork the repo
into your own copy, make the changes, then raise a PR to bring them in.

If there's a bug you'd like to report or new feature you'd like to see in this
library, feel free to [open a new issue for it](https://github.com/dekarrin/rosed/issues/new).

## Example Text Source
This explains the refrance [sic] of the sample text, most of the samples,
test cases, and example texts used in this library's documentation have been
sampled from the text of [Homestuck](https://www.homestuck.com/story/1), a
primarily text-based web comic with occasional animations whose topic matter
ranges from data-structure jokes to an analysis of the self and how our past
makes us. Sadly due to Flash dying a lot of the site is non-functional.

May it live on in our hearts and naming conventions longer than the darkness of
the furthest ring. Or, you know, just generally exist in [The Unofficial
Homestuck Collection](https://bambosh.dev/unofficial-homestuck-collection/).
