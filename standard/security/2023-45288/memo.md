## issue

https://github.com/golang/go/issues/65051

https://github.com/golang/go/issues/65387


## go src 修正内容

https://go-review.googlesource.com/c/go/+/576076/2/src/net/http/h2_bundle.go




## curl で http2 メモ

``` sh
#    -k, --insecure
curl -vk --http2 https://localhost:8080


$ man curl
       -k, --insecure
              (TLS SFTP SCP) By default, every secure connection curl makes is verified to be secure
              before the transfer takes place. This option makes curl skip the verification step and
              proceed without checking.

              When this option is not used for protocols using TLS, curl verifies the server's TLS
              certificate before it continues: that the certificate contains the right name which
              matches the host name used in the URL and that the certificate has been signed by a CA
              certificate present in the cert store.  See this online resource for further details:
               https://curl.se/docs/sslcerts.html

              For SFTP and SCP, this option makes curl skip the known_hosts verification.
              known_hosts is a file normally stored in the user's home directory in the .ssh
              subdirectory, which contains host names and their public keys.

              WARNING: using this option makes the transfer insecure.

              Example:
               curl --insecure https://example.com

              See also --proxy-insecure, --cacert and --capath.
```


``` sh
❯ curl -vk --http2 https://localhost:8080

*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
* ALPN: offers h2
* ALPN: offers http/1.1
* (304) (OUT), TLS handshake, Client hello (1):
* (304) (IN), TLS handshake, Server hello (2):
* (304) (IN), TLS handshake, Unknown (8):
* (304) (IN), TLS handshake, Certificate (11):
* (304) (IN), TLS handshake, CERT verify (15):
* (304) (IN), TLS handshake, Finished (20):
* (304) (OUT), TLS handshake, Finished (20):
* SSL connection using TLSv1.3 / AEAD-CHACHA20-POLY1305-SHA256
* ALPN: server accepted h2
* Server certificate:
*  subject: C=AU; ST=Some-State; O=Internet Widgits Pty Ltd
*  start date: Apr  7 12:15:07 2024 GMT
*  expire date: Apr  7 12:15:07 2025 GMT
*  issuer: C=AU; ST=Some-State; O=Internet Widgits Pty Ltd
*  SSL certificate verify result: self signed certificate (18), continuing anyway.
* Using HTTP2, server supports multiplexing
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* h2h3 [:method: GET]
* h2h3 [:path: /]
* h2h3 [:scheme: https]
* h2h3 [:authority: localhost:8080]
* h2h3 [user-agent: curl/7.84.0]
* h2h3 [accept: */*]
* Using Stream ID: 1 (easy handle 0x14000cc00)
> GET / HTTP/2
> Host: localhost:8080
> user-agent: curl/7.84.0
> accept: */*
> 
* Connection state changed (MAX_CONCURRENT_STREAMS == 250)!
< HTTP/2 200 
< content-type: text/plain; charset=utf-8
< content-length: 20
< date: Sun, 07 Apr 2024 14:30:32 GMT
< 
* Connection #0 to host localhost left intact
Hello, HTTP/2 world!
```

## CONTINUATION

[rfc7540-6.10](https://datatracker.ietf.org/doc/rfc7540/)

> A CONTINUATION frame MUST be preceded by a HEADERS, PUSH_PROMISE or CONTINUATION frame without the END_HEADERS flag set. A recipient that observes violation of this rule MUST respond with a connection error (Section 5.4.1) of type PROTOCOL_ERROR.

CONTINUATION の前には END_HEADERS のセットされてない Header とかが必要！

## Links

- https://datatracker.ietf.org/doc/html/rfc9204

