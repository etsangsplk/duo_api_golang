package main

import (
	//"encoding/json"
	//"fmt"
	//"net/http"
	//"net/url"
	"os"
	//"strings"
	"time"

	duoapi "github.com/duosecurity/duo_api_golang"
	//"github.com/duosecurity/duo_api_golang/authapi"
)

// Example basic. This shows how to use Duo Security Go API

func main() {
	iKey, _ := os.LookupEnv("iKey")
	sKey, _ := os.LookupEnv("sKey")
	apiHost, _ := os.LookupEnv("authMgrUrl")
	//factor, _ := os.LookupEnv("factor")
	userAgent := "Go Test Agent"
	 duoapi.NewDuoApi(
		iKey,
		sKey,
		apiHost,
		userAgent,
		duoapi.SetTimeout(1*time.Second),
		duoapi.SetInsecure())



}
