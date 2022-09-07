package agora

type ServiceRtc struct {
	Service
	ChannelName string
	UID         string
}

func NewServiceRtc(channelName, uid string) *ServiceRtc {
	return &ServiceRtc{
		Service:     Service{Type: ServiceTypeRtc},
		ChannelName: channelName,
		UID:         uid,
	}
}

func (s *ServiceRtc) Pack() []byte {
	return append(append(s.Service.Pack(), PackString(s.ChannelName)...), PackString(s.UID)...)
}

func (s *ServiceRtc) UnPack(data []byte) []byte {
	data = s.Service.UnPack(data)
	s.ChannelName, data = UnPackString(data)
	s.UID, data = UnPackString(data)
	return data
}
