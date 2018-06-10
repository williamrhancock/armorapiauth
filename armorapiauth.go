package aromorapiauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	username   = os.Getenv("MasterUser")
	password   = os.Getenv("ArmorPass")
	accept     = "application/json"
	fetchError = false
)

type access struct {
	URL         string
	User        string
	Pass        string
	code        string
	accessToken string
	idToken     string
}
type authentication struct {
	Code        string `json:"code,omitempty"`
	RedirectURI string `json:"redirect_uri,omitempty"`
	Success     bool   `json:"success,omitempty"`
}
type token struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
	IDToken     string `json:"id_token,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
}

var auth authentication
var toke token
var a access

// GenBearer is the only exported function: BearerToken:= GenBearer()
func GenBearer() string {
	if len(username) == 0 {
		a.die("You must declare an evn variable for ArmorPass, and MasterPass")
	}

	fetchError = false
	a = access{
		URL:  "https://api.armor.com/",
		User: username,
		Pass: password,
	}

	if ok := a.authpost("auth/authorize",
		`{
		"username": "`+username+`",
		"password": "`+password+`"
	  }`); ok != nil {
		a.die(ok)
	} else {
		a.code = auth.Code
	}

	if ok := a.authpost("auth/token",
		`{
		"code":"`+auth.Code+`",
		"grant_type":"authorization_code"
	  }`); ok != nil {
		a.die(ok)
	} else {
		a.accessToken = toke.AccessToken
		a.idToken = toke.IDToken
	}

	return "FH-AUTH " + toke.AccessToken
}

func (a access) authpost(path, payload string) error {
	buf := bytes.NewBufferString(payload)

	resp, fetchError := http.Post(a.URL+path, "application/json", buf)
	if fetchError != nil {
		err := fmt.Errorf("Failed request: %s%s ", a.URL, path)
		return err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	a.parse(body)

	if resp.Status != "200 OK" {
		err := fmt.Errorf("Failed request: %v: %s%s ", resp.Status, a.URL, path)
		return err
	}

	return nil
}

func (a access) parse(s []byte) {
	if err := json.Unmarshal(s, &auth); err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(s, &toke); err != nil {
		log.Fatal(err)
	}
}

func (a access) die(p ...interface{}) {
	for _, v := range p {
		log.Printf("%#v\n", v)
	}
	log.Printf("%#v\n", a)
	log.Fatal("Ungraceful death.")
}
