package main

import "fmt"

type IpAddr [4]byte

func (ip IpAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IpAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for key, value := range hosts {
		fmt.Printf("%v: %v\n", key, value)
	}
}
