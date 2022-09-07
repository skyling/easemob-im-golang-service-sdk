package agora

type ServiceChat struct {
	Service
	UserID string
}

func NewServiceChat(userID string) *ServiceChat {
	return &ServiceChat{
		Service: Service{Type: ServiceTypeChat},
		UserID:  userID,
	}
}

func (s *ServiceChat) Pack() []byte {
	return append(s.Service.Pack(), PackString(s.UserID)...)
}

func (s *ServiceChat) UnPack(data []byte) []byte {
	data = s.Service.UnPack(data)
	s.UserID, data = UnPackString(data)
	return data
}
