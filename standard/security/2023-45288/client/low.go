package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"net"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
)

const (
	// src/net/http/h2_bundle.go
	http2ClientPreface = "PRI * HTTP/2.0\r\n\r\nSM\r\n\r\n"
)

var (
	//go:embed dummy-key
	dummyKey string

	//go:embed dummy-value
	dummyValue string
)

func customHttp2Call(conn net.Conn) {
	var err error

	// RFC 7540 3.5 HTTP/2 Connection Preface
	conn.Write([]byte(http2ClientPreface))

	framer := http2.NewFramer(conn, conn)

	// framer.WriteRawFrame(http2.FrameSettings, 0, 0, []byte{})
	framer.WriteSettings(
		http2.Setting{
			ID:  http2.SettingEnablePush,
			Val: 0,
		},
		http2.Setting{
			ID:  http2.SettingInitialWindowSize,
			Val: 4194304,
		},
		http2.Setting{
			ID:  http2.SettingHeaderTableSize,
			Val: 4096,
		},
	)

	hbuf := bytes.NewBuffer([]byte{})
	henc := hpack.NewEncoder(hbuf)

	henc.WriteField(hpack.HeaderField{Name: ":authority", Value: "localhost:8080"})
	henc.WriteField(hpack.HeaderField{Name: ":method", Value: "GET"})
	henc.WriteField(hpack.HeaderField{Name: ":path", Value: "/"})
	henc.WriteField(hpack.HeaderField{Name: ":scheme", Value: "https"})
	henc.WriteField(hpack.HeaderField{Name: "custom-header", Value: "MyValue"})
	henc.WriteField(hpack.HeaderField{Name: "accept-encoding", Value: "gzip"})
	henc.WriteField(hpack.HeaderField{Name: "user-agent", Value: "Foo Bar"})

	// fmt.Printf("len(hbuf.Bytes()): %v\n", len(hbuf.Bytes()))
	err = framer.WriteHeaders(http2.HeadersFrameParam{
		// StreamID:      settingsFrame.StreamID,
		StreamID:      1,
		BlockFragment: hbuf.Bytes(),
		// EndHeaders:    true,
		// EndStream:     true,
		EndHeaders: false,
	})

	if err != nil {
		log.Fatal("write headers error: ", err)
	}

	// N := 1_000_000_000_000
	N := 1_000_000
	for idx := range N {
		// frame, err := framer.ReadFrame()
		// fmt.Printf("err: %v\n", err)
		// fmt.Printf("frame: %v\n", frame)
		writeHeader(framer, idx == N-1)
	}

	frames := make([]http2.Frame, 0)
	for {
		fmt.Println("----- for -----")
		frame, err := framer.ReadFrame()
		if err != nil {
			log.Fatal("read frame error: ", err)
		}
		frames = append(frames, frame)
		// if frame.Header().Flags.Has(http2.FlagSettingsAck) {
		// 	framer.WriteSettingsAck()
		// }
		if frame.Header().Flags.Has(http2.FlagHeadersEndStream) {
			fmt.Printf("head ended-----: %v\n", frame)
		}

		fmt.Printf("-------- frame: %v\n", frame)

		switch frame := frame.(type) {
		case *http2.DataFrame:
			log.Printf("data frame: %s\n", frame.Data())
			data := frame.Data()
			fmt.Printf("data: %v\n", data)
			fmt.Printf("string(data): %v\n", string(data))
		default:
		}

		if frame.Header().Type == http2.FrameData && frame.Header().Flags.Has(http2.FlagDataEndStream) {
			// end of stream !!!
			fmt.Printf("data ended-----: %v\n\n\n", frame)
			break
		}
	}

	for _, frame := range frames {
		switch frame := frame.(type) {
		case *http2.DataFrame:
			log.Printf("data frame: %s\n", frame.Data())
			data := frame.Data()
			fmt.Printf("data: %v\n", data)
			fmt.Printf("string(data): %v\n", string(data))
		case *http2.HeadersFrame:
			log.Printf("headers frame: %s\n", frame.Header())
		default:
			log.Printf("frame: %v\n", frame.Header())
		}
	}
}

func writeHeader(framer *http2.Framer, last bool) {
	hbuf := bytes.NewBuffer([]byte{})
	henc := hpack.NewEncoder(hbuf)
	henc.WriteField(hpack.HeaderField{Name: dummyKey, Value: dummyValue})
	// 逆にしたのも送る。
	henc.WriteField(hpack.HeaderField{Name: dummyValue, Value: dummyKey})

	// if err := framer.WriteHeaders(http2.HeadersFrameParam{
	// 	// StreamID:      settingsFrame.StreamID,
	// 	StreamID:      1,
	// 	BlockFragment: hbuf.Bytes(),
	// 	EndHeaders:    last,
	// 	EndStream:     last,
	// }); err != nil {
	// 	log.Fatal("write headers error: ", err)
	// }

	if err := framer.WriteContinuation(1, last, hbuf.Bytes()); err != nil {
		log.Fatal("write headers error: ", err)
	}
}
