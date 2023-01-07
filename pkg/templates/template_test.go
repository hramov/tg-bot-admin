package template

import "testing"

func Test_GetTemplate(t *testing.T) {
	_, err := GetTemplate(Http, "login.html")
	if err != nil {
		t.Error(err.Error())
	}
}
