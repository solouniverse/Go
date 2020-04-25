package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	fmt.Println(os.Environ())

	fmt.Println()

	// blank identifier takes key or index, env takes value
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	name := os.Getenv("USER")

	fmt.Println("Namaskaram", name)

	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	// Current User
	fmt.Println("Namaskaram " + user.Name + "(id: )" + user.Uid + ")")
	fmt.Println("Username: " + user.Username)
	fmt.Println("Dir: " + user.HomeDir)
	fmt.Println("Real User: " + os.Getenv("SUDO_USER"))
}
