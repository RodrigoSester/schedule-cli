package main

import (
	"fmt"

	"schedule.com/schedule"
)

func main() {
	message := schedule.Hello("Rodrigo")
	fmt.Println(message)
}
