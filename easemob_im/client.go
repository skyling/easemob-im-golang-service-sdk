package easemob_im

import (
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/skyling/easemob-im-golang-service-sdk/types"
)

func ErrorHook(req *resty.Request, err error) {
	if v, ok := err.(*resty.ResponseError); ok {
		// todo log
		logrus.Error(v)
		// v.Response contains the last response from the server
		// v.Err contains the original error
	}
}

func RestyClient() *resty.Client {
	return resty.New()
}

func HttpGet(url string, res interface{}, headers ...map[string]string) *types.ErrResp {
	var err *types.ErrResp
	request := RestyClient().R().SetError(&err).SetResult(res)
	if len(headers) > 0 {
		request.SetHeaders(headers[0])
	}
	_, e := request.Get(url)
	if e != nil && err == nil {
		err.Err = e.Error()
		err.ErrorDescription = e.Error()
	}
	return err
}

func HttpPost(url string, req interface{}, res interface{}, headers ...map[string]string) *types.ErrResp {
	var err *types.ErrResp
	request := RestyClient().R().SetError(&err).SetBody(req).SetResult(res)
	if len(headers) > 0 {
		request.SetHeaders(headers[0])
	}
	_, e := request.Post(url)
	if e != nil && err == nil {
		err.Err = e.Error()
		err.ErrorDescription = e.Error()
	}
	return err
}

func HttpPut(url string, req interface{}, res interface{}, headers ...map[string]string) *types.ErrResp {
	var err *types.ErrResp
	request := RestyClient().R().SetError(&err).SetBody(req).SetResult(res)
	if len(headers) > 0 {
		request.SetHeaders(headers[0])
	}
	_, e := request.Put(url)
	if e != nil && err == nil {
		err.Err = e.Error()
		err.ErrorDescription = e.Error()
	}
	return err
}

func HttpDelete(url string, req interface{}, res interface{}, headers ...map[string]string) *types.ErrResp {
	var err *types.ErrResp
	request := RestyClient().R().SetError(&err).SetResult(res)
	if len(headers) > 0 {
		request.SetHeaders(headers[0])
	}

	if req != nil {
		request.SetBody(req)
	}
	_, e := request.Delete(url)
	if e != nil && err == nil {
		err.Err = e.Error()
		err.ErrorDescription = e.Error()
	}
	return err
}
