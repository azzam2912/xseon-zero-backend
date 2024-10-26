package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"xseon-zero/lib/log"
)

const (
	LibHTTPRequestInfo  = "INFO.STATEMENT.LIB.HTTP.REQUEST"
	LibHTTPResponseInfo = "INFO.STATEMENT.LIB.HTTP.RESPONSE"
	LibHTTPRequestErr   = "ERR.STATEMENT.LIB.HTTP.REQUEST"
	LibHTTPResponseErr  = "ERR.STATEMENT.LIB.HTTP.RESPONSE"

	LogInfoSeverity  = "INFO"
	LogErrorSeverity = "ERROR"
)

func requestToMap(req *http.Request) map[string]interface{} {
	reqCopy := req.Clone(req.Context())
	if req.Body != nil {
		var buf bytes.Buffer
		io.Copy(&buf, req.Body)
		reqCopy.Body = io.NopCloser(bytes.NewReader(buf.Bytes()))
		req.Body = io.NopCloser(bytes.NewReader(buf.Bytes()))
	}

	var result map[string]interface{}
	body, _ := io.ReadAll(reqCopy.Body)
	_ = json.Unmarshal(body, &result)
	return result
}

func responseToMap(resp *http.Response) map[string]interface{} {
	var result map[string]interface{}
	var buf bytes.Buffer
	tee := io.TeeReader(resp.Body, &buf)
	body, _ := io.ReadAll(tee)
	_ = json.Unmarshal(body, &result)
	resp.Body = io.NopCloser(&buf)
	return result
}

func MakeRequest(url string, body interface{}, headers map[string]string, method string, formFileBody *bytes.Buffer) ([]byte, int, error) {
	var requestBody *bytes.Buffer
	if formFileBody != nil {
		requestBody = formFileBody
	} else {
		reqBody, err := json.Marshal(body)
		if err != nil {
			return nil, 0, err
		}
		requestBody = bytes.NewBuffer(reqBody)
	}
	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		log.WriteLog(LogErrorSeverity, LibHTTPRequestErr, err.Error(), "", nil)
		return nil, 0, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	log.WriteLog(LogInfoSeverity, LibHTTPRequestInfo, "Request", "", requestToMap(req))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			log.WriteLog(LogErrorSeverity, LibHTTPResponseErr, "Timeout", "", nil)
			return nil, 504, err
		}

		log.WriteLog(LogErrorSeverity, LibHTTPResponseErr, err.Error(), "", nil)
		return nil, 500, err
	}

	if resp == nil {
		log.WriteLog(LogErrorSeverity, LibHTTPResponseErr, "No response received", "", map[string]interface{}{
			"target_url": url,
		})
		return nil, 500, errors.New("no response received")
	}

	log.WriteLog(LogInfoSeverity, LibHTTPResponseInfo, "Response", "", responseToMap(resp))
	defer resp.Body.Close()

	var responseBody []byte
	responseBody, err = io.ReadAll(resp.Body)
	if err != nil {
		log.WriteLog(LogErrorSeverity, LibHTTPResponseErr, err.Error(), "", nil)
		return nil, 500, err
	}
	return responseBody, resp.StatusCode, nil
}
