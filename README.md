# armorapiauth

This package is a canned Armor Secure Hosting API authorization key generator.  It simply returns "FH-AUTH $bearer_token".

# Requirements
The only real requirements is to set your user/password in the environment, you can do this at run time via.

```MasterPass="password" ArmorUser="username" go run *.go (or the upstream go proggy if using this as a library.)```

I would encourage you to instrument your own way of securely getting your creds.  I'm simply grabbing them via os.Env(VAR) as everyone's shop will have their own way of storage and retreval of secure creds.

