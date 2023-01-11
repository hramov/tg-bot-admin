package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

func SendResponse[T any](code int, data T, w http.ResponseWriter) {
	bytes, err := json.Marshal(data)
	if err != nil {
		SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(bytes)
}

func SendError(code int, err string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write([]byte("{\"error\": \"" + err + "\"}"))
}
