package util

import (
	"net/http"
	netUrl "net/url"
	"strings"

	"github.com/mozillazg/request"
)

func HttpPost3(url string, body string) (int, []byte, error) {
	return HttpPost(url, body, 3)
}
func HttpPost(url string, body string, retryCount int) (int, []byte, error) {
	c := new(http.Client)
	req := request.NewRequest(c)
	req.Body = strings.NewReader(body)
	resp, err := req.Post(url)
	defer resp.Body.Close()
	if err != nil {
		if retryCount > 0 {
			return HttpPost(url, body, retryCount-1)
		}
		return 0, nil, err
	}
	res, err := resp.Content()
	return resp.StatusCode, res, nil
}

func HttpPostForm3(url string, form map[string]string) (int, []byte, error) {
	return HttpPostForm(url, form, 3)
}
func HttpPostForm(url string, form map[string]string, retryCount int) (int, []byte, error) {
	c := new(http.Client)
	req := request.NewRequest(c)
	req.Data = form
	resp, err := req.Post(url)
	if err != nil {
		if retryCount > 0 {
			return HttpPostForm(url, form, retryCount-1)
		}
		return 0, nil, err
	}
	defer resp.Body.Close()
	res, err := resp.Content()
	return resp.StatusCode, res, nil
}

func HttpQuery3(url string, param map[string]string, retryCount int) (int, []byte, error) {
	return HttpQuery(url, param, 3)
}
func HttpQuery(url string, param map[string]string, retryCount int) (int, []byte, error) {
	u, err := netUrl.Parse(url)
	if nil != err {
		return 0, nil, err
	}
	qs := u.Query()
	for k, v := range param {
		qs.Add(k, v)
	}
	u.RawQuery = qs.Encode()
	return HttpGet(u.String(), 3)
}
func HttpGet3(url string) (int, []byte, error) {
	return HttpGet(url, 3)
}
func HttpGet(url string, retryCount int) (int, []byte, error) {
	c := new(http.Client)
	req := request.NewRequest(c)
	resp, err := req.Get(url)
	if err != nil {
		if retryCount > 0 {
			return HttpGet(url, retryCount-1)
		}
		return 0, nil, err
	}
	defer resp.Body.Close()
	res, err := resp.Content()
	return resp.StatusCode, res, nil
}
func HttpDelete(url string, form map[string]string, retryCount int) (int, []byte, error) {
	u, err := netUrl.Parse(url)
	if nil != err {
		return 0, nil, err
	}
	qs := u.Query()
	for k, v := range form {
		qs.Add(k, v)
	}
	u.RawQuery = qs.Encode()
	c := new(http.Client)
	req := request.NewRequest(c)
	resp, err := req.Delete(u.String())
	if err != nil {
		if retryCount > 0 {
			return HttpGet(url, retryCount-1)
		}
		return 0, nil, err
	}
	defer resp.Body.Close()
	res, err := resp.Content()
	return resp.StatusCode, res, nil
}
