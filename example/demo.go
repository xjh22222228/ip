package main

import (
	"fmt"
	"github.com/xjh22222228/ip"
)

func main()  {
	v4, err := ip.V4()

	if err != nil {
		fmt.Println("ipv4 err: ", err)
	}

	fmt.Println("v4: ", v4)

	v6, err := ip.V6()

	if err != nil {
		fmt.Println("ipv6 err: ", err)
	}

	fmt.Println("v6: ", v6)
}
