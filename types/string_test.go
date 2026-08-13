package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringToUTF16(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		utf16 := StringToUTF16("abc")

		expected := []byte{0x61, 0x00, 0x62, 0x00, 0x63, 0x00, 0x00, 0x00}
		require.Equal(t, expected, utf16)
	})

	t.Run("empty", func(t *testing.T) {
		utf16 := StringToUTF16("")
		require.Zero(t, utf16)
	})
}

func TestUTF16ToString(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		utf16 := []byte{0x61, 0x00, 0x62, 0x00, 0x63, 0x00, 0x00, 0x00}

		s, err := UTF16ToString(utf16)
		require.NoError(t, err)

		expected := "abc"
		require.Equal(t, expected, s)
	})

	t.Run("invalid", func(t *testing.T) {
		utf16 := []byte{0x61}

		s, err := UTF16ToString(utf16)
		require.EqualError(t, err, "invalid utf16 string")
		require.Zero(t, s)
	})
}
