package application

import (
	"fmt"
	"log"
	"net/http"

	"github.com/uladzislaubarysau/gleragames/internal"
	"github.com/uladzislaubarysau/gleragames/internal/repository"
)

type Application struct {
	c  *internal.Config
	db *repository.DBRepository
}

func NewApplication(c *internal.Config) *Application {
	return &Application{
		c:  c,
		db: repository.NewDBRepository(c.DataSourceName, c.DBQueryTimeout.Duration),
	}
}

func (a *Application) SendRequest() {
	resp, err := http.Get(a.c.Url)
	if err != nil {
		log.Fatalf("can't make GET request: %v", err)
	}

	fmt.Printf("status code: %d\n", resp.StatusCode)
	a.db.SaveRequest(resp.StatusCode)
}
