rosed
=====

![Tests Status Badge](https://github.com/dekarrin/rosed/actions/workflows/tests.yml/badge.svg?branch=main&event=push)

Fluent text editing and layout library for Go. Supports concept of "graphemes" (user-visible characters) vs actual
runes.

Take something like "Ý", which appears to be one single character but behind the scenes could be a single
precomposed character rune or multiple runes that compose into a single apparent character using Unicode
combining marks. This library will treat both as the same, assuming it is used correctly.

Note that it is entirely possible to use it incorrectly.

This is an early pre-release version of the library. Version 1.0.0 may feature code changes, although since
the main thing that still needs to be done is adding proper documentation, this version has been published
in the mean time.
