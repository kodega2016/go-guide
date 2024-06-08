package main

import "fmt"

type User struct {
	name    string
	age     int
	address Address
}

type Address struct {
	city    string
	street  string
	zipCode int
}

func main() {
	kode := User{
		name: "Khadga Bahadur Shrestha",
		age:  27,
		address: Address{
			city:    "Madhumalla",
			street:  "Miklajung,Morang",
			zipCode: 6500,
		},
	}
	kode.print()

	rakas := User{
		"Sakar Subedi",
		29,
		Address{
			"Biratnagar",
			"Tintolia Road",
			56000,
		},
	}
	rakas.print()
}
func (u User) print() {
	fmt.Printf("My name is %s and i am %d years old and i live in %s\n", u.name, u.age, u.address.street)
}
