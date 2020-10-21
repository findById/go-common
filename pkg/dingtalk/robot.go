// https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq/uKPlK

package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	RobotSendAPI = "https://oapi.dingtalk.com/robot/send?access_token="
)

type Robot struct {
	Token  string
	Secret string
	debug  bool
}

func NewRobot(token, secret string, debug bool) *Robot {
	return &Robot{
		Token:  token,
		Secret: secret,
		debug:  debug,
	}
}

func hmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (robot *Robot) sendRobotMessage(message string) error {
	var url string
	if robot.Secret != "" {
		ts := time.Now().UnixNano() / 1e6
		sign := hmacSha256(fmt.Sprint(ts, "\n", robot.Secret), robot.Secret)
		temp := fmt.Sprintf("&timestamp=%v&sign=%v", ts, sign)
		url = fmt.Sprint(RobotSendAPI, robot.Token, temp)
	} else {
		url = fmt.Sprint(RobotSendAPI, robot.Token)
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(message))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	_, err = io.Copy(os.Stderr, res.Body)
	return err
}

func (robot *Robot) SendTextMessage(data map[string]string) error {
	content := data["content"]
	at := data["at"]

	var msg string
	if at == "true" || at == "all" {
		msg = `{
			"msgtype":"text",
			"text":{
				"content":"%v"
			},
			"at":{
				"isAtAll":%v
			}
		}`
		msg = fmt.Sprintf(msg, content, true)
	} else {
		msg = `{
			"msgtype":"text",
			"text":{
				"content":"%v"
			},
			"at":{
				"atMobiles":%v,
				"isAtAll":false
			}
		}`
		mobiles := fmt.Sprintf(`["%v"]`, strings.Join(strings.Split(at, ","), `","`)) // (1,1) to (["1","1"])
		msg = fmt.Sprintf(msg, content, mobiles)
	}
	if !robot.debug {
		err := robot.sendRobotMessage(msg)
		if err != nil {
			return err
		}
	} else {
		return errors.New(msg)
	}
	return nil
}
