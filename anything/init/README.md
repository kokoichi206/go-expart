## 1 package

only `main.go`

``` go
package main

import "fmt"

// global variable
var a = func() string {
	fmt.Println("main global variable")
	return "a"
}()

// init function
func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("main")
}
```

output

``` sh
$ go run main.go
main global variable
main init
main
```

## package dependencies

`main.go` and `repository` package

``` go
package main

import (
	"fmt"
	_ "initialize-order/repository"
	_ "initialize-order/usecase"
	_ "initialize-order/a"
)

// global variable
var a = func() string {
	fmt.Println("main global variable")
	return "a"
}()

// init function
func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("main")
}
```

output

``` sh
$ go run main.go
a global variable
a init
repository global variable
repository init
usecase global variable
usecase init
main global variable
main init
main
```

In terms of importing packages, they are initialized in ascending alphabetical order of the **package names**.
(Regardless of the importing order, the results are consistent.)

``` go
import (
	"fmt"
	_ "initialize-order/usecase"
	_ "initialize-order/repository"
	_ "initialize-order/a"
)
```
