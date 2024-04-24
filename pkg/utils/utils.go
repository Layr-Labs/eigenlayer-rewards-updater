package utils

import (
	"encoding/hex"
)

func ConvertBytesToStrings(b [][]byte) []string {
	var s []string
	for _, v := range b {
		s = append(s, "0x"+hex.EncodeToString(v[:]))
	}
	return s
}

func ConvertBytesToString(b []byte) string {
	return "0x" + hex.EncodeToString(b[:])
}

func ConvertBytes32ToString(b [32]byte) string {
	return "0x" + hex.EncodeToString(b[:])
}
