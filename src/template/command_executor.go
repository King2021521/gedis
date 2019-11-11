package template

import (
	"net"
	"bytes"
	"protocol"
	"strconv"
	"fmt"
	"os"
	"strings"
	"github.com/emirpasic/gods/lists/arraylist"
)

func SendCommand(conn *net.TCPConn, cmd string, a ...[]byte) string {
	var buffer bytes.Buffer
	buffer.Write(protocol.SafeEncode(protocol.ASTERISKBYTE))
	buffer.Write(protocol.SafeEncode(strconv.Itoa(len(a) + 1)))
	buffer.Write(protocol.SafeEncode(protocol.CRLF))
	buffer.Write(protocol.SafeEncode(protocol.DOLLARBYTE))
	buffer.Write(protocol.SafeEncode(strconv.Itoa(len(cmd))))
	buffer.Write(protocol.SafeEncode(protocol.CRLF))
	buffer.Write(protocol.SafeEncode(cmd))
	buffer.Write(protocol.SafeEncode(protocol.CRLF))

	for _, arg := range a {
		buffer.Write(protocol.SafeEncode(protocol.DOLLARBYTE))
		buffer.Write(protocol.SafeEncode(strconv.Itoa(len(arg))))
		buffer.Write(protocol.SafeEncode(protocol.CRLF))
		buffer.Write(arg)
		buffer.Write(protocol.SafeEncode(protocol.CRLF))
	}
	return send(conn, buffer)
}

func send(conn *net.TCPConn, content bytes.Buffer) string {
	//send to server
	_, err := conn.Write(content.Bytes())

	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), "server response")
		os.Exit(1)
	}

	buffer := make([]byte, 1024)
	//receive server info
	msg, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), "server response:"+err.Error())
		os.Exit(1)
	}
	return string(buffer[:msg])
}

func HandleComplexResult(result string) interface{} {
	array := strings.Split(result, protocol.CRLF)
	results := arraylist.New()
	for i := 1; i < len(array)-1; i++ {
		if array[i] == protocol.NONEXIST {
			results.Add(nil)
			continue
		}
		results.Add(array[i+1])
		i++
		if i > len(array)-2 {
			break
		}
	}
	return results.Values()
}
