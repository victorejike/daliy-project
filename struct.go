package main

import (
	"fmt"
)

type c_pannel struct {
	name            string
	phone_number    string
	addres_of_user1 string
	location        string
}

func main() {
	var user c_pannel

	user.name = "victor ejike"
	user.phone_number = "08062463468"
	user.addres_of_user1 = "life"
	user.location = "abuja"

	fmt.Println("user name:  ", user.name)
	fmt.Println("phone number: ", user.phone_number)
	fmt.Println("address of user: ", user.addres_of_user1)
	fmt.Println("location: ", user.location)

}
