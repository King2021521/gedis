package protocol

import "strconv"

func SafeEncode(arg string) []byte {
	return []byte(arg)
}

func SafeEncodeInt(arg int64) []byte {
	return SafeEncode(strconv.FormatInt(arg, 10))
}
