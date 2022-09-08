package agora

type ServiceFpa struct {
	Service
}

func NewServiceFpa(userID string) *ServiceFpa {
	return &ServiceFpa{
		Service: Service{Type: ServiceTypeFpa, Privileges: map[uint16]uint32{}},
	}
}

func (s *ServiceFpa) Pack() []byte {
	return s.Service.Pack()
}

func (s *ServiceFpa) UnPack(data []byte) []byte {
	return s.Service.UnPack(data)
}
