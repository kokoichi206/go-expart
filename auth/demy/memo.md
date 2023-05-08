## Authentication

- three ways
  - who you are (biometrics)
  - what you have (eg. atm card, key, phone)
  - what you know (eg. password,)

## Authorization

- says what you are allowed to do

## Basic authentication

``` sh
# Authorization: Basic dXNlcjpwYXNz
curl -u user:pass -v google.com

*   Trying 2404:6800:4004:80f::200e:80...
* Connected to google.com (2404:6800:4004:80f::200e) port 80 (#0)
* Server auth using Basic with user 'user'
> GET / HTTP/1.1
> Host: google.com
> Authorization: Basic dXNlcjpwYXNz
> User-Agent: curl/7.84.0
> Accept: */*
> ...

echo -n dXNlcjpwYXNz | base64 -d
user:pass%
```

- 可逆なので https でのみ使うこと！


## Bearer Token and Hmac

- Bearer Token
  - http spec, OAuth2
  - authorization header and keyword "bearer"
- hmac
  - Keyed-Hash Message Authentication Code
  - 偽の bearer token が送られてくるのを防ぐために hmac を使う
  - https://pkg.go.dev/crypto/hmac



## Links

- https://curlbuilder.com/
