package main

import "fmt"

func main() {
	accountAge := 2.6

	// create a new "accountAgeInt" here
	accountAgeInt := int64(accountAge)
	// it should be the result of casting "accountAge" to an integer

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}
