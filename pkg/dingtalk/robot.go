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
	token  string
	secret string
	debug  bool
}

func NewRobot(token, secret string, debug bool) *Robot {
	return &Robot{
		token:  token,
		secret: secret,
		debug:  debug,
	}
}

func (robot *Robot) hmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (robot *Robot) SendRobotMessage(message string) error {
	var url string
	if robot.secret != "" {
		ts := time.Now().UnixNano() / 1e6
		sign := robot.hmacSha256(fmt.Sprint(ts, "\n", robot.secret), robot.secret)
		temp := fmt.Sprintf("&timestamp=%v&sign=%v", ts, sign)
		url = fmt.Sprint(RobotSendAPI, robot.token, temp)
	} else {
		url = fmt.Sprint(RobotSendAPI, robot.token)
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
	fmt.Println()
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
	} else if at != "" {
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
	} else {
		msg = `{
			"msgtype":"text",
			"text":{
				"content":"%v"
			}
		}`
		msg = fmt.Sprintf(msg, content)
	}
	if !robot.debug {
		err := robot.SendRobotMessage(msg)
		if err != nil {
			return err
		}
	} else {
		return errors.New(msg)
	}
	return nil
}

func (robot *Robot) SendMarkdownMessage(data map[string]string) error {
	title := data["title"]
	content := data["content"]
	at := data["at"]

	var msg string
	if at == "true" || at == "all" {
		msg = `{
			"msgtype":"markdown",
			"title":"%v",
			"text":{
				"content":"%v"
			},
			"at":{
				"isAtAll":%v
			}
		}`
		msg = fmt.Sprintf(msg, title, content, true)
	} else if at != "" {
		msg = `{
			"msgtype":"markdown",
			"title":"%v",
			"text":{
				"content":"%v"
			},
			"at":{
				"atMobiles":%v,
				"isAtAll":false
			}
		}`
		mobiles := fmt.Sprintf(`["%v"]`, strings.Join(strings.Split(at, ","), `","`)) // (1,1) to (["1","1"])
		msg = fmt.Sprintf(msg, title, content, mobiles)
	} else {
		msg = `{
			"msgtype":"markdown",
			"title":"%v",
			"text":{
				"content":"%v"
			}
		}`
		msg = fmt.Sprintf(msg, title, content)
	}
	if !robot.debug {
		err := robot.SendRobotMessage(msg)
		if err != nil {
			return err
		}
	} else {
		return errors.New(msg)
	}
	return nil
}
