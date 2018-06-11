# armorapiauth

This package is a canned Armor Secure Hosting API authorization key generator.  It simply returns "FH-AUTH $bearer_token".

# Requirements
None, it's totally native.

I would encourage you to instrument your own way of securely getting your creds.  I'm simply grabbing them via os.Getenv(VAR) as a fall back but you can pass them to `func GenBearer(string, string) string` if you want. Everyone's shop will have their own way of storage and retreval of secure creds that will be best for them.

# Example
For this example you need your account number for context, it can be found on the amp portal (G4) under "Account"->"Overview".

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
var user = "username"
var pass = "password"

func main() {
	client := &http.Client{}

	// This is the only line you need to add to your code
	// to get all of the access request back and forth out 
	// of the way.  Once you have it, just pass it to the 
	// "Authorization" header (2 lines below).
	token := aromorapiauth.GenBearer(user, pass)

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

This returns all users in the context of the company.  Incidentally the struct in this case looks like this.


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