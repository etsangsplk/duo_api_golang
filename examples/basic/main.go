package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/duosecurity/duo_api_golang"
	"github.com/duosecurity/duo_api_golang/authapi"
)

// Example basic. This shows how to use Duo Security Go API

func buildAuthApi(url string, proxy func(*http.Request) (*url.URL, error)) *authapi.AuthApi {
	ikey := "eyekey"
	skey := "esskey"
	host := strings.Split(url, "//")[1]
	userAgent := "GoTestClient"
	return authapi.NewAuthApi(*duoapi.NewDuoApi(ikey,
		skey,
		host,
		userAgent,
		duoapi.SetTimeout(1*time.Second),
		duoapi.SetInsecure(),
		duoapi.SetProxy(proxy)))
}

func main() {

}
