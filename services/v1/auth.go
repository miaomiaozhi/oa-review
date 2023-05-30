package v1

type AuthResult struct {
	Token    string `json:"Token"`    // token
	UserID   int64  `json:"UserId"`   //用户id
	UserName string `json:"UserName"` //用户名
	Priority int64  `json:"Priority"` // 优先级
}
