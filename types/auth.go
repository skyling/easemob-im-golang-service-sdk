package types

import (
	jsoniter "github.com/json-iterator/go"
)

type ServerResp struct {
	FileVersion string `json:"file_version"`
	ClusterName string `json:"cluster_name"`
	Resolver    struct {
		Hosts []struct {
			Protocol string `json:"protocol,omitempty"`
			Port     string `json:"port"`
			Domain   string `json:"domain"`
			Channel  string `json:"channel"`
			Ip       string `json:"ip,omitempty"`
		} `json:"hosts"`
	} `json:"resolver"`
	Rest struct {
		Hosts []struct {
			Protocol string `json:"protocol"`
			Port     string `json:"port"`
			Domain   string `json:"domain"`
			Channel  string `json:"channel"`
			Priority string `json:"priority"`
			Ip       string `json:"ip,omitempty"`
		} `json:"hosts"`
	} `json:"rest"`
	MsyncWs struct {
		Hosts []struct {
			Protocol string `json:"protocol"`
			Port     string `json:"port"`
			Domain   string `json:"domain,omitempty"`
			Channel  string `json:"channel"`
			Priority string `json:"priority"`
			Ip       string `json:"ip,omitempty"`
		} `json:"hosts"`
	} `json:"msync-ws"`
	Im struct {
		Hosts []struct {
			Port     string `json:"port"`
			Domain   string `json:"domain"`
			Ip       string `json:"ip"`
			Channel  string `json:"channel"`
			Priority string `json:"priority"`
		} `json:"hosts"`
	} `json:"im"`
	MsyncWx struct {
		Hosts []struct {
			Protocol string `json:"protocol"`
			Port     string `json:"port"`
			Domain   string `json:"domain"`
			Channel  string `json:"channel"`
			Priority string `json:"priority"`
		} `json:"hosts"`
	} `json:"msync-wx"`
	OpenReportId string `json:"openReportId"`
	MsyncIm      struct {
		Hosts []struct {
			Port     string `json:"port"`
			Domain   string `json:"domain"`
			Ip       string `json:"ip"`
			Channel  string `json:"channel"`
			Priority string `json:"priority"`
		} `json:"hosts"`
	} `json:"msync-im"`
	CloseReportId string `json:"closeReportId"`
	Xmpp          struct {
		Hosts []struct {
			Protocol string `json:"protocol"`
			Port     string `json:"port"`
			Domain   string `json:"domain"`
			Ip       string `json:"ip,omitempty"`
			Channel  string `json:"channel"`
			Priority string `json:"priority"`
		} `json:"hosts"`
	} `json:"xmpp"`
	ValidBefore      string `json:"valid_before"`
	DeployName       string `json:"deploy_name"`
	EnableDataReport string `json:"enableDataReport"`
	TcpResolver      struct {
		Hosts []struct {
			Port    string `json:"port"`
			Domain  string `json:"domain"`
			Ip      string `json:"ip"`
			Channel string `json:"channel"`
		} `json:"hosts"`
	} `json:"tcp_resolver"`
	GcmEnabled string `json:"gcm_enabled"`
}

type ErrResp struct {
	Err              string `json:"error"`
	Exception        string `json:"exception"`
	Timestamp        int64  `json:"timestamp"`
	Duration         int    `json:"duration"`
	ErrorDescription string `json:"error_description"`
}

func (s *ErrResp) Error() string {
	str, _ := jsoniter.ConfigDefault.MarshalToString(s)
	return str
}

type AccessTokenResp struct {
	Application string `json:"application"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}
