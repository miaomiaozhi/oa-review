package v1

// unauth
type UserLoginRequest struct {
	UserId       string `json:"UserId,omitempty"`       // user id
	UserPassword string `json:"UserPassword,omitempty"` // user password
}
type UserRegisterRequest struct {
	UserId       string `json:"UserId,omitempty"`
	UserPassword string `json:"UserPassword,omitempty"`
	UserName     string `json:"UserName,omitempty"`
	Priority     int32  `json:"Priority,omitempty"`
}

// auth
type UserGetInfoRequest struct {
	UserId int64 `json:"UserId,omitempty"`
}
type UserSubmitApplicationRequest struct {
	UserId             int64  `json:"UserId,omitempty"`
	ApplicationContext string `json:"ApplicationContext,omitempty"`
}
