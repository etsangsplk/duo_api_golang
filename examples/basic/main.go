package main

import (
	//"encoding/json"
	"fmt"
	//"net/http"
	//"net/url"
	"log"
	"os"
	//"strings"
	"time"

	duoapi "github.com/duosecurity/duo_api_golang"
	"github.com/duosecurity/duo_api_golang/authapi"
)

func getKeyOrDefault(key string, defVal string) (string, bool) {
	val, ok := os.LookupEnv(key)
	if ok {
		return val, ok
	}
	return defVal, ok
}

// Dam you need to ask for sales to enable

func main() {
	// Must be partner API on application list in duo secuiryt for application to protect
	iKey, _ := getKeyOrDefault("iKey", "DIM7ZWB0986Y84TG1DI9")
	sKey, _ := getKeyOrDefault("sKey", "b0rEroFKOZXrrbtwUs5WDtvZD1rnusi8EuWfxCHN")
	apiHost, _ := getKeyOrDefault("duoHost", "api-fb2615f6.duosecurity.com")
	factor := "passcode"
	//uID := "4213-5772-06"
	uName := "sc_admin_admin"
	passCode := "043415"
	//apiUrl, _ := getKeyOrDefault("authMgrUrl", "/auth/v2/auth")
	//apiServer, err := url.Parse(apiHost+apiUrl)

	userAgent := "Go Test Agent"
	duoApiClient := duoapi.NewDuoApi(
		iKey,
		sKey,
		apiHost,
		userAgent,
		duoapi.SetTimeout(1*time.Second),
		duoapi.SetInsecure())

	authClient := authapi.NewAuthApi(*duoApiClient)

	resP, err := authClient.Ping()
	fmt.Printf("ping %p response: %#v \n", resP, resP)

	if err != nil {
		fmt.Printf("ping error %#v \n", err)
		log.Fatal(err)
	}

	resA, err := authClient.Auth(
		factor,
		//authapi.AuthUserId(uID),
		authapi.AuthUsername(uName),
		//authapi.AuthIpAddr("40.40.40.10"),
		//authapi.AuthAsync(),
		authapi.AuthPasscode(passCode),
	)
	//resAjson, _ := json.Marshal(resA)

	fmt.Printf("auth %p response*******: %#v \n", resA, resA)

	fmt.Printf("authResponse Stat Result: %#v \n Code %#v \n Message %#v \n Message_Detail %#v\n", resA.StatResult,
		*resA.Code,
		*resA.Message,
		*resA.Message_Detail,
	)

	fmt.Printf("authResponse*******: %#v \n", resA.Response)

	if err != nil {
		fmt.Printf("auth error %#v \n", err)
		log.Fatal(err)
	}

}
