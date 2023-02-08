package application

import (
	"fmt"

	"github.com/uladzislaubarysau/gleragames/internal"
)

type Application struct {
}

func NewApplication(c *internal.Config) *Application {
	return &Application{}
}

func (a *Application) Run() {
	fmt.Println("Hello world")
}
