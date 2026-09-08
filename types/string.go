package types

import (
	"encoding/binary"
	"errors"
	"unicode/utf16"
)

// StringToUTF16 is used to convert go string to utf16 string.
func StringToUTF16(s string) []byte {
	if s == "" {
		return nil
	}
	w := utf16.Encode([]rune(s))
	output := make([]byte, (len(w)+1)*2)
	for i := 0; i < len(w); i++ {
		binary.LittleEndian.PutUint16(output[i*2:], w[i])
	}
	return output
}

// UTF16ToString is used to convert utf16 string to go string.
func UTF16ToString(b []byte) (string, error) {
	n := len(b)
	if n == 0 {
		return "", nil
	}
	n -= 2 // remove the "0x00, 0x00" at tail
	if n%2 != 0 {
		return "", errors.New("invalid utf16 string")
	}
	u16 := make([]uint16, n/2)
	for i := 0; i < len(u16); i++ {
		u16[i] = binary.LittleEndian.Uint16(b[i*2:])
	}
	return string(utf16.Decode(u16)), nil
}
