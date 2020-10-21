package version

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Count = 3 // major.minor.revision
	Len   = 3 // 000.000.000
)

func CompareVersion(source, target, operation string) bool {
	left, err := ParseVersionToInt(source)
	if err != nil {
		return false
	}
	right, err := ParseVersionToInt(target)
	if err != nil {
		return false
	}
	if operation == ">" {
		return left > right
	} else if operation == "<" {
		return left < right
	} else if operation == "==" {
		return left == right
	} else if operation == ">=" {
		return left >= right
	} else if operation == "<=" {
		return left <= right
	} else if operation == "!=" {
		return left != right
	}
	return false
}

// 反序列化为可读版本  1000001 >> 1.0.1
func ToVersionString(version string) string {
	version = appendBit(version, Count*Len) // 补位
	ver := make([]string, 0)
	for i := 0; i < len(version); i += Count {
		item := ""
		for j := 0; j < Len; j++ {
			item += version[i+j : i+j+1]
		}
		ver = append(ver, item)
	}

	str := ""
	for i := 0; i < Count; i++ {
		temp, err := strconv.Atoi(ver[i])
		if err != nil {
			temp = 0
		}
		if i > 0 && i < Count {
			str += "."
		}
		str += fmt.Sprint(temp)
	}
	return str
}

// 序列化为存库版本 1.0.1 >> 1000001
func ParseVersion(version string) string {
	ver := strings.Split(version, ".")
	ver = appendArray(ver, Count)

	str := ""
	for i := 0; i < Count; i++ {
		str += appendBit(ver[i], Len)
	}
	return str
}

func ParseVersionToInt(version string) (int, error) {
	v := ParseVersion(version)
	return strconv.Atoi(v)
}

// 序列化为存库版本 数组形式
func ParseVersionToArray(version string) []string {
	ver := strings.Split(version, ".")
	ver = appendArray(ver, Count)

	data := make([]string, Count)
	for i := 0; i < Count; i++ {
		data[i] = appendBit(ver[i], Len)
	}
	return data
}

// 版本号段补位 length=3 major.minor.revision
func appendArray(array []string, length int) []string {
	size := length - len(array)
	if size <= 0 {
		return array
	}
	for i := 0; i < size; i++ {
		array = append(array, "")
	}
	return array
}

// 版本号补位 length=2 00.00.00
func appendBit(str string, length int) string {
	size := length - len(str)
	if size <= 0 {
		return str
	}
	temp := ""
	for i := 0; i < size; i++ {
		temp += "0"
	}
	return temp + str
}
