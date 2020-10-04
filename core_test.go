package ip

import (
	"fmt"
	"testing"
)

func TestIsIPV4(t *testing.T) {
	ok := IsIPv4("192.168.0.1")

	if ok == false {
		t.Fatal("=====> TestIsIpv4")
	}
}

func TestIsIpv6(t *testing.T) {
	ok := IsIPv6("fe80::200:f8ff:fe21:67cf")

	if ok == false {
		t.Fatal("=====> TestIsIpv6")
	}
}

func TestV4(t *testing.T) {
	ipv4, err := V4()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("IP: %v", ipv4)
}

func TestV6(t *testing.T) {
	ipv6, err := V6()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("IP: %v", ipv6)
}

