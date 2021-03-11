// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html

package wxmp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	ApiGetTokenFormat    = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v"
	ApiMessageSendFormat = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%v"
)

type MiniSubscribeMessage struct {
	appId  string
	secret string
	token  string
	retry  int
}

func NewMiniSubscribeMessage(appId, secret string) *MiniSubscribeMessage {
	return &MiniSubscribeMessage{
		appId:  appId,
		secret: secret,
		retry:  2,
	}
}

func (m *MiniSubscribeMessage) Send(msg string) error {
	url := fmt.Sprintf(ApiMessageSendFormat, m.token)

	res, err := m.sendMessage(url, msg)
	if err != nil {
		return err
	}
	temp, err := m.parseResp([]byte(res))
	if err != nil {
		if temp != nil && (fmt.Sprint(temp["errcode"]) == "41001" || fmt.Sprint(temp["errcode"]) == "42001") {
			m.retry--
			if m.retry > 0 && m.getAccessToken() == nil {
				return m.Send(msg)
			}
		}
		return err
	}
	log.Println(res)
	return nil
}

func (m *MiniSubscribeMessage) GetAccessToken() error {
	err := m.getAccessToken()
	log.Println(m.token)
	return err
}

func (m *MiniSubscribeMessage) getAccessToken() error {
	url := fmt.Sprintf(ApiGetTokenFormat, m.appId, m.secret)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	temp, err := m.parseResp(b)
	if err != nil {
		return err
	}
	m.token = fmt.Sprint(temp["access_token"])
	return nil
}

func (m *MiniSubscribeMessage) sendMessage(url, msg string) (string, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(msg))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (m *MiniSubscribeMessage) parseResp(b []byte) (map[string]interface{}, error) {
	temp := make(map[string]interface{}, 0)
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return nil, err
	}
	if temp["errcode"] == nil || fmt.Sprint(temp["errcode"]) == "<nil>" {
		return temp, nil
	}
	return temp, errors.New(string(b))
}
