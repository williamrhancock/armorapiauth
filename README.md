# armorapiauth

This package is a canned Armor Secure Hosting API authorization key generator.  It simply returns "FH-AUTH $bearer_token".

# Requirements
The only real requirements is to set your user/password in the environment, you can do this at run time via.

```MasterPass="password" ArmorUser="username" go run *.go (or the upstream go proggy if using this as a library.)```

I would encourage you to instrument your own way of securely getting your creds.  I'm simply grabbing them via os.Env(VAR) as everyone's shop will have their own way of storage and retreval of secure creds.

# Example
For this example you need your account number for context, it can be found on the amp portal under "Account"->"Overview".

```
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/williamrhancock/armorapiauth"
)

var account = "YourArmorCompanyNumber-4 digits"

func main() {
	client := &http.Client{}

	token := aromorapiauth.GenBearer()

	request, error := http.NewRequest("GET", "https://api.armor.com/users", nil)

	request.Header.Set("Authorization", token)
	request.Header.Set("X-Account-Context", account)

	responce, error := client.Do(request)
	if error != nil {
		log.Fatal(error)
	}

	defer responce.Body.Close()
	body, _ := ioutil.ReadAll(responce.Body)
	fmt.Println(string(body))
}
```

This is effectly return all users in the context of the company, the struct in this case looks like this.


```
type Users struct {
	Culture            string    `json:"culture,omitempty"`
	Email              string    `json:"email,omitempty"`
	FirstName          string    `json:"firstName,omitempty"`
	ID                 float64   `json:"id,omitempty"`
	IsMfaEnabled       bool      `json:"isMfaEnabled,omitempty"`
	LastLogin          string    `json:"lastLogin,omitempty"`
	LastModified       string    `json:"lastModified,omitempty"`
	LastName           string    `json:"lastName,omitempty"`
	MfaMode            string    `json:"mfaMode,omitempty"`
	MfaPinMode         string    `json:"mfaPinMode,omitempty"`
	MustChangePassword bool      `json:"mustChangePassword,omitempty"`
	PasswordLastSet    string    `json:"passwordLastSet,omitempty"`
	Permissions        []float64 `json:"permissions,omitempty"`
	PhonePrimary       struct {
		CountryCode    float64     `json:"countryCode,omitempty"`
		CountryIsoCode interface{} `json:"countryIsoCode,omitempty"`
		Number         string      `json:"number,omitempty"`
		PhoneExt       interface{} `json:"phoneExt,omitempty"`
	} `json:"phonePrimary,omitempty"`
	Status   string `json:"status,omitempty"`
	Timezone string `json:"timezone,omitempty"`
	Title    string `json:"title,omitempty"`
}```