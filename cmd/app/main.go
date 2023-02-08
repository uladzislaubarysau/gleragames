package main

import (
	"fmt"

	"github.com/uladzislaubarysau/gleragames/internal"
	"github.com/uladzislaubarysau/gleragames/internal/application"
)

func main() {
	c, err := internal.NewConfig()
	if err != nil {
		fmt.Printf("error while parsing config: %s\n", err)
		return
	}
	a := application.NewApplication(c)
	a.SendRequest()
}
