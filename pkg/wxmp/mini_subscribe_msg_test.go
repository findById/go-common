package wxmp

import (
	"fmt"
	"testing"
)

func TestMiniSubscribeMessage_Send(t *testing.T) {
	msg := `
{
  "touser": "",
  "template_id": "",
  "page": "pages/index/index",
  "data": {"thing1": {"value": "xxx"},"thing2": {"value": "xxx"},"time3": {"value": "2020-02-12 16:30"} ,"thing4": {"value": "xxx"}}
}`

	mini := NewMiniSubscribeMessage("", "")
	fmt.Println(mini.Send(msg))
}
