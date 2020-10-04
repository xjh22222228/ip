


## Install
```bash
$ go get github.com/xjh22222228/ip
```

## Usage
```go
package main
import (

"fmt"
"github.com/xjh22222228/ip"
)

func main()  {
    ipv4, err := ip.V4()
    ipv6, err2 := ip.V6()

    if err != nil {
        fmt.Println(err)
    }
    if err2 != nil {
        fmt.Println(err)
    }

    
    fmt.Println(ipv4)
    fmt.Println(ipv6)
}
```
