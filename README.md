rosed
=====

![Tests Status Badge](https://github.com/dekarrin/rosed/actions/workflows/tests.yml/badge.svg?branch=main&event=push)
[![Go Reference](https://pkg.go.dev/badge/github.com/dekarrin/rosed.svg)](https://pkg.go.dev/github.com/dekarrin/rosed)

Pronounced "ROSE-edd". Fluent text editing and layout library for Go. Supports
concept of "graphemes" (user-visible characters) vs actual runes. This allows
wrapping and other character-counting/index operations to function in an
expected way regardless of whether visible characters are made up of Unicode
decomposed sequences or pre-composed characters. Basically, it's Unicode aware.

### Graphemes? What Are Those?

Take something like "Ý", which appears to be one single character but behind the
scenes could be a single precomposed character rune or multiple runes that
compose into a single apparent character using Unicode combining marks. This
library will treat both as the same, assuming it is used correctly.

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

In general, to use rosed, first create an Editor, use it to modify text, the
get it back by converting it to a string. As easy as plucking a rose from the
bed of thorns that grows in the midst of the most eldritch of gardens. Perhaps
easier, in fact.

### Wrap Text

You can use rosed to wrap text to a certain width:

*Try it on the [Go Playground](https://go.dev/play/p/XdoblhFv3XX)*

```golang
package main

import (
	"fmt"
	
	"github.com/dekarrin/rosed"
)

func main() {
	text := "It's hard, being a kid and growing up. It's hard and nobody understands."
	
	wrapped := rosed.Edit(text).Wrap(30).String()
	
	fmt.Println(wrapped)
	// It's hard, being a kid and
	// growing up. It's hard and
	// nobody understands.
}
```

You can do this even if it's already been hard-wrapped:

*Try it on the [Go Playground](https://go.dev/play/p/6O7jjoft1Qr)*

```golang
package main

import (
	"fmt"
	
	"github.com/dekarrin/rosed"
)

func main() {
	text := "It's hard, being\na kid and growing up.\nIt's hard and nobody\nunderstands."
	
	wrapped := rosed.Edit(text).Wrap(30).String()
	
	fmt.Println(wrapped)
	// It's hard, being a kid and
	// growing up. It's hard and
	// nobody understands.
}
```

### Justify Text

You can use rosed to ensure that all text takes up the same width of lines, and
add spacing equally to ensure this is the case, also known as justifying the
block of text:

*Try it on the [Go Playground](https://go.dev/play/p/3bNywhvrFch)*

```golang
package main

import (
	"fmt"
	
	"github.com/dekarrin/rosed"
)

func main() {
	text := "I WARNED YOU ABOUT STAIRS\n"
	text += "BRO! I TOLD YOU DOG!\n"
	text += "I TOLD YOU MAN!"
	
	justified := rosed.Edit(text).Justify(30).String()
	
	fmt.Println(justified)
	// I   WARNED  YOU  ABOUT  STAIRS
	// BRO!    I   TOLD   YOU    DOG!
	// I      TOLD      YOU      MAN!
}
```

You could combine this with wrapping to account for the lines that didn't quite
end up the same length as the others:

*Try it on the [Go Playground](https://go.dev/play/p/P6KievkJ_cQ)*

```golang
package main

import (
	"fmt"
	
	"github.com/dekarrin/rosed"
)

func main() {
	text := "Your name is KANAYA MARYAM. You are one of the few of your "
	text += "kind who can withstand the BLISTERING ALTERNIAN SUN, and "
	text += "perhaps the only who enjoys the feel of its rays. As such, "
	text += "you are one of the few of your kind who has taken a shining "
	text += "to LANDSCAPING. You have cultivated a lush oasis."
	
	wrappedAndJustified := rosed.Edit(text).Wrap(30).Justify(30).String()
	
	fmt.Println(wrappedAndJustified)
}
```

TODO: before 1.0, make all output like this (? after checking the output!!!!!!!!).
This also means that we need to upd8 the Playground links 8ecause their code
won't be right.

Output:
```
Your  name is  KANAYA  MARYAM.
You are one of the few of your
kind  who  can  withstand  the
BLISTERING  ALTERNIAN SUN, and
perhaps  the only  who  enjoys
the feel of its rays. As such,
you are one of the few of your
kind  who has taken  a shining
to   LANDSCAPING.   You   have
cultivated   a   lush   oasis.
```

### Definitions Table

### Two-Column Layout

### Sub-Section Selection

## Contributing
This library uses its [GitHub Issues Page](https://github.com/dekarrin/rosed/issues)
for coordination of work. If you'd like to assist with an open issue, feel free
to drop a comment on the issue mentioning it. Fork the repo into your own, make
the changes, then raise a PR to bring them in.

If there's a bug you'd like to report or new feature you'd like to see in this
library, feel free to open a new issue on its [GitHub Issues Page](https://github.com/dekarrin/rosed/issues).

## Example Text Source
In case you didn't quite comprehend the refrance [sic], most of the samples,
test cases, and example texts used in this library's documentation have been
from the text of [Homestuck](https://www.homestuck.com/story/1), a primarily
text-based web comic with occasional animations whose topic matter ranges from
data-structure jokes to an analysis of the self and how our past makes us.

The name of the library also happens to be a reference to one of the characters,
Rose Lalonde, as part of the never-ending drive to name things after entities
from favored media.

The comic is okay. If you want to read it, due to its reliance on Flash, the
best way to do so at the time of this writing is by using [The Unofficial
Homestuck Collection](https://bambosh.dev/unofficial-homestuck-collection/).