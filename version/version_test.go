package version

import (
	"fmt"
	"strconv"
	"testing"
)

func TestParseVersion(t *testing.T) {
	fmt.Println(ParseVersionToArray("1.12."))
	fmt.Println(ParseVersion("1.12"))
	fmt.Println(ToVersionString(ParseVersion("1.12")))
	v1 := ParseVersion("1.12.123")
	vv, _ := strconv.Atoi(v1)
	fmt.Println(v1, ToVersionString(fmt.Sprint(vv)))
	fmt.Println(ToVersionString(ParseVersion("1.1.1")))
}
