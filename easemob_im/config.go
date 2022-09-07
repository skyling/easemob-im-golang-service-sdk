package easemob_im

type EaseMobConfig struct {
	AppKey       string `yaml:"app_key" json:"app_key"`
	ClientID     string `yaml:"client_id" json:"client_id"`
	ClientSecret string `yaml:"client_secret" json:"client_secret"`
	ApiURI       string `yaml:"api_uri" json:"api_uri"`
}

type AgoraConfig struct {
	AppKey         string `yaml:"app_key" json:"app_key"`
	AppID          string `yaml:"app_id" json:"app_id"`
	AppCertificate string `yaml:"app_certificate" json:"app_certificate"`
	ExpireTime     uint   `yaml:"expire_time" json:"expire_time"`
}

var (
	DefaultEaseMobConfig = &EaseMobConfig{
		AppKey:       "",
		ClientID:     "",
		ClientSecret: "",
	}
	DefaultAgoraConfig = &EaseMobConfig{}
)
