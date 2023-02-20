## Go

- Goroutines
  - light weight threads of execution managed by the go runtime

## WebRTC

Web Real-Time Communication

enable Web applications and sites to capture and optionally stream audio and/or video media as well as to exchange arbitrary data.

something like Google Meet using standard web APIs.

### Lifecycle

- Offers and Answers between 2 peers
  - This negotiation is called Signalling
  - expressed using the Session Description Protocol (SDP)
- p2p connectivity information: ICE Candidates

## Signalling, SDP

- session description -> SDP
- websocket
- sdp
  - local description: itself
  - remote description: other end of the call
- p2p connection
- ICE: interactive connectivity establishment
- stun: session traversal utilities for NAT





## Links

- [WebRTC の仕組み](https://cloudapi.kddi-web.com/magazine/webrtc/understood-webrtc-mechanism)
- [Session 2 : Building Video Chat Apps using WebRTC and Golang](https://www.youtube.com/watch?v=JTIm3ChI-6w&ab_channel=GDSCKIIT)
- [rfc6455: websocket](https://www.rfc-editor.org/rfc/rfc6455)
- [stun server list](https://gist.github.com/sagivo/3a4b2f2c7ac6e1b5267c2f1f59ac6c6b)
