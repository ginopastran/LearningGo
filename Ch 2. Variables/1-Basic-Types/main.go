package main

import "fmt"

	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

func main() {
	// initialize variables here

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
