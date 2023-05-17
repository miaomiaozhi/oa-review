package protoresp

type BaseResponse struct {
	StatusCode int32       `json:"StatusCode"` // 状态码
	StatusMsg  string      `json:"StatusMsg"`  // 状态信息
	Data       interface{} `json:"Data"`       // 数据
}
