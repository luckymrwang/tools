package common

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	client *http.Client

	DefaultTransport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          10,
		MaxIdleConnsPerHost:   2,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}
)

var DEFAULT_HTTP_TIMEOUT = 5 * time.Minute

func init() {
	client = &http.Client{Timeout: time.Duration(DEFAULT_HTTP_TIMEOUT), Transport: DefaultTransport}
}

func HttpGet(url string, header *P, param *P, tr *http.Transport) (body string, e error) {
	if query := GetUrlQuery(param); !IsEmpty(query) {
		url = fmt.Sprintf("%s?%s", url, query)
	}
	r, err := HttpGetBytes(url, header, param, tr)
	if err != nil {
		Error("HttpGet异常:", err.Error())
	}
	e = err
	body = string(r)
	return
}

func GetUrlQuery(params *P) (query string) {
	if params == nil {
		return
	}

	for key, val := range *params {
		if strings.Contains(key, "[]") {
			for _, v := range ToString(val) {
				query += fmt.Sprintf("%v=%v&", key, v)
			}
		} else {
			query += fmt.Sprintf("%v=%v&", key, val)
		}
	}

	if len(query) > 0 {
		query = query[0 : len(query)-1]
	}
	return
}

func HttpGetBytes(url string, header *P, param *P, tr *http.Transport) (body []byte, e error) {
	return HttpDo("GET", url, header, param, tr)
}

func HttpPost(url string, header *P, param *P) (body string, err error) {
	r, e := HttpDo("POST", url, header, param, nil)
	if e != nil {
		Error("HttpPost异常:", e.Error())
		body = e.Error()
		err = e
	} else {
		body = string(r)
	}
	return
}

func HttpDelete(url string, header *P, param *P) (body []byte, e error) {
	return HttpDo("DELETE", url, header, param, nil)
}

func HttpDo(method string, httpurl string, header *P, param *P, tr *http.Transport) (body []byte, err error) {
	var req *http.Request
	vs := url.Values{}
	if param != nil {
		for k, v := range *param {
			key := ToString(k)
			if IsMapArray(v) {
				vs.Set(key, JSONEncode(v))
			} else if IsArray(v) {
				a, _ := v.([]interface{})
				for i, iv := range a {
					if i == 0 {
						vs.Set(key, ToString(iv))
					} else {
						vs.Add(key, ToString(iv))
					}
				}
			} else {
				vs.Set(key, ToString(v))
			}
		}
	}
	method = strings.ToUpper(method)
	req, err = http.NewRequest(method, httpurl, strings.NewReader(vs.Encode()))
	if header != nil {
		for k, v := range *header {
			req.Header.Set(ToString(k), ToString(v))
		}
	}
	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	var resp *http.Response
	if strings.Contains(httpurl, "https") {
		httpsClient := *client
		if tr != nil {
			transport := *DefaultTransport
			transport.TLSClientConfig = tr.TLSClientConfig
			httpsClient.Transport = &transport
		}
		resp, err = httpsClient.Do(req)
	} else {
		resp, err = client.Do(req)
	}
	if err != nil {
		Error("HttpDo异常:", err.Error())
		return []byte(ToString(resp)), err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

func HttpDoWithCookie(method string, httpurl string, header *P, body []byte, cookies []*http.Cookie) (result []byte, err error) {
	client := &http.Client{Timeout: time.Duration(DEFAULT_HTTP_TIMEOUT)}
	req, err := http.NewRequest(strings.ToUpper(method), httpurl, bytes.NewBuffer(body))
	if header != nil {
		for k, v := range *header {
			req.Header.Set(ToString(k), ToString(v))
		}
	}
	if cookies != nil {
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		Error("HttpDoWithCookie异常:", err.Error())
		return []byte(ToString(resp)), err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	result, err = ioutil.ReadAll(resp.Body)
	return
}

func HttpCookies(method string, httpurl string, header *P, param *P) (cookie []*http.Cookie, err error) {
	client := &http.Client{Timeout: time.Duration(DEFAULT_HTTP_TIMEOUT)}
	var req *http.Request
	vs := url.Values{}
	if param != nil {
		for k, v := range *param {
			key := ToString(k)
			if IsMapArray(v) {
				vs.Set(key, JSONEncode(v))
			} else if IsArray(v) {
				a, _ := v.([]interface{})
				for i, iv := range a {
					if i == 0 {
						vs.Set(key, ToString(iv))
					} else {
						vs.Add(key, ToString(iv))
					}
				}
			} else {
				vs.Set(key, ToString(v))
			}
		}
	}
	method = strings.ToUpper(method)
	req, err = http.NewRequest(method, httpurl, strings.NewReader(vs.Encode()))
	if header != nil {
		for k, v := range *header {
			req.Header.Set(ToString(k), ToString(v))
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		Error("HttpCookies异常:", err.Error())
		return nil, err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	return resp.Cookies(), nil
}

func HttpRequest(method string, httpurl string, param *P) (body []byte, cookies []*http.Cookie, err error) {
	client := &http.Client{Timeout: time.Duration(DEFAULT_HTTP_TIMEOUT)}
	var req *http.Request
	vs := url.Values{}
	if param != nil {
		for k, v := range *param {
			key := ToString(k)
			if IsMapArray(v) {
				vs.Set(key, JSONEncode(v))
			} else if IsArray(v) {
				a, _ := v.([]interface{})
				for i, iv := range a {
					if i == 0 {
						vs.Set(key, ToString(iv))
					} else {
						vs.Add(key, ToString(iv))
					}
				}
			} else {
				vs.Set(key, ToString(v))
			}
		}
	}
	method = strings.ToUpper(method)
	req, err = http.NewRequest(method, httpurl, strings.NewReader(vs.Encode()))
	resp, err := client.Do(req)
	if err != nil {
		Error("HttpRequest异常:", err.Error())
		return []byte(ToString(resp)), nil, err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	cookies = resp.Cookies()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

func HttpPostBody(url string, header *P, body []byte) (string, error) {
	client := &http.Client{Timeout: DEFAULT_HTTP_TIMEOUT}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if header != nil {
		for k, v := range *header {
			req.Header.Set(ToString(k), ToString(v))
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		Error("HttpPostBody异常:", err.Error())
		return ToString(resp), err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	b, err := ioutil.ReadAll(resp.Body)
	return string(b), err
}

func UrlEncoded(str string) (string, error) {
	str = strings.Replace(str, "%", "%25", -1)
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func Upload(url, file string) (body []byte, err error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	// Add your file
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()
	fw, err := w.CreateFormFile("bin", file)
	if err != nil {
		return
	}
	if _, err = io.Copy(fw, f); err != nil {
		return
	}
	// Add the other fields
	if fw, err = w.CreateFormField("key"); err != nil {
		return
	}
	if _, err = fw.Write([]byte("KEY")); err != nil {
		return
	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		Error("Upload异常:", err.Error())
		return []byte(ToString(res)), err
	}
	defer func() {
		if res != nil {
			res.Body.Close()
		}
	}()
	body, err = ioutil.ReadAll(res.Body)
	return
}
