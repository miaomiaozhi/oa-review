package v1

type UserRegisterResponse struct {
	StatusCode int32  `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=StatusMsg,proto3" json:"StatusMsg,omitempty"`
}

type UserLoginResponse struct {
	StatusCode int32  `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=StatusMsg,proto3" json:"StatusMsg,omitempty"`
	Token      string `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
}
