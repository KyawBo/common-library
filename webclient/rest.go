package webclient

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/KyawBo/common-library/configuration"

	"github.com/KyawBo/common-library/constants"

	"github.com/KyawBo/common-library/json"
)

type webclientRest struct {
	httpClient *http.Client
}

type WebclientRestInterface interface {
	CallREST(method string, endpoint string, headers map[string]string, body interface{}, result interface{}, resultErr interface{}) (bool, int, error)
}

func NewWebclientRest() WebclientRestInterface {
	timeout := constants.WEBCLIENT_TIMEOUT
	if configuration.Config.WsClient.Timeout > 0 {
		timeout = configuration.Config.WsClient.Timeout
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   time.Duration(timeout) * time.Millisecond,
		Transport: tr,
	}

	return &webclientRest{
		httpClient: client,
	}
}

func (w *webclientRest) CallREST(method string, endpoint string, headers map[string]string, body interface{}, result interface{}, resultErr interface{}) (bool, int, error) {
	// headerB, err := json.Json.Marshal(headers)
	// if err != nil {
	// 	return false, 0, err
	// }

	bodyB, err := json.Json.Marshal(body)
	if err != nil {
		return false, 0, err
	}

	fmt.Println(body)

	// logFields := map[string]interface{}{
	// 	constants.LOG_MSG_KEY: "rest_request",
	// 	"endpoint":            endpoint,
	// 	"method":              method,
	// 	"headers":             log.FilterLog(string(headerB)),
	// 	"body":                log.FilterLog(string(bodyB)),
	// }

	// log.CommonLog(c, traceData, logFields)

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(bodyB))
	if err != nil {
		return false, 0, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := w.httpClient.Do(req)
	if err != nil {
		return false, 0, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, 0, err
	}

	isSuccess := resp.StatusCode > 199 && resp.StatusCode < 300

	fmt.Println(string(respBody))

	// logFields = map[string]interface{}{
	// 	constants.LOG_MSG_KEY: "rest_response",
	// 	"body":                log.FilterLog(string(respBody)),
	// 	"is_success":          isSuccess,
	// 	"status_code":         resp.StatusCode,
	// }

	//log.CommonLog(c, traceData, logFields)

	if isSuccess {
		if result != nil {
			err = json.Json.Unmarshal(respBody, result)
		}
	} else {
		if resultErr != nil {
			err = json.Json.Unmarshal(respBody, resultErr)
		}
	}

	if err != nil {
		return false, 0, err
	}

	return isSuccess, resp.StatusCode, nil
}
