package v1

// unauth
type UserLoginRequest struct {
	UserId       string `json:"UserId" validate:"numeric,gte=1,required"`      // user id
	UserPassword string `json:"UserPassword" validate:"required,gte=3,lte=10"` // user password
}
type UserRegisterRequest struct {
	UserId       string `json:"UserId" validate:"numeric,gte=1,required"`
	UserPassword string `json:"UserPassword" validate:"required,gte=3,lte=10"`
	UserName     string `json:"UserName" validate:"required,len=2,lt=10"`
	Priority     int32  `json:"Priority" validate:"required,gte=0"`
}

// auth
type UserGetInfoRequest struct {
	UserId int64 `json:"UserId" validate:"required,gte=1"`
}
type UserSubmitApplicationRequest struct {
	UserId             int64  `json:"UserId" validate:"required,gte=1"`
	ApplicationContext string `json:"ApplicationContext" validate:"required"`
}
