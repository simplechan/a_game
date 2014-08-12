// Code generated by protoc-gen-go.
// source: login.proto
// DO NOT EDIT!

package protodata

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// ************************************登陆************************************
type LoginRequest struct {
	PlatId           *int32  `protobuf:"varint,1,opt,name=platId" json:"platId,omitempty"`
	Username         *string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password         *string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	OtherId          *string `protobuf:"bytes,4,opt,name=otherId" json:"otherId,omitempty"`
	OtherSession     *string `protobuf:"bytes,5,opt,name=otherSession" json:"otherSession,omitempty"`
	OtherSign        *string `protobuf:"bytes,6,opt,name=otherSign" json:"otherSign,omitempty"`
	OtherData        *string `protobuf:"bytes,7,opt,name=otherData" json:"otherData,omitempty"`
	Ip               *string `protobuf:"bytes,8,opt,name=ip" json:"ip,omitempty"`
	Imei             *string `protobuf:"bytes,9,opt,name=imei" json:"imei,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}

func (m *LoginRequest) GetPlatId() int32 {
	if m != nil && m.PlatId != nil {
		return *m.PlatId
	}
	return 0
}

func (m *LoginRequest) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *LoginRequest) GetOtherId() string {
	if m != nil && m.OtherId != nil {
		return *m.OtherId
	}
	return ""
}

func (m *LoginRequest) GetOtherSession() string {
	if m != nil && m.OtherSession != nil {
		return *m.OtherSession
	}
	return ""
}

func (m *LoginRequest) GetOtherSign() string {
	if m != nil && m.OtherSign != nil {
		return *m.OtherSign
	}
	return ""
}

func (m *LoginRequest) GetOtherData() string {
	if m != nil && m.OtherData != nil {
		return *m.OtherData
	}
	return ""
}

func (m *LoginRequest) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *LoginRequest) GetImei() string {
	if m != nil && m.Imei != nil {
		return *m.Imei
	}
	return ""
}

type LoginResponse struct {
	TokenStr         *string `protobuf:"bytes,1,opt,name=tokenStr" json:"tokenStr,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}

func (m *LoginResponse) GetTokenStr() string {
	if m != nil && m.TokenStr != nil {
		return *m.TokenStr
	}
	return ""
}

func init() {
}
