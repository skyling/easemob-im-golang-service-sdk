package agora

type ServiceRtm struct {
	Service
	UserID string
}

func NewServiceRtm(userID string) *ServiceRtm {
	return &ServiceRtm{
		Service: Service{Type: ServiceTypeRtm},
		UserID:  userID,
	}
}

func (s *ServiceRtm) Pack() []byte {
	return s.Service.Pack()
}

func (s *ServiceRtm) UnPack(data []byte) []byte {
	data = s.Service.UnPack(data)
	s.UserID, data = UnPackString(data)
	return data
}
