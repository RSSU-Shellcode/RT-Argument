package serialization

// serialized data structure
// +---------+----------+----------+----------+------------+
// |  magic  |  item 1  |  item 2  | item END |  raw data  |
// +---------+----------+----------+----------+------------+
// |  uint32 |  uint32  |  uint32  |  uint32  |    var     |
// +---------+----------+----------+----------+------------+
//
// item data structure
// 0······· value or pointer
// ·0000000 data length

const (
	magic   = 0xACFFFFEE
	itemEnd = 0x00000000

	maskType   = 0x80000000
	maskLength = 0x7FFFFFFF

	typeValue   = 0x00000000
	typePointer = 0x80000000
)
