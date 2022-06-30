package handler

import "gedis/src/protocol"

func HandleMultiBulkRequest(key string, elements []string) [][]byte {
	bytes := make([][]byte, len(elements)+1)
	bytes[0] = protocol.SafeEncode(key)
	for i := 0; i < len(elements); i++ {
		bytes[i+1] = protocol.SafeEncode(elements[i])
	}
	return bytes
}

func HandleBulkRequest(elements []string) [][]byte {
	bytes := make([][]byte, len(elements))
	for i := 0; i < len(elements); i++ {
		bytes[i] = protocol.SafeEncode(elements[i])
	}
	return bytes
}
