package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GitlabToken struct {
	AccessToken string `json:"access_token"`
}

// 代理 gitlab 接口

func Login(w http.ResponseWriter, r *http.Request) {
	v := url.Values{}
	v.Add("client_id", "712f0fa4c635f752f4efb99e7c0a8f90f6db97f8bb56f0f7fef3325e4b7659c5")
	v.Add("redirect_uri", "http://localhost:4000/gitlab-callback")
	v.Add("response_type", "code")
	v.Add("state", "12345")
	v.Add("scope", "api")
	http.Redirect(w, r, "https://gitlab.com/oauth/authorize?"+v.Encode(), 302)
}

func Callback(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")
	v := url.Values{}
	v.Add("client_id", "712f0fa4c635f752f4efb99e7c0a8f90f6db97f8bb56f0f7fef3325e4b7659c5")
	v.Add("client_secret", "3656bf06ebd645b21ca7088e99bfb40aa583ba5d4006c905a23b492ebc2b0896")
	v.Add("redirect_uri", "http://localhost:4000/gitlab-callback")
	v.Add("code", code)
	v.Add("grant_type", "authorization_code")

	resp, _ := http.PostForm("https://gitlab.com/oauth/token", v)
	body, _ := io.ReadAll(resp.Body)

	gitlabToken := &GitlabToken{}
	err := json.Unmarshal([]byte(body), gitlabToken)
	if err != nil {
		fmt.Printf("序列化错误%v\n", err)
	}
	fmt.Println(gitlabToken)

	c1 := http.Cookie{Name: "gitlab-token", Value: gitlabToken.AccessToken, HttpOnly: false, MaxAge: 7200, Path: "/", SameSite: 4, Secure: true}
	w.Header().Add("Set-Cookie", c1.String())

	http.Redirect(w, r, "http://localhost:3000/", 302)

}
