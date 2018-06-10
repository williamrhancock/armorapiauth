package aromorapiauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	username    = os.Getenv("MasterUser")
	password    = os.Getenv("ArmorPass")
	accept      = "application/json"
	fetchError  = false
	code        string
	accessToken string
	idToken     string
)

type access struct {
	URL         string
	User        string
	Pass        string
	code        string
	accessToken string
	idToken     string
	laststatus  string
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

func armorapiauth() string {
	if len(username) == 0 {
		a.print("You must declare an evn variable for ArmorPass, and MasterPass")
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
		a.print(ok)
		os.Exit(99)
	} else {
		a.code = auth.Code
	}

	if ok := a.authpost("auth/token",
		`{
		"code":"`+auth.Code+`",
		"grant_type":"authorization_code"
	  }`); ok != nil {
		a.print(ok)
		os.Exit(99)
	} else {
		a.accessToken = toke.AccessToken
		a.idToken = toke.IDToken
	}

	//log.Print("Successful authentication.")
	// Use to check the access struct for full data
	//fmt.Printf("%#v\n", a)
	return "FH-AUTH " + toke.IDToken
}

func (a access) print(p ...interface{}) {
	for _, v := range p {
		log.Printf("%#v\n", v)
	}
	log.Printf("%#v\n", a)
	log.Fatal("Ungraceful death.")
}

func (a access) authpost(path, payload string) error {
	buf := bytes.NewBufferString(payload)

	resp, fetchError := http.Post(a.URL+path, "application/json", buf)
	if fetchError != nil {
		err := fmt.Errorf("Failed request: %s%s ", a.URL, path)
		_ = err
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

func graceful() { fmt.Println(time.Now(), "End"); os.Exit(0) }

func reuse() {}
func new()   {}
