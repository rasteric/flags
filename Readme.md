 # Flags
*-- get, set, toggle, or clear 64 flags

[![GoDoc](https://godoc.org/github.com/rasteric/flags/go?status.svg)](https://godoc.org/github.com/rasteric/flags)
[![Go Report Card](https://goreportcard.com/badge/github.com/rasteric/flags)](https://goreportcard.com/report/github.com/rasteric/flags)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

Flags is a simple package to set and get 64 bit flags. If you're familiar with basic bit twiddling you won't need it, but if you're like me and need to look up these operations every time because you're not using them very often, this package might be a useful shortcut.

## API

`type Bits uint64`

Bits represents 64 bits. The constants Flag0 .. Flag63 can be used to address the flags in this bitfield.

`Set(b, flag Bits) Bits`

Set the flag in the given bitfield `b`, return the new bitfield.

`Clear(b, flag Bits) Bits`

Clear the flag in the given bitfield `b`, return the new bitfield.

`Toggle(b, flag Bits) Bits`

Toggle the flag in the given bitfield `b`, return the new bitfield.

`Has(b, flag Bits) bool`

Return true if the flag in bitfield `b` is set, false otherwise.




