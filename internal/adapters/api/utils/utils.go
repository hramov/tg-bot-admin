package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func GetBody[T any](c *gin.Context) (T, error) {
	var data T

	reqBody, exists := c.Get("body")

	if !exists {
		rawData, err := io.ReadAll(c.Request.Body)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
			}
		}(c.Request.Body)
		if err != nil {
			return data, err
		}

		err = json.Unmarshal(rawData, &data)

		if err != nil {
			return data, err
		}

		c.Set("body", data)
		return data, nil
	}

	data = reqBody.(T)
	return data, nil
}

func SendResponse[T any](code int, data T, c *gin.Context) {
	c.JSON(code, &gin.H{
		"version": "0.0.1",
		"data":    data,
	})
}

func SendError[T any](code int, err T, c *gin.Context) {
	c.AbortWithStatusJSON(code, err)
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

func GetTokenFromContext(c *gin.Context) (string, error) {
	req, _ := GetReqResFromContext(c)
	return GetTokenFromRequest(req)
}

func GetReqResFromContext(c *gin.Context) (*http.Request, *http.Response) {
	return c.Request, c.Request.Response
}

func GetParam(name string, c *gin.Context) string {
	return c.Param(name)
}

func GetQuery(name string, c *gin.Context) string {
	return c.Query(name)
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
