package main

import (
	"fmt"
)

func main() {
	var company = map[string]string{"name": "victor", "job": "software", "location": "abuja"}
	b := map[string]int{"year": 2026, "month": 05, "day": 28}
	fmt.Printf("company\t%s\n", company)
	fmt.Printf("b\t%v\n", b)
}
