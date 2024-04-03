## 環境確認

``` sh
❯ which -a go
/opt/homebrew/bin/go
/usr/local/go/bin/go
❯ cz | grep -i go
export PATH="$PATH:/usr/local/go/bin/go"
export PATH="$PATH:$HOME/go/bin"
export GOPATH="$HOME/go"
export PATH="$PATH:$(go env GOPATH)/bin"


❯ /opt/homebrew/bin/go version
go version go1.21.7 darwin/arm64
❯ /usr/local/go/bin/go version
go version go1.20.7 darwin/arm64
```

## 初回実行時エラー

``` sh
$ go-licenses report github.com/android-project-46group/api-server

E0403 23:21:58.468922   59961 library.go:117] Package context does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.471205   59961 library.go:117] Package database/sql does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.471222   59961 library.go:117] Package encoding/json does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.471227   59961 library.go:117] Package errors does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.471231   59961 library.go:117] Package fmt does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.478501   59961 library.go:117] Package io does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.478522   59961 library.go:117] Package path does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
W0403 23:21:58.478529   59961 library.go:101] "runtime" contains non-Go code that can't be inspected for further dependencies:
/opt/homebrew/Cellar/go/1.21.7/libexec/src/runtime/asm_ppc64x.h
/opt/homebrew/Cellar/go/1.21.7/libexec/src/runtime/funcdata.h
/opt/homebrew/Cellar/go/1.21.7/libexec/src/runtime/go_tls.h
/opt/homebrew/Cellar/go/1.21.7/libexec/src/runtime/textflag.h
/opt/homebrew/Cellar/go/1.21.7/libexec/src/runtime/tls_arm64.h
/opt/homebrew/Cellar/go/1.21.7/libexec/src/runtime/asm.s
/opt/homebrew/Cellar/go/1.21.7/libexec/src/runtime/asm_arm64.s
/opt/homebrew/Cellar/go/1.21.7/libexec/src/runtime/atomic_arm64.s
/opt/homebrew/Cellar/go/1.21.7/libexec/src/runtime/duff_arm64.s
/opt/homebrew/Cellar/go/1.21.7/libexec/src/runtime/memclr_arm64.s
/opt/homebrew/Cellar/go/1.21.7/libexec/src/runtime/memmove_arm64.s
/opt/homebrew/Cellar/go/1.21.7/libexec/src/runtime/preempt_arm64.s
/opt/homebrew/Cellar/go/1.21.7/libexec/src/runtime/rt0_darwin_arm64.s
/opt/homebrew/Cellar/go/1.21.7/libexec/src/runtime/sys_darwin_arm64.s
/opt/homebrew/Cellar/go/1.21.7/libexec/src/runtime/tls_arm64.s
E0403 23:21:58.478559   59961 library.go:117] Package runtime does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.478571   59961 library.go:117] Package strconv does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.478575   59961 library.go:117] Package strings does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.482749   59961 library.go:117] Package bytes does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.482767   59961 library.go:117] Package database/sql/driver does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
W0403 23:21:58.487382   59961 library.go:101] "reflect" contains non-Go code that can't be inspected for further dependencies:
/opt/homebrew/Cellar/go/1.21.7/libexec/src/reflect/asm_arm64.s
E0403 23:21:58.487407   59961 library.go:117] Package reflect does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.487411   59961 library.go:117] Package time does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
W0403 23:21:58.491761   59961 library.go:101] "crypto/md5" contains non-Go code that can't be inspected for further dependencies:
/opt/homebrew/Cellar/go/1.21.7/libexec/src/crypto/md5/md5block_arm64.s
E0403 23:21:58.491781   59961 library.go:117] Package crypto/md5 does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.494935   59961 library.go:117] Package crypto/rand does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
W0403 23:21:58.494953   59961 library.go:101] "crypto/sha1" contains non-Go code that can't be inspected for further dependencies:
/opt/homebrew/Cellar/go/1.21.7/libexec/src/crypto/sha1/sha1block_arm64.s
E0403 23:21:58.494959   59961 library.go:117] Package crypto/sha1 does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.494963   59961 library.go:117] Package encoding/binary does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.494967   59961 library.go:117] Package encoding/hex does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.495306   59961 library.go:117] Package hash does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.495323   59961 library.go:117] Package net does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.495331   59961 library.go:117] Package os does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.495337   59961 library.go:117] Package sync does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.505057   59961 library.go:117] Package regexp does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.505108   59961 library.go:117] Package unicode does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.505112   59961 library.go:117] Package unicode/utf8 does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
W0403 23:21:58.505116   59961 library.go:101] "math" contains non-Go code that can't be inspected for further dependencies:
/opt/homebrew/Cellar/go/1.21.7/libexec/src/math/dim_arm64.s
/opt/homebrew/Cellar/go/1.21.7/libexec/src/math/exp_arm64.s
/opt/homebrew/Cellar/go/1.21.7/libexec/src/math/floor_arm64.s
/opt/homebrew/Cellar/go/1.21.7/libexec/src/math/modf_arm64.s
E0403 23:21:58.505123   59961 library.go:117] Package math does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.505390   59961 library.go:117] Package sort does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.505397   59961 library.go:117] Package math/rand does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
W0403 23:21:58.505401   59961 library.go:101] "sync/atomic" contains non-Go code that can't be inspected for further dependencies:
/opt/homebrew/Cellar/go/1.21.7/libexec/src/sync/atomic/asm.s
E0403 23:21:58.505405   59961 library.go:117] Package sync/atomic does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.566375   59961 library.go:117] Package html/template does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.566421   59961 library.go:117] Package io/ioutil does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.566426   59961 library.go:117] Package os/exec does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.566430   59961 library.go:117] Package path/filepath does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.641753   59961 library.go:117] Package net/http does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.644420   59961 library.go:117] Package encoding/csv does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
W0403 23:21:58.647922   59961 library.go:101] "golang.org/x/sys/unix" contains non-Go code that can't be inspected for further dependencies:
/Users/kokoichi/go/pkg/mod/golang.org/x/sys@v0.4.0/unix/asm_bsd_arm64.s
/Users/kokoichi/go/pkg/mod/golang.org/x/sys@v0.4.0/unix/zsyscall_darwin_arm64.s
W0403 23:21:58.666729   59961 library.go:101] "syscall" contains non-Go code that can't be inspected for further dependencies:
/opt/homebrew/Cellar/go/1.21.7/libexec/src/syscall/asm_darwin_arm64.s
/opt/homebrew/Cellar/go/1.21.7/libexec/src/syscall/zsyscall_darwin_arm64.s
E0403 23:21:58.666754   59961 library.go:117] Package syscall does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.669206   59961 library.go:117] Package encoding does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.709636   59961 library.go:117] Package io/fs does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.709655   59961 library.go:117] Package log does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.717561   59961 library.go:117] Package encoding/base64 does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.717582   59961 library.go:117] Package flag does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:58.725935   59961 library.go:117] Package bufio does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:59.109117   59961 library.go:117] Package text/tabwriter does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
W0403 23:21:59.170760   59961 library.go:101] "math/big" contains non-Go code that can't be inspected for further dependencies:
/opt/homebrew/Cellar/go/1.21.7/libexec/src/math/big/arith_arm64.s
E0403 23:21:59.170785   59961 library.go:117] Package math/big does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:59.367093   59961 library.go:117] Package compress/zlib does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:59.367117   59961 library.go:117] Package crypto/tls does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:59.498535   59961 library.go:117] Package net/url does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
W0403 23:21:59.622872   59961 library.go:101] "crypto/sha256" contains non-Go code that can't be inspected for further dependencies:
/opt/homebrew/Cellar/go/1.21.7/libexec/src/crypto/sha256/sha256block_arm64.s
E0403 23:21:59.622890   59961 library.go:117] Package crypto/sha256 does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:59.622894   59961 library.go:117] Package crypto/x509 does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:59.627454   59961 library.go:117] Package crypto/hmac does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:59.627467   59961 library.go:117] Package os/user does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
W0403 23:21:59.689169   59961 library.go:101] "github.com/dgraph-io/ristretto/z" contains non-Go code that can't be inspected for further dependencies:
/Users/kokoichi/go/pkg/mod/github.com/dgraph-io/ristretto@v0.1.0/z/rtutil.s
W0403 23:21:59.700006   59961 library.go:101] "github.com/cespare/xxhash/v2" contains non-Go code that can't be inspected for further dependencies:
/Users/kokoichi/go/pkg/mod/github.com/cespare/xxhash/v2@v2.2.0/xxhash_arm64.s
E0403 23:21:59.703262   59961 library.go:117] Package math/bits does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:59.740007   59961 library.go:117] Package unicode/utf16 does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:59.767110   59961 library.go:117] Package hash/fnv does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:59.796319   59961 library.go:117] Package go/token does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:21:59.828302   59961 library.go:117] Package compress/gzip does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
W0403 23:21:59.838348   59961 library.go:101] "hash/crc32" contains non-Go code that can't be inspected for further dependencies:
/opt/homebrew/Cellar/go/1.21.7/libexec/src/hash/crc32/crc32_arm64.s
E0403 23:21:59.838366   59961 library.go:117] Package hash/crc32 does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:22:00.137481   59961 library.go:117] Package crypto does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:22:00.137500   59961 library.go:117] Package crypto/ecdsa does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:22:00.137504   59961 library.go:117] Package crypto/ed25519 does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:22:00.137508   59961 library.go:117] Package crypto/elliptic does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:22:00.137512   59961 library.go:117] Package crypto/rsa does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:22:00.137585   59961 library.go:117] Package encoding/asn1 does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:22:00.137591   59961 library.go:117] Package encoding/pem does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
W0403 23:22:00.140510   59961 library.go:101] "crypto/sha512" contains non-Go code that can't be inspected for further dependencies:
/opt/homebrew/Cellar/go/1.21.7/libexec/src/crypto/sha512/sha512block_arm64.s
E0403 23:22:00.140521   59961 library.go:117] Package crypto/sha512 does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
W0403 23:22:00.241995   59961 library.go:101] "runtime/debug" contains non-Go code that can't be inspected for further dependencies:
/opt/homebrew/Cellar/go/1.21.7/libexec/src/runtime/debug/debug.s
E0403 23:22:00.242058   59961 library.go:117] Package runtime/debug does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:22:00.242063   59961 library.go:117] Package runtime/pprof does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:22:00.242067   59961 library.go:117] Package runtime/trace does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:22:01.352102   59961 library.go:117] Package container/list does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:22:01.352138   59961 library.go:117] Package net/textproto does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:22:01.356319   59961 library.go:117] Package net/http/httptrace does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:22:01.504617   59961 library.go:117] Package net/http/httputil does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:22:01.632943   59961 library.go:117] Package embed does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:22:01.678502   59961 library.go:117] Package net/netip does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
E0403 23:22:01.678532   59961 library.go:117] Package net/http/cgi does not have module info. Non go modules projects are no longer supported. For feedback, refer to https://github.com/google/go-licenses/issues/128.
F0403 23:22:01.678558   59961 main.go:77] some errors occurred when loading direct and transitive dependency packages
```


## GOROOT

今回はこれで動いたな。
GitHub だから？

``` sh
$ GOROOT=$(go env GOROOT) go-licenses report github.com/android-project-46group/api-server

```


## indirect を除く

``` sh
# まずは indirect を除いたパッケージ一覧を出す。
awk '!/indirect/ && /v[0-9]+\.[0-9]+\.[0-9]+/ {print $1}' go.mod


GOROOT=$(go env GOROOT) go-licenses report github.com/android-project-46group/api-server > licenses.txt

# あわせる。
# 何個か足りないのがあるが go-licenses の時点で出されてない気がする（エラーにも）。
awk '!/indirect/ && /v[0-9]+\.[0-9]+\.[0-9]+/ {print $1}' go.mod | grep -Ff - licenses.txt > filtered
```

適当なプロジェクトで確認した時に出なかったのは

- https://github.com/golang/mock
- https://github.com/kat-co/vala
- https://github.com/stretchr/testify


## Links

- [/...](https://github.com/google/go-licenses/issues/118#issuecomment-1092407323)
- [GOROOT=](https://github.com/google/go-licenses/issues/193#issuecomment-1462384548)

