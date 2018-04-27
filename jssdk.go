package dd

import (
	"crypto/sha1"
	"github.com/rs/xid"
	"strconv"
	"encoding/hex"
)

type JSSign struct {
	Url       string `json:"url" form:"url"`             // 当前网页的URL，不包含#及其后面部分
	NonceStr  string `json:"nonceStr" form:"nonceStr"`   // 随机串，自己定义
	AgentId   string `json:"agentId" form:"agentId"`     // 应用的标识	编辑企业应用可以看到
	TimeStamp int64  `json:"timeStamp" form:"timeStamp"` //时间戳	当前时间，但是前端和服务端进行校验时候的值要一致
	CorpId    string `json:"corpId" form:"corpId"`       //企业ID	企业ID，在//open-dev.dingtalk.com/上企业视图下开发者账号设置里面可以看到
	Signature string `json:"signature" form:"signature"`
}

type JSTicket struct {
	CommonReply
	ExpiresIn int    `json:"expires_in"` // 过期时间
	Ticket    string `json:"ticket"`     // 获取到的凭证
}

// 获取jsticket
// 默认缓存在client里
// TODO 过期刷新
func (c *Client) GetJSTicket() (t JSTicket, err error) {
	err = c.getJSON(DD_API_HOST+"/get_jsapi_ticket?type=jsapi&access_token="+c.cacheToken.AccessToken, nil, &t)
	c.cacheJSTicket = t
	return
}

// 获取js签名
// ticket为空时，使用缓存的ticket
func (c *Client) GetJSSign(url, agentId, ticket string) JSSign {
	if ticket == "" {
		ticket = c.cacheJSTicket.Ticket
	}
	id := xid.New()
	sig := JSSign{
		Url:       url,
		NonceStr:  id.String(),
		AgentId:   agentId,
		TimeStamp: id.Time().Unix(),
		CorpId:    c.corpID,
	}
	plain := "jsapi_ticket=" + c.cacheJSTicket.Ticket + "&noncestr=" + sig.NonceStr +
		"&timestamp=" + strconv.FormatInt(sig.TimeStamp, 10) + "&url=" + sig.Url
	hashsum := sha1.Sum([]byte(plain))
	sig.Signature = hex.EncodeToString(hashsum[:])
	return sig
}
