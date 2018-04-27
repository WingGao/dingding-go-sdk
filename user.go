package dd

type UserInfo struct {
	CommonReply
	UserId   string `json:"userid"`    // 员工在企业内的UserID
	DeviceId string `json:"deviceId"`  // 手机设备号,由钉钉在安装时随机产生
	IsSys    bool   `json:"is_sys"`    // 是否是管理员
	SysLevel int    `json:"sys_level"` // 级别，0：非管理员 1：超级管理员（主管理员） 2：普通管理员（子管理员） 100：老板
}

// 获取成员详情 `/user/get`
func (c *Client) GetUserInfo(code string) (info UserInfo, err error) {
	err = c.getJSON(DD_API_HOST+"/user/getuserinfo?access_token="+c.cacheToken.AccessToken+"&code="+code, nil, &info)
	return
}

type UserInfoDetail struct {
	CommonReply
	UserId string `json:"userid"` // 员工在企业内的UserID
	OpenId string `json:"openid"` // 在本 服务窗运营服务商 范围内,唯一标识关注者身份的id（不可修改）
	Name   string `json:"name"`   // 成员名称
}

// 获取成员详情  `/user/get`
func (c *Client) GetUserInfoDetail(userid string) (info UserInfoDetail, err error) {
	err = c.getJSON(DD_API_HOST+"/user/get?access_token="+c.cacheToken.AccessToken+"&userid="+userid, nil, &info)
	return
}
