package dd

import (
	"fmt"
)

type Token struct {
	CommonReply
	ExpiresIn   int    `json:"expires_in"`   // 过期时间
	AccessToken string `json:"access_token"` // 获取到的凭证
}

// 获取access_token
// 应该全局缓存该token，不能频繁刷新
// 目前会缓存在Client里
// https://open-doc.dingtalk.com/docs/doc.htm?spm=a219a.7629140.0.0.jATUnK&treeId=385&articleId=104980&docType=1
// TODO 过期刷新
func (c *Client) GetToken() (t Token, err error) {
	err = c.getJSON(fmt.Sprintf("%s/gettoken?corpid=%s&corpsecret=%s", DD_API_HOST, c.corpID, c.corpSecret),
		nil, &t)
	c.cacheToken = t
	return
}
