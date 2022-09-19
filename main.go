package main

import (
	"fmt"
	"pinosint/pkg/iplogging"
	_ "pinosint/pkg/iplogging"
)

func main() {
	Menu()
}

func Menu() {
	fmt.Println("PinOsint \n \n")

	for {
		var selected string
		fmt.Println("Choose from Menu ")
		fmt.Println("1: IP Logging ")
		fmt.Println("2: In Development ")
		fmt.Println("0 : End Program ")
		fmt.Scanln(&selected)
		if selected == "1" {
			iplogging.Iplog()
		}
		if selected == "2" || selected == "0" {
			break
		}
	}
}
