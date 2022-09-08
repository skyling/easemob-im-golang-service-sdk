package agora

type ServiceStreaming struct {
	Service
	ChannelName string
	UID         string
}

func NewServiceStreaming(channelName, uid string) *ServiceStreaming {
	return &ServiceStreaming{
		Service:     Service{Type: ServiceTypeRtc, Privileges: map[uint16]uint32{}},
		ChannelName: channelName,
		UID:         uid,
	}
}

func (s *ServiceStreaming) Pack() []byte {
	return append(append(s.Service.Pack(), PackString(s.ChannelName)...), PackString(s.UID)...)
}

func (s *ServiceStreaming) UnPack(data []byte) []byte {
	data = s.Service.UnPack(data)
	s.ChannelName, data = UnPackString(data)
	s.UID, data = UnPackString(data)
	return data
}
