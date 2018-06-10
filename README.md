# armorapiauth

This package is a canned Armor Secure Hosting API authorization key generator.  It simply returns "FH-AUTH $bearer_token".

# Requirements
The only real requirements is to set your user/password in the environment, you can do this at run time via.

```MasterPass="password" ArmorUser="username" go run *.go (or the upstream go proggy if using this as a library.)```