package dd

import (
	"github.com/parnurzeal/gorequest"
	"github.com/go-playground/form"
	"github.com/json-iterator/go"
	"errors"
	"fmt"
)

// 主要处理Client
// 请使用单例模式
type Client struct {
	corpID        string
	corpSecret    string
	cacheToken    Token
	cacheJSTicket JSTicket
	formDecoder   *form.Decoder
	formEncode    *form.Encoder
}

func NewClient(corpID, corpSecret string) (client *Client, err error) {
	client = &Client{
		corpID:      corpID,
		corpSecret:  corpSecret,
		formDecoder: form.NewDecoder(),
		formEncode:  form.NewEncoder(),
	}
	return
}

type CommonReply struct {
	ErrCode int    `json:"errcode"` // 错误码
	ErrMsg  string `json:"errmsg"`  // 错误信息
}

func (rep CommonReply) GetErrorReply() (code int, msg string) {
	return rep.ErrCode, rep.ErrMsg
}

type ICommonReply interface {
	GetErrorReply() (code int, msg string)
}

func (c *Client) getJSON(url string, params, out interface{}) error {
	var err error
	if params != nil {
		if vals, err1 := c.formEncode.Encode(params); err1 != nil {
			return err1
		} else {
			url += "?" + vals.Encode()
		}
	}
	gorequest.New().Get(url).End(func(response gorequest.Response, body string, errs []error) {
		if len(errs) > 0 {
			err = errs[0]
			return
		}
		err = jsoniter.UnmarshalFromString(body, out)
		if err == nil {
			ec, em := out.(ICommonReply).GetErrorReply()
			if ec > 0 {
				out = nil
				err = errors.New(fmt.Sprintf("code: %d, msg: %s", ec, em))
				return
			}
		}
	})
	return err
}
