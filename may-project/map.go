package main

import (
	"fmt"
)

func main() {
	// write  progrom using map to desplay the code of a country and their full
	country := map[string]string{
		"NG": "NIGERIA",
		"CA": "CANANADA",
		"UK": "UNITED KINGDOM",
		"US": "UNITED STATE",
		"GA": "GHANA",
	}

	country["EU"] = "ERUOPE"
	fmt.Println(country)
	delete(country, "GA")
	fmt.Println(country)

	value, exit := country["NG"]
	fmt.Println(value, exit)

}
