package wechatwork

import (
	"encoding/json"
	"fmt"
	"github.com/chamzzzzzz/wechatwork/user"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	defaultURL = "https://qyapi.weixin.qq.com/cgi-bin/"
)

type Client struct {
	CorpId     string
	CorpSecret string
	Token      string
	ExpireAt   int64
}

func (c *Client) ResetToken(force bool) error {
	type Body struct {
		Errcode     int64  `json:"errcode"`
		Errmsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}

	now := time.Now().Unix()
	if now < c.ExpireAt && !force {
		return nil
	}

	res, err := http.Get(fmt.Sprintf("%sgettoken?corpid=%s&corpsecret=%s", defaultURL, c.CorpId, c.CorpSecret))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bodyData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	body := Body{}
	if err := json.Unmarshal(bodyData, &body); err != nil {
		return err
	}

	if body.Errcode != 0 {
		return fmt.Errorf("%d-%s", body.Errcode, body.Errmsg)
	}

	c.Token = body.AccessToken
	c.ExpireAt = time.Now().Unix() + body.ExpiresIn
	return nil
}

func (c *Client) GetDepartmentUsers(departmentId int64) ([]*user.User, error) {
	type Body struct {
		Errcode  int64        `json:"errcode"`
		Errmsg   string       `json:"errmsg"`
		Userlist []*user.User `json:"userlist"`
	}

	err := c.ResetToken(false)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(fmt.Sprintf("%suser/list?access_token=%s&department_id=%d&fetch_child=1", defaultURL, c.Token, departmentId))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	body := Body{}
	if err := json.Unmarshal(bodyData, &body); err != nil {
		return nil, err
	}

	if body.Errcode != 0 {
		return nil, fmt.Errorf("%d-%s", body.Errcode, body.Errmsg)
	}

	return body.Userlist, nil
}
