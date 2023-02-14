## WebSocket

- → Connection: Upgrade
- ← HTTP 101 Switching Protocols
- WebSocket
- → Close

## When ws

- chats-app
- games
- real-time app

## Events

- Close
- Error
- Message
- Open

## Auth

- Regular HTTP auth
  - one-time password, ticket -> parameter
  - this approach
- allow connect and authenticate

## cert

``` sh
openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.crt -batch -days 3650
```
