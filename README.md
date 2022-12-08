rosed
=====

![Tests Status Badge](https://github.com/dekarrin/rosed/actions/workflows/tests.yml/badge.svg?branch=main&event=push)
[![Go Reference](https://pkg.go.dev/badge/github.com/dekarrin/rosed.svg)](https://pkg.go.dev/github.com/dekarrin/rosed)

Pronounced "ROSE-edd". Fluent text editing and layout library for Go. Supports
concept of "graphemes" (user-visible characters) vs actual runes. This allows
wrapping and other character-counting/index operations to function in an
expected way regardless of whether visible characters are made up of Unicode
decomposed sequences or pre-composed characters. Basically, it's Unicode aware.

Take something like "Ý", which appears to be one single character but behind the
scenes could be a single precomposed character rune or multiple runes that
compose into a single apparent character using Unicode combining marks. This
library will treat both as the same, assuming it is used correctly.

This is an early pre-release version of the library. Version 1.0.0 may feature
code changes, although since the main thing that still needs to be done is
adding proper documentation, this version has been published in the meantime.

TODO: examples of what this lib can do

### What This Library Is Not

* Performant - Especially for initial release the idea is to create a library
that is easy to use. Better performance may come in later versions but that is
not a goal at the time.
* Collating - This libarary does not support Unicode collation. It may come at
a later date.
