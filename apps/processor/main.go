package main

import (
	"fmt"
)

func main() {
	config := LoadConfig()

	fmt.Printf("Config: %+v\n", config.ProjectID)
}
