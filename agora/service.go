package agora

const (
	ServiceTypeRtc                 = uint16(1)
	PrivilegeRtcJoinChannel        = uint16(1)
	PrivilegeRtcPublishAudioStream = uint16(2)
	PrivilegeRtcPublishVideoStream = uint16(3)
	PrivilegeRtcPublishDataStream  = uint16(4)

	ServiceTypeRtm    = uint16(2)
	PrivilegeRtmLogin = uint16(1)

	ServiceTypeStreaming               = uint16(3)
	PrivilegeStreamingPublishMixStream = uint16(1)
	PrivilegeStreamingPublishRawStream = uint16(2)

	ServiceTypeFpa    = uint16(4)
	PrivilegeFpaLogin = uint16(1)

	ServiceTypeChat   = uint16(5)
	PrivilegeChatUser = uint16(1)
	PrivilegeChatApp  = uint16(2)
)

type Servicer interface {
	AddPrivilege(privilege uint16, expire uint32)
	Pack() []byte
	UnPack(data []byte) (ret []byte)
	GetServiceType() uint16
}
type Service struct {
	Type       uint16
	Privileges map[uint16]uint32
}

func (s *Service) AddPrivilege(privilege uint16, expire uint32) {
	s.Privileges[privilege] = expire
}

func (s *Service) Pack() []byte {
	ret := PackUint16(s.Type)
	ret = append(ret, PackMapUint32(s.Privileges)...)
	return ret
}

func (s *Service) UnPack(data []byte) (ret []byte) {
	s.Privileges, ret = UnpackMapUint32(data)
	return ret
}

func (s *Service) GetServiceType() uint16 {
	return s.Type
}
