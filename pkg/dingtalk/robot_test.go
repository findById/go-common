package dingtalk

import (
	"log"
	"testing"
)

func TestRobot_SendTextMessage(t *testing.T) {
	token := ""
	secret := ""

	data := make(map[string]string, 0)
	data["content"] = "Hello, 世界！"
	data["at"] = "all"

	robot := NewRobot(token, secret, true)
	err := robot.SendTextMessage(data)
	if err != nil {
		log.Println(err)
	}
}
