``` sh
$ go doc fmt.Println

package fmt // import "fmt"

func Println(a ...any) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline
    is appended. It returns the number of bytes written and any write error
    encountered.
```

## TCP/IP

- Open Port
  - three-way handshake
- **Closed Port**
  - syn (client -> server)
  - rst (client <- server)
- Filtered Port (firewall etc...)
  - syn (client -> server)
    - **Timeout**
- port forwarding

``` sh
nc -lp 13337 -e /bin/bash
```


## Enlighs

- errant
  - (考え・行為など)常軌を逸した，逸脱した
- brevity
  - 簡潔
- tweak
  - 微調整する
- adequately
  - 十分に
