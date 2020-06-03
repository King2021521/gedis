package handler

import (
	"strings"
	"protocol"
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
)

/**
 * 处理redis响应的结果
 * 1、复杂的*...
 * 2、:数字
 * 3、简单的$...
 * 4、+OK
 * 5、-ERR(WRONGTYPE...)
 * 6、+QUEUED
 */
func HandleReply(result string) (interface{}, error) {
	if strings.HasPrefix(result, protocol.MOVED) {
		return handleMovedReply(result)
	}

	if strings.HasPrefix(result, protocol.MINUS_BYTE) {
		return handleMinusReply(result)
	}

	if strings.HasPrefix(result, protocol.PLUSBYTE) {
		return handlePlusReply(result)
	}

	if strings.HasPrefix(result, protocol.COLON_BYTE) {
		return handleColonReply(result)
	}

	if strings.HasPrefix(result, protocol.DOLLARBYTE) {
		return handleDollarReply(result)
	}

	if strings.HasPrefix(result, protocol.ASTERISKBYTE) {
		return HandleAsteriskReply(result)
	}
	return nil, fmt.Errorf("reply handle err")
}

func HandleAsteriskReply(result string) (interface{}, error) {
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
	return results.Values(), nil
}

func handlePlusReply(result string) (interface{}, error) {
	if result != protocol.OK && !strings.HasPrefix(result, protocol.PLUSBYTE+protocol.PONG) && !strings.HasPrefix(result, protocol.PLUSBYTE+protocol.QUEUED) {
		return nil, fmt.Errorf(result)
	}
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.PLUSBYTE, protocol.BLANK), nil
}

func handleDollarReply(result string) (interface{}, error) {
	if strings.HasPrefix(result, protocol.NONEXIST) {
		return nil, nil
	}

	if !strings.HasPrefix(result, protocol.DOLLARBYTE) {
		return nil, fmt.Errorf(result)
	}
	array := strings.Split(result, protocol.CRLF)
	return array[1], nil
}

func handleColonReply(result string) (interface{}, error) {
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK), nil
}

func handleMinusReply(result string) (interface{}, error) {
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.MINUS_BYTE, protocol.BLANK), nil
}

//访问集群模式时，处理响应的MOVED指令
func handleMovedReply(result string) (interface{}, error) {
	movedInfo := strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.MINUS_BYTE, protocol.BLANK)
	array := strings.Split(movedInfo, " ")
	return array[2], fmt.Errorf(protocol.MOVED)
}

//处理事务执行结果
func HandleTransactionReply(result string)(interface{}, error){
	elements := strings.Split(result, protocol.CRLF)
	var values []string
	//解析事务批量执行的返回结果
	for i := 1; i < len(elements); {
		if elements[i] == protocol.BLANK {
			i++
			continue
		}

		if strings.HasPrefix(elements[i], protocol.DOLLARBYTE) {
			values = append(values, elements[i+1])
			i += 2
		} else {
			values = append(values, elements[i])
			i++
		}
	}
	return values, nil
}
