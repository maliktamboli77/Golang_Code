package main

import (
	"fmt"
	"strings"
)

func joinstr(elements ...string) string {
	return strings.Join(elements, "-")
}

func main() {
	fmt.Println(joinstr())

	fmt.Println(joinstr("Malik", "Tamboli"))
	fmt.Println(joinstr("H", "E", "L", "L", "O"))
}
