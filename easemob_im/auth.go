package easemob_im

import (
	"errors"
	"fmt"
	"github.com/skyling/easemob-im-golang-service-sdk/agora"
	"github.com/skyling/easemob-im-golang-service-sdk/cache"
	"github.com/skyling/easemob-im-golang-service-sdk/types"
	"strings"
	"time"
)

type Auth struct {
	apiURI              string // rest 域名
	orgName             string // 企业唯一标识 org_name
	appName             string // 企业下 APP 唯一标识 app_name
	appKey              string // APP 的唯一标识，规则是 ${org_name}#${app_name}
	clientID            string // App 的 client_id
	clientSecret        string // App 的 client_secret
	easeMobToken        string // 环信 token
	isAgora             bool   // 是否以声网 token 初始化标识
	appID               string // 声网 App ID
	appCertificate      string // 声网 App 证书
	expireTimeInSeconds int64  //  声网 token 有效期，单位（秒）
	uuid                string // 环信用户 uuid
	//agoraToken2EaseMobToken string // 声网 token 置换的环信 token
}

func NewAuth(appKey, clientIDOrAppID, clientSecretOrAppCertificate string, isAgora bool) (*Auth, error) {
	if !strings.Contains(appKey, "#") {
		return nil, errors.New("appKey 数据结构有误")
	}
	temp := strings.Split(appKey, "#")
	auth := &Auth{
		orgName:             temp[0],
		appName:             temp[1],
		appKey:              appKey,
		expireTimeInSeconds: 259200,
		isAgora:             isAgora,
	}
	if isAgora {
		auth.appID = clientIDOrAppID
		auth.appCertificate = clientSecretOrAppCertificate
	} else {
		auth.clientID = clientIDOrAppID
		auth.clientSecret = clientSecretOrAppCertificate
	}
	return auth, nil
}

func (s *Auth) SetExpireTimeInSeconds(expireTimeInSeconds int64) *Auth {
	s.expireTimeInSeconds = expireTimeInSeconds
	return s
}

func (s *Auth) SetUUID(uuid string) *Auth {
	s.uuid = uuid
	return s
}

func (s *Auth) GetBaseURI() string {
	return fmt.Sprintf("%s/%s/%s", s.GetApiURI(), s.orgName, s.appName)
}

func (s *Auth) BuildURI(path string) string {
	return fmt.Sprintf("%s/%s", s.GetBaseURI(), strings.TrimPrefix(path, "/"))
}

func (s *Auth) GetApiURI() string {
	if s.apiURI != "" {
		return s.apiURI
	}
	return s.getRemoteApiURI()
}

func (s *Auth) SetApiURI(apiURI string) *Auth {
	s.apiURI = apiURI
	return s
}

func (s *Auth) getEaseMobTokenCacheKey() string {
	return s.appKey + "_easemob_token"
}

func (s *Auth) getAgoraTokenCacheKey() string {
	return s.appKey + "_easemob_token"
}

func (s *Auth) getAgoraToken2EaseMobTokenCacheKey() string {
	return s.appKey + "_agora_2_easemob_token"
}

func (s *Auth) GetEaseMobToken() string {
	if token, ok := cache.Cache.Get(s.getEaseMobTokenCacheKey()); ok {
		s.easeMobToken = token.(string)
		return s.easeMobToken
	}
	return s.getEaseMobAccessToken()
}

// GetAgoraToken 获取声网 token
func (s *Auth) GetAgoraToken() string {
	if s.isAgora {
		if agoraToken, ok := cache.Cache.Get(s.getAgoraTokenCacheKey()); ok {
			return agoraToken.(string)
		}
		var agoraToken string
		if s.uuid != "" {
			agoraToken = agora.ChatTokenBuilderI.BuildUserToken(s.appID, s.appCertificate, s.uuid, s.expireTimeInSeconds)
		} else {
			agoraToken = agora.ChatTokenBuilderI.BuildAppToken(s.appID, s.appCertificate, s.expireTimeInSeconds)
		}
		if agoraToken != "" {
			cache.Cache.Set(s.getAgoraTokenCacheKey(), agoraToken, time.Duration(s.expireTimeInSeconds)*time.Second)
		}
		return agoraToken
	}
	return ""
}

func (s *Auth) GetAgoraToken2easemobToken() string {
	if s.isAgora {
		if agoraToken2EaseMobToken, ok := cache.Cache.Get(s.getAgoraToken2EaseMobTokenCacheKey()); ok {
			return agoraToken2EaseMobToken.(string)
		}
		return s.agoraToken2EaseMobToken()
	}
	return ""
}

type AuthGetUserTokenParams struct {
	Username        string           // 用户名 或uuid
	Password        string           // 密码
	ExpireInSeconds int64            // 过期时间
	Configuration   []agora.Servicer // 权限
}

func (s *Auth) GetUserToken(params *AuthGetUserTokenParams) (*types.AccessTokenResp, error) {
	if params.ExpireInSeconds == 0 {
		params.ExpireInSeconds = 3600
	}
	if params.Password != "" {
		body := map[string]interface{}{
			"grant_type": "password",
			"username":   params.Username,
			"password":   params.Password,
		}
		if params.ExpireInSeconds > 0 {
			body["ttl"] = params.ExpireInSeconds
		}
		uri := s.GetBaseURI() + "/token"
		res := types.AccessTokenResp{}
		err := HttpPost(uri, body, &res)
		return &res, err
	} else if s.isAgora {
		res := types.AccessTokenResp{
			ExpiresIn: s.expireTimeInSeconds,
		}
		if params.Configuration != nil {
			token2 := agora.NewAccessToken2(s.appID, s.appCertificate, params.ExpireInSeconds)
			for _, v := range params.Configuration {
				token2.AddService(v)
			}
			res.AccessToken = token2.Build()
		} else {
			res.AccessToken = agora.ChatTokenBuilderI.BuildUserToken(s.appID, s.appCertificate, params.Username, params.ExpireInSeconds)
		}

		return &res, nil
	}
	return nil, errors.New("get user token error")
}

// Headers 获取请求头
func (s *Auth) Headers() map[string]string {
	token := ""
	if s.isAgora {
		token = s.GetAgoraToken2easemobToken()
	} else {
		token = s.GetEaseMobToken()
	}
	return map[string]string{
		"Authorization": "Bearer " + token,
	}
}

// getRemoteApiURI 获取 REST API 域名
func (s *Auth) getRemoteApiURI() string {
	uri := fmt.Sprintf("%s/easemob/server.json?app_key=%s", DnsURL, s.appKey)
	var result types.ServerResp
	err := HttpGet(uri, &result)
	if err != nil {
		return ""
	}
	for _, v := range result.Rest.Hosts {
		if v.Protocol == "https" {
			s.apiURI = fmt.Sprintf("%s://%s", v.Protocol, v.Domain)
			return s.apiURI
		}
	}
	return ""
}

// getEaseMobAccessToken 获取环信 token
func (s *Auth) getEaseMobAccessToken() string {
	uri := s.GetBaseURI() + "/token"
	body := map[string]interface{}{
		"grant_type":    "client_credentials",
		"client_id":     s.clientID,
		"client_secret": s.clientSecret,
		"ttl":           s.expireTimeInSeconds,
	}
	token := types.AccessTokenResp{}
	err := HttpPost(uri, body, &token)
	if err != nil {
		return ""
	}
	// set cache
	if token.AccessToken != "" {
		cache.Cache.Set(s.getEaseMobTokenCacheKey(), token.AccessToken, time.Duration(s.expireTimeInSeconds/2)*time.Second)
		return token.AccessToken
	}
	return ""
}

// agoraToken2EaseMobToken 声网 token 置换环信 token
func (s *Auth) agoraToken2EaseMobToken() string {
	agoraToken := s.GetAgoraToken()
	uri := fmt.Sprintf("http://a41.easemob.com/%s/%s/token", s.orgName, s.appName)
	body := map[string]interface{}{
		"grant_type": "agora",
	}
	headers := map[string]string{
		"Authorization": "Bearer " + agoraToken,
	}
	res := types.AccessTokenResp{}
	err := HttpPost(uri, body, &res, headers)
	fmt.Println(err, agoraToken)
	if err != nil {
		return ""
	}
	if res.AccessToken != "" {
		cache.Cache.Set(s.getAgoraToken2EaseMobTokenCacheKey(), res.AccessToken, 2592000)
		return res.AccessToken
	}
	return ""
}
