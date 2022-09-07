package agora

import (
	"encoding/base64"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

const (
	Version       = "007"
	VersionLength = 3
)

type AccessToken2 struct {
	appCert  string
	appID    string
	expire   int64
	issueTs  int64
	salt     int
	services map[uint16]Servicer
}

func NewAccessToken2(appID, appCert string, expire int64) *AccessToken2 {
	return &AccessToken2{
		appID:   appID,
		appCert: appCert,
		expire:  expire,
		issueTs: time.Now().Unix(),
		salt:    rand.Int(),
	}
}

func (s *AccessToken2) AddService(service Servicer) {
	s.services[service.GetServiceType()] = service
}

func (s *AccessToken2) Build() string {
	if !s.isUUID(s.appID) || s.isUUID(s.appCert) {
		return ""
	}
	signKey := s.getSign()
	data := PackString(s.appID)
	data = append(data, PackUint32(uint32(s.issueTs))...)
	data = append(data, PackUint32(uint32(s.expire))...)
	data = append(data, PackUint32(uint32(s.salt))...)
	data = append(data, PackUint16(uint16(len(s.services)))...)

	for _, v := range s.services {
		data = append(data, v.Pack()...)
	}

	sign := HmacSha256(signKey, data)
	signData := PackString(string(sign))
	signData = append(signData, data...)
	return Version + base64.StdEncoding.EncodeToString(ZlibEncode(signData))
}

func (s *AccessToken2) getSign() []byte {
	hh := HmacSha256(PackUint32(uint32(s.issueTs)), []byte(s.appCert))
	return HmacSha256(PackUint32(uint32(s.salt)), hh)
}

func (s *AccessToken2) isUUID(str string) bool {
	if len(str) != 32 {
		return false
	}
	ok, _ := regexp.MatchString("[0-9a-fA-F]{32}", str)
	return ok
}

// Parse TODO 解析
func (s *AccessToken2) Parse(token string) bool {
	if !strings.HasPrefix(token, Version) {
		return false
	}
	str := strings.TrimPrefix(Version, token)
	destr, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return false
	}
	data := ZlibDecode(destr)
	var sign string
	var issueTs, expire, salt uint32
	var serviceNum, serviceType uint16
	sign, data = UnPackString(data)
	s.appID, data = UnPackString(data)
	issueTs, data = UnPackUint32(data)
	expire, data = UnPackUint32(data)
	salt, data = UnPackUint32(data)
	serviceNum, data = UnPackUint16(data)

	s.issueTs = int64(issueTs)
	s.expire = int64(expire)
	s.salt = int(salt)

	// 对比签名
	if string(s.getSign()) != sign {
		return false
	}

	for i := uint16(0); i < serviceNum; i++ {
		serviceType, data = UnPackUint16(data)
		var srv Servicer
		switch serviceType {
		case ServiceTypeRtc:
			srv = NewServiceRtc("", "")
		case ServiceTypeRtm:
			srv = NewServiceRtc("", "")
		case ServiceTypeStreaming:
			srv = NewServiceRtc("", "")
		case ServiceTypeChat:
			srv = NewServiceRtc("", "")
		}
		if srv != nil {
			data = srv.UnPack(data)
			s.AddService(srv)
		}
	}
	return true
}
