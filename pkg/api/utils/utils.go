package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func GetBody[T any](r *http.Request) (T, error) {
	var data T

	rawData, err := io.ReadAll(r.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(r.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(rawData, &data)

	if err != nil {
		return data, err
	}

	return data, nil
}

func SendResponse[T any](code int, data T, w http.ResponseWriter) {
	bytes, err := json.Marshal(data)
	if err != nil {
		SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}

func SendError(code int, err string, w http.ResponseWriter) {
	w.WriteHeader(code)
	_, _ = w.Write([]byte(err))
}

func GetTokenFromRequest(req *http.Request) (string, error) {
	auth := req.Header.Get("authorization")
	if auth != "" {
		cred := strings.Split(auth, " ")
		if len(cred) > 1 && cred[0] == "Bearer" {
			if cred[1] != "" {
				return cred[1], nil
			}
			return "", fmt.Errorf("no token")
		}
		return "", fmt.Errorf("wrong auth header format")
	}
	return "", fmt.Errorf("wo auth header")
}

func GetReqResFromContext(c *gin.Context) (*http.Request, *http.Response) {
	return c.Request, c.Request.Response
}

func IsEqualClientIp(serviceIp, clientIp string) (bool, error) {
	ip := ""
	u, err := url.Parse(serviceIp)
	if err != nil {
		return false, err
	}

	if u.Hostname() == "localhost" {
		ip = "127.0.0.1"
	} else {
		ip = u.Hostname()
	}

	return ip == clientIp, nil

}

func MaintainRequest(requestCtx context.Context, cancel context.CancelFunc) {
	<-requestCtx.Done()
	cancel()
}
