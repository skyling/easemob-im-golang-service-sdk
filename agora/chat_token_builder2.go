package agora

var ChatTokenBuilderI = ChatTokenBuilder2{}

type ChatTokenBuilder2 struct{}

func (s *ChatTokenBuilder2) BuildUserToken(appID, appCertificate, userID string, expire int64) string {
	accessToken := NewAccessToken2(appID, appCertificate, expire)
	serviceChat := NewServiceChat(userID)

	serviceChat.AddPrivilege(PrivilegeChatUser, uint32(expire))
	accessToken.AddService(serviceChat)
	return accessToken.Build()
}

func (s *ChatTokenBuilder2) BuildAppToken(appID, appCertificate string, expire int64) string {
	accessToken := NewAccessToken2(appID, appCertificate, expire)
	serviceChat := NewServiceChat("")

	serviceChat.AddPrivilege(PrivilegeChatApp, uint32(expire))
	accessToken.AddService(serviceChat)
	return accessToken.Build()
}
