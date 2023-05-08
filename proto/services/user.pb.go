// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 申请的结构
type Application struct {
	ApplicationId        int64    `protobuf:"varint,1,opt,name=ApplicationId,proto3" json:"ApplicationId,omitempty"`
	Context              string   `protobuf:"bytes,2,opt,name=Context,proto3" json:"Context,omitempty"`
	ReviewStatus         bool     `protobuf:"varint,3,opt,name=ReviewStatus,proto3" json:"ReviewStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *Application) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Application.Unmarshal(m, b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Application.Marshal(b, m, deterministic)
}
func (m *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(m, src)
}
func (m *Application) XXX_Size() int {
	return xxx_messageInfo_Application.Size(m)
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetApplicationId() int64 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *Application) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func (m *Application) GetReviewStatus() bool {
	if m != nil {
		return m.ReviewStatus
	}
	return false
}

// 用户基本信息
type User struct {
	UserId               int64          `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Applications         []*Application `protobuf:"bytes,3,rep,name=Applications,proto3" json:"Applications,omitempty"`
	Priority             int32          `protobuf:"varint,4,opt,name=Priority,proto3" json:"Priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetApplications() []*Application {
	if m != nil {
		return m.Applications
	}
	return nil
}

func (m *User) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

// 用户登录请求 POST
type UserLoginRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	UserPassword         string   `protobuf:"bytes,2,opt,name=UserPassword,proto3" json:"UserPassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLoginRequest) Reset()         { *m = UserLoginRequest{} }
func (m *UserLoginRequest) String() string { return proto.CompactTextString(m) }
func (*UserLoginRequest) ProtoMessage()    {}
func (*UserLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginRequest.Unmarshal(m, b)
}
func (m *UserLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginRequest.Marshal(b, m, deterministic)
}
func (m *UserLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginRequest.Merge(m, src)
}
func (m *UserLoginRequest) XXX_Size() int {
	return xxx_messageInfo_UserLoginRequest.Size(m)
}
func (m *UserLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginRequest proto.InternalMessageInfo

func (m *UserLoginRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserLoginRequest) GetUserPassword() string {
	if m != nil {
		return m.UserPassword
	}
	return ""
}

type UserLoginResponse struct {
	StatusCode           int32    `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	StatusMsg            string   `protobuf:"bytes,2,opt,name=StatusMsg,proto3" json:"StatusMsg,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLoginResponse) Reset()         { *m = UserLoginResponse{} }
func (m *UserLoginResponse) String() string { return proto.CompactTextString(m) }
func (*UserLoginResponse) ProtoMessage()    {}
func (*UserLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UserLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginResponse.Unmarshal(m, b)
}
func (m *UserLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginResponse.Marshal(b, m, deterministic)
}
func (m *UserLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginResponse.Merge(m, src)
}
func (m *UserLoginResponse) XXX_Size() int {
	return xxx_messageInfo_UserLoginResponse.Size(m)
}
func (m *UserLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginResponse proto.InternalMessageInfo

func (m *UserLoginResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *UserLoginResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

func (m *UserLoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// 用户注册 POST
type UserRegisterRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	UserPassword         string   `protobuf:"bytes,2,opt,name=UserPassword,proto3" json:"UserPassword,omitempty"`
	UserName             string   `protobuf:"bytes,3,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Priority             int32    `protobuf:"varint,4,opt,name=Priority,proto3" json:"Priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegisterRequest) Reset()         { *m = UserRegisterRequest{} }
func (m *UserRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*UserRegisterRequest) ProtoMessage()    {}
func (*UserRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UserRegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegisterRequest.Unmarshal(m, b)
}
func (m *UserRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegisterRequest.Marshal(b, m, deterministic)
}
func (m *UserRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegisterRequest.Merge(m, src)
}
func (m *UserRegisterRequest) XXX_Size() int {
	return xxx_messageInfo_UserRegisterRequest.Size(m)
}
func (m *UserRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegisterRequest proto.InternalMessageInfo

func (m *UserRegisterRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserRegisterRequest) GetUserPassword() string {
	if m != nil {
		return m.UserPassword
	}
	return ""
}

func (m *UserRegisterRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserRegisterRequest) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

type UserRegisterResponse struct {
	StatusCode           int32    `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	StatusMsg            string   `protobuf:"bytes,2,opt,name=StatusMsg,proto3" json:"StatusMsg,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegisterResponse) Reset()         { *m = UserRegisterResponse{} }
func (m *UserRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*UserRegisterResponse) ProtoMessage()    {}
func (*UserRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *UserRegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegisterResponse.Unmarshal(m, b)
}
func (m *UserRegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegisterResponse.Marshal(b, m, deterministic)
}
func (m *UserRegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegisterResponse.Merge(m, src)
}
func (m *UserRegisterResponse) XXX_Size() int {
	return xxx_messageInfo_UserRegisterResponse.Size(m)
}
func (m *UserRegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegisterResponse proto.InternalMessageInfo

func (m *UserRegisterResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *UserRegisterResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

func (m *UserRegisterResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// 用户信息 GET
type UserGetInfoRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	UserPassword         string   `protobuf:"bytes,2,opt,name=UserPassword,proto3" json:"UserPassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGetInfoRequest) Reset()         { *m = UserGetInfoRequest{} }
func (m *UserGetInfoRequest) String() string { return proto.CompactTextString(m) }
func (*UserGetInfoRequest) ProtoMessage()    {}
func (*UserGetInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *UserGetInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGetInfoRequest.Unmarshal(m, b)
}
func (m *UserGetInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGetInfoRequest.Marshal(b, m, deterministic)
}
func (m *UserGetInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGetInfoRequest.Merge(m, src)
}
func (m *UserGetInfoRequest) XXX_Size() int {
	return xxx_messageInfo_UserGetInfoRequest.Size(m)
}
func (m *UserGetInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGetInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserGetInfoRequest proto.InternalMessageInfo

func (m *UserGetInfoRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserGetInfoRequest) GetUserPassword() string {
	if m != nil {
		return m.UserPassword
	}
	return ""
}

type UserGetInfoResponse struct {
	StatusCode           int32    `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	StatusMsg            string   `protobuf:"bytes,2,opt,name=StatusMsg,proto3" json:"StatusMsg,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGetInfoResponse) Reset()         { *m = UserGetInfoResponse{} }
func (m *UserGetInfoResponse) String() string { return proto.CompactTextString(m) }
func (*UserGetInfoResponse) ProtoMessage()    {}
func (*UserGetInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *UserGetInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGetInfoResponse.Unmarshal(m, b)
}
func (m *UserGetInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGetInfoResponse.Marshal(b, m, deterministic)
}
func (m *UserGetInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGetInfoResponse.Merge(m, src)
}
func (m *UserGetInfoResponse) XXX_Size() int {
	return xxx_messageInfo_UserGetInfoResponse.Size(m)
}
func (m *UserGetInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGetInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserGetInfoResponse proto.InternalMessageInfo

func (m *UserGetInfoResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *UserGetInfoResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

func (m *UserGetInfoResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// 用户提交申请请求 POST
type UserSubmitApplicationRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	ApplicationContext   string   `protobuf:"bytes,2,opt,name=ApplicationContext,proto3" json:"ApplicationContext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSubmitApplicationRequest) Reset()         { *m = UserSubmitApplicationRequest{} }
func (m *UserSubmitApplicationRequest) String() string { return proto.CompactTextString(m) }
func (*UserSubmitApplicationRequest) ProtoMessage()    {}
func (*UserSubmitApplicationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *UserSubmitApplicationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSubmitApplicationRequest.Unmarshal(m, b)
}
func (m *UserSubmitApplicationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSubmitApplicationRequest.Marshal(b, m, deterministic)
}
func (m *UserSubmitApplicationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSubmitApplicationRequest.Merge(m, src)
}
func (m *UserSubmitApplicationRequest) XXX_Size() int {
	return xxx_messageInfo_UserSubmitApplicationRequest.Size(m)
}
func (m *UserSubmitApplicationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSubmitApplicationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserSubmitApplicationRequest proto.InternalMessageInfo

func (m *UserSubmitApplicationRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserSubmitApplicationRequest) GetApplicationContext() string {
	if m != nil {
		return m.ApplicationContext
	}
	return ""
}

type UserSubmitApplicationResponse struct {
	StatusCode           int32    `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	StatusMsg            string   `protobuf:"bytes,2,opt,name=StatusMsg,proto3" json:"StatusMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSubmitApplicationResponse) Reset()         { *m = UserSubmitApplicationResponse{} }
func (m *UserSubmitApplicationResponse) String() string { return proto.CompactTextString(m) }
func (*UserSubmitApplicationResponse) ProtoMessage()    {}
func (*UserSubmitApplicationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}

func (m *UserSubmitApplicationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSubmitApplicationResponse.Unmarshal(m, b)
}
func (m *UserSubmitApplicationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSubmitApplicationResponse.Marshal(b, m, deterministic)
}
func (m *UserSubmitApplicationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSubmitApplicationResponse.Merge(m, src)
}
func (m *UserSubmitApplicationResponse) XXX_Size() int {
	return xxx_messageInfo_UserSubmitApplicationResponse.Size(m)
}
func (m *UserSubmitApplicationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSubmitApplicationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserSubmitApplicationResponse proto.InternalMessageInfo

func (m *UserSubmitApplicationResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *UserSubmitApplicationResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

// 用户检索申请请求 GET
type UserRetrievalApplicationRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRetrievalApplicationRequest) Reset()         { *m = UserRetrievalApplicationRequest{} }
func (m *UserRetrievalApplicationRequest) String() string { return proto.CompactTextString(m) }
func (*UserRetrievalApplicationRequest) ProtoMessage()    {}
func (*UserRetrievalApplicationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{10}
}

func (m *UserRetrievalApplicationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRetrievalApplicationRequest.Unmarshal(m, b)
}
func (m *UserRetrievalApplicationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRetrievalApplicationRequest.Marshal(b, m, deterministic)
}
func (m *UserRetrievalApplicationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRetrievalApplicationRequest.Merge(m, src)
}
func (m *UserRetrievalApplicationRequest) XXX_Size() int {
	return xxx_messageInfo_UserRetrievalApplicationRequest.Size(m)
}
func (m *UserRetrievalApplicationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRetrievalApplicationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRetrievalApplicationRequest proto.InternalMessageInfo

func (m *UserRetrievalApplicationRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type UserRetrievalApplicationResponse struct {
	StatusCode           int32          `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	StatusMsg            string         `protobuf:"bytes,2,opt,name=StatusMsg,proto3" json:"StatusMsg,omitempty"`
	Applications         []*Application `protobuf:"bytes,3,rep,name=Applications,proto3" json:"Applications,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UserRetrievalApplicationResponse) Reset()         { *m = UserRetrievalApplicationResponse{} }
func (m *UserRetrievalApplicationResponse) String() string { return proto.CompactTextString(m) }
func (*UserRetrievalApplicationResponse) ProtoMessage()    {}
func (*UserRetrievalApplicationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{11}
}

func (m *UserRetrievalApplicationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRetrievalApplicationResponse.Unmarshal(m, b)
}
func (m *UserRetrievalApplicationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRetrievalApplicationResponse.Marshal(b, m, deterministic)
}
func (m *UserRetrievalApplicationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRetrievalApplicationResponse.Merge(m, src)
}
func (m *UserRetrievalApplicationResponse) XXX_Size() int {
	return xxx_messageInfo_UserRetrievalApplicationResponse.Size(m)
}
func (m *UserRetrievalApplicationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRetrievalApplicationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserRetrievalApplicationResponse proto.InternalMessageInfo

func (m *UserRetrievalApplicationResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *UserRetrievalApplicationResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

func (m *UserRetrievalApplicationResponse) GetApplications() []*Application {
	if m != nil {
		return m.Applications
	}
	return nil
}

// 提交审核 POST
type UserSubmitReviewRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	ApplicationId        int64    `protobuf:"varint,2,opt,name=ApplicationId,proto3" json:"ApplicationId,omitempty"`
	ReviewStatus         bool     `protobuf:"varint,3,opt,name=ReviewStatus,proto3" json:"ReviewStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSubmitReviewRequest) Reset()         { *m = UserSubmitReviewRequest{} }
func (m *UserSubmitReviewRequest) String() string { return proto.CompactTextString(m) }
func (*UserSubmitReviewRequest) ProtoMessage()    {}
func (*UserSubmitReviewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{12}
}

func (m *UserSubmitReviewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSubmitReviewRequest.Unmarshal(m, b)
}
func (m *UserSubmitReviewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSubmitReviewRequest.Marshal(b, m, deterministic)
}
func (m *UserSubmitReviewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSubmitReviewRequest.Merge(m, src)
}
func (m *UserSubmitReviewRequest) XXX_Size() int {
	return xxx_messageInfo_UserSubmitReviewRequest.Size(m)
}
func (m *UserSubmitReviewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSubmitReviewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserSubmitReviewRequest proto.InternalMessageInfo

func (m *UserSubmitReviewRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserSubmitReviewRequest) GetApplicationId() int64 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *UserSubmitReviewRequest) GetReviewStatus() bool {
	if m != nil {
		return m.ReviewStatus
	}
	return false
}

type UserSubmitReviewResponse struct {
	StatusCode           int32    `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	StatusMsg            string   `protobuf:"bytes,2,opt,name=StatusMsg,proto3" json:"StatusMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSubmitReviewResponse) Reset()         { *m = UserSubmitReviewResponse{} }
func (m *UserSubmitReviewResponse) String() string { return proto.CompactTextString(m) }
func (*UserSubmitReviewResponse) ProtoMessage()    {}
func (*UserSubmitReviewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{13}
}

func (m *UserSubmitReviewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSubmitReviewResponse.Unmarshal(m, b)
}
func (m *UserSubmitReviewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSubmitReviewResponse.Marshal(b, m, deterministic)
}
func (m *UserSubmitReviewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSubmitReviewResponse.Merge(m, src)
}
func (m *UserSubmitReviewResponse) XXX_Size() int {
	return xxx_messageInfo_UserSubmitReviewResponse.Size(m)
}
func (m *UserSubmitReviewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSubmitReviewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserSubmitReviewResponse proto.InternalMessageInfo

func (m *UserSubmitReviewResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *UserSubmitReviewResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

// 撤回审核
type UserWithdrawReviewRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserWithdrawReviewRequest) Reset()         { *m = UserWithdrawReviewRequest{} }
func (m *UserWithdrawReviewRequest) String() string { return proto.CompactTextString(m) }
func (*UserWithdrawReviewRequest) ProtoMessage()    {}
func (*UserWithdrawReviewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{14}
}

func (m *UserWithdrawReviewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserWithdrawReviewRequest.Unmarshal(m, b)
}
func (m *UserWithdrawReviewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserWithdrawReviewRequest.Marshal(b, m, deterministic)
}
func (m *UserWithdrawReviewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserWithdrawReviewRequest.Merge(m, src)
}
func (m *UserWithdrawReviewRequest) XXX_Size() int {
	return xxx_messageInfo_UserWithdrawReviewRequest.Size(m)
}
func (m *UserWithdrawReviewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserWithdrawReviewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserWithdrawReviewRequest proto.InternalMessageInfo

func (m *UserWithdrawReviewRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type UserWithdrawReviewResponse struct {
	StatusCode           int32    `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	StatusMsg            string   `protobuf:"bytes,2,opt,name=StatusMsg,proto3" json:"StatusMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserWithdrawReviewResponse) Reset()         { *m = UserWithdrawReviewResponse{} }
func (m *UserWithdrawReviewResponse) String() string { return proto.CompactTextString(m) }
func (*UserWithdrawReviewResponse) ProtoMessage()    {}
func (*UserWithdrawReviewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{15}
}

func (m *UserWithdrawReviewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserWithdrawReviewResponse.Unmarshal(m, b)
}
func (m *UserWithdrawReviewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserWithdrawReviewResponse.Marshal(b, m, deterministic)
}
func (m *UserWithdrawReviewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserWithdrawReviewResponse.Merge(m, src)
}
func (m *UserWithdrawReviewResponse) XXX_Size() int {
	return xxx_messageInfo_UserWithdrawReviewResponse.Size(m)
}
func (m *UserWithdrawReviewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserWithdrawReviewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserWithdrawReviewResponse proto.InternalMessageInfo

func (m *UserWithdrawReviewResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *UserWithdrawReviewResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*Application)(nil), "user.Application")
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*UserLoginRequest)(nil), "user.UserLoginRequest")
	proto.RegisterType((*UserLoginResponse)(nil), "user.UserLoginResponse")
	proto.RegisterType((*UserRegisterRequest)(nil), "user.UserRegisterRequest")
	proto.RegisterType((*UserRegisterResponse)(nil), "user.UserRegisterResponse")
	proto.RegisterType((*UserGetInfoRequest)(nil), "user.UserGetInfoRequest")
	proto.RegisterType((*UserGetInfoResponse)(nil), "user.UserGetInfoResponse")
	proto.RegisterType((*UserSubmitApplicationRequest)(nil), "user.UserSubmitApplicationRequest")
	proto.RegisterType((*UserSubmitApplicationResponse)(nil), "user.UserSubmitApplicationResponse")
	proto.RegisterType((*UserRetrievalApplicationRequest)(nil), "user.UserRetrievalApplicationRequest")
	proto.RegisterType((*UserRetrievalApplicationResponse)(nil), "user.UserRetrievalApplicationResponse")
	proto.RegisterType((*UserSubmitReviewRequest)(nil), "user.UserSubmitReviewRequest")
	proto.RegisterType((*UserSubmitReviewResponse)(nil), "user.UserSubmitReviewResponse")
	proto.RegisterType((*UserWithdrawReviewRequest)(nil), "user.UserWithdrawReviewRequest")
	proto.RegisterType((*UserWithdrawReviewResponse)(nil), "user.UserWithdrawReviewResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x9b, 0xa4, 0x6d, 0x26, 0x01, 0xd1, 0x21, 0x6a, 0x1d, 0xab, 0x4d, 0xad, 0xe5, 0x47,
	0x39, 0x05, 0x29, 0x15, 0x12, 0x08, 0x09, 0xa9, 0xf4, 0x80, 0x2a, 0xd1, 0x2a, 0x72, 0x40, 0xa0,
	0x0a, 0x0e, 0x6e, 0xb3, 0x0d, 0x86, 0xd6, 0x4e, 0xbd, 0x9b, 0x04, 0xc4, 0x99, 0x03, 0x4f, 0xc0,
	0x83, 0xf0, 0x82, 0x68, 0x7f, 0x12, 0xaf, 0x1b, 0xbb, 0x09, 0x10, 0x4e, 0xde, 0xd9, 0x9d, 0xfd,
	0xe6, 0x9b, 0xd9, 0x99, 0x4f, 0x06, 0x18, 0x32, 0x1a, 0xb7, 0x06, 0x71, 0xc4, 0x23, 0x2c, 0x8a,
	0x35, 0xb9, 0x82, 0xca, 0xfe, 0x60, 0x70, 0x11, 0x9c, 0xf9, 0x3c, 0x88, 0x42, 0xbc, 0x0f, 0xb7,
	0x0c, 0xf3, 0xb0, 0x67, 0x5b, 0xae, 0xd5, 0x2c, 0x78, 0xe9, 0x4d, 0xb4, 0x61, 0xed, 0x20, 0x0a,
	0x39, 0xfd, 0xc2, 0xed, 0x15, 0xd7, 0x6a, 0x96, 0xbd, 0x89, 0x89, 0x04, 0xaa, 0x1e, 0x1d, 0x05,
	0x74, 0xdc, 0xe5, 0x3e, 0x1f, 0x32, 0xbb, 0xe0, 0x5a, 0xcd, 0x75, 0x2f, 0xb5, 0x47, 0xbe, 0x5b,
	0x50, 0x7c, 0xc3, 0x68, 0x8c, 0x9b, 0xb0, 0x2a, 0xbe, 0xd3, 0x28, 0xda, 0x42, 0x84, 0xe2, 0xb1,
	0x7f, 0x49, 0x35, 0xb6, 0x5c, 0xe3, 0x63, 0xa8, 0x1a, 0x1c, 0x04, 0x70, 0xa1, 0x59, 0x69, 0x6f,
	0xb4, 0x64, 0x42, 0xc6, 0x89, 0x97, 0x72, 0x43, 0x07, 0xd6, 0x3b, 0x71, 0x10, 0xc5, 0x01, 0xff,
	0x6a, 0x17, 0x5d, 0xab, 0x59, 0xf2, 0xa6, 0x36, 0x39, 0x86, 0x3b, 0x22, 0xe0, 0xab, 0xa8, 0x1f,
	0x84, 0x1e, 0xbd, 0x1a, 0x52, 0xc6, 0xaf, 0x51, 0x2a, 0x4f, 0x29, 0x11, 0xa8, 0x8a, 0x55, 0xc7,
	0x67, 0x6c, 0x1c, 0xc5, 0x3d, 0x4d, 0x2d, 0xb5, 0x47, 0xfa, 0xb0, 0x61, 0xe0, 0xb1, 0x41, 0x14,
	0x32, 0x8a, 0x0d, 0x00, 0x95, 0xf6, 0x41, 0xd4, 0xa3, 0x12, 0xb4, 0xe4, 0x19, 0x3b, 0xb8, 0x0d,
	0x65, 0x65, 0x1d, 0xb1, 0xbe, 0x46, 0x4d, 0x36, 0xb0, 0x06, 0xa5, 0xd7, 0xd1, 0x67, 0x1a, 0xca,
	0x3a, 0x96, 0x3d, 0x65, 0x90, 0x1f, 0x16, 0xdc, 0x15, 0x91, 0x3c, 0xda, 0x0f, 0x18, 0x17, 0xdf,
	0x7f, 0x26, 0x2f, 0x0a, 0x25, 0x6c, 0x59, 0x77, 0x15, 0x6c, 0x6a, 0xdf, 0x58, 0xc4, 0x4f, 0x50,
	0x4b, 0x53, 0xf9, 0x8f, 0x79, 0x77, 0x00, 0x45, 0xac, 0x97, 0x94, 0x1f, 0x86, 0xe7, 0x51, 0x76,
	0xd6, 0x85, 0x3f, 0x7a, 0x32, 0xa6, 0x0a, 0x39, 0x45, 0x5c, 0x0a, 0xf9, 0x06, 0xc8, 0xd1, 0x92,
	0xdc, 0x2b, 0x6d, 0x50, 0x2d, 0x2a, 0x8b, 0xa4, 0x46, 0xee, 0x1c, 0xb6, 0x85, 0xd5, 0x1d, 0x9e,
	0x5e, 0x06, 0xdc, 0x6c, 0xdd, 0x39, 0x09, 0xb5, 0x00, 0x0d, 0xef, 0xf4, 0x00, 0x66, 0x9c, 0x90,
	0x0f, 0xb0, 0x93, 0x13, 0x67, 0x19, 0x69, 0x92, 0xa7, 0xb0, 0xab, 0x5e, 0x9e, 0xc7, 0x01, 0x1d,
	0xf9, 0x17, 0x8b, 0x67, 0x42, 0x7e, 0x5a, 0xe0, 0xe6, 0xdf, 0x5d, 0xca, 0x23, 0xfc, 0x9d, 0x5e,
	0x90, 0x6f, 0xb0, 0x95, 0xd4, 0x4c, 0xa9, 0xd6, 0xbc, 0x67, 0x99, 0x91, 0xcc, 0x95, 0x2c, 0xc9,
	0x5c, 0x44, 0x18, 0xdf, 0x81, 0x3d, 0x1b, 0x7c, 0x29, 0x6f, 0xb5, 0x07, 0x75, 0x81, 0xfc, 0x36,
	0xe0, 0x1f, 0x7b, 0xb1, 0x3f, 0x5e, 0x28, 0x31, 0x72, 0x02, 0x4e, 0xd6, 0xa5, 0x65, 0x10, 0x6a,
	0xff, 0x2a, 0x42, 0x45, 0xe6, 0x4a, 0xe3, 0x51, 0x70, 0x46, 0xf1, 0x09, 0x94, 0xa4, 0x6e, 0xe2,
	0x66, 0x32, 0x2e, 0xa6, 0x30, 0x3b, 0x5b, 0x33, 0xfb, 0x9a, 0xc7, 0x3e, 0xac, 0x4f, 0xc4, 0x07,
	0xeb, 0xc6, 0xac, 0xa5, 0xb5, 0xd1, 0x71, 0xb2, 0x8e, 0x34, 0xc4, 0x73, 0x58, 0xd3, 0x0a, 0x80,
	0x76, 0xe2, 0x96, 0x96, 0x19, 0xa7, 0x9e, 0x71, 0xa2, 0xef, 0xbf, 0x87, 0x8d, 0x99, 0x21, 0x43,
	0x92, 0xf8, 0xe7, 0x4d, 0xba, 0x73, 0xef, 0x46, 0x1f, 0x8d, 0x4e, 0xa1, 0x96, 0x35, 0x27, 0xf8,
	0xc0, 0xcc, 0x28, 0x77, 0x06, 0x9d, 0x87, 0xf3, 0xdc, 0x74, 0x98, 0x23, 0xa8, 0x9a, 0x8d, 0x87,
	0x3b, 0xd7, 0xb9, 0xa5, 0x9a, 0xc6, 0x69, 0xe4, 0x1d, 0x6b, 0xb8, 0x2e, 0xdc, 0x4e, 0x37, 0x0e,
	0xee, 0x26, 0x37, 0x32, 0xfb, 0xd0, 0x71, 0xf3, 0x1d, 0x14, 0xe8, 0x8b, 0xcd, 0x93, 0x5a, 0xab,
	0xf5, 0x88, 0xa9, 0x9e, 0x61, 0xcf, 0x26, 0x8b, 0xd3, 0x55, 0xf9, 0x47, 0xb3, 0xf7, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0xc3, 0x0a, 0x9f, 0x92, 0xdf, 0x08, 0x00, 0x00,
}