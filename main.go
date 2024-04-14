package main

import (
	"fmt"

	os "github.com/maitaken/go_deadcode_sample/internal/entities"
)

func main() {
	device := os.Windows{}
	fmt.Println(device.Type())
}

// unused
func deadcode() {
	device := os.Android{}
	fmt.Println(device.Type())
}
