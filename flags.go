package flags

// Bits represents 64 bits from 0 .. 63.
type Bits uint64

// Flag constants are used to access individual bits from Flag0 .. Flag63.
const (
	Flag0 Bits = 1 << iota
	Flag1
	Flag2
	Flag3
	Flag4
	Flag5
	Flag6
	Flag7
	Flag8
	Flag9
	Flag10
	Flag11
	Flag12
	Flag13
	Flag14
	Flag15
	Flag16
	Flag17
	Flag18
	Flag19
	Flag20
	Flag21
	Flag22
	Flag23
	Flag24
	Flag25
	Flag26
	Flag27
	Flag28
	Flag29
	Flag30
	Flag31
	Flag32
	Flag33
	Flag34
	Flag35
	Flag36
	Flag37
	Flag38
	Flag39
	Flag40
	Flag41
	Flag42
	Flag43
	Flag44
	Flag45
	Flag46
	Flag47
	Flag48
	Flag49
	Flag50
	Flag51
	Flag52
	Flag53
	Flag54
	Flag55
	Flag56
	Flag57
	Flag58
	Flag59
	Flag60
	Flag61
	Flag62
	Flag63
)

// Set the flag in the given Bits field, return the new Bits field.
func Set(b, flag Bits) Bits { return b | flag }

// Clear the flag in the given Bits field, return the new Bits field.
func Clear(b, flag Bits) Bits { return b &^ flag }

// Toggle the flag in the given Bits field, return the new Bits field.
func Toggle(b, flag Bits) Bits { return b ^ flag }

// Has returns true if the flag is set in the given Bits field, false otherwise.
func Has(b, flag Bits) bool { return (b & flag) > 0 }
