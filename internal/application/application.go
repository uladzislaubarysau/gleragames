package application

import (
	"log"
	"net/http"

	"github.com/uladzislaubarysau/gleragames/internal"
	"github.com/uladzislaubarysau/gleragames/internal/db"
)

type Application struct {
	c  *internal.Config
	db *db.DB
}

func NewApplication(c *internal.Config) *Application {
	return &Application{
		c:  c,
		db: db.NewDB(c.DataSourceName, c.DBQueryTimeout.Duration),
	}
}

func (a *Application) SendRequest() {
	resp, err := http.Get(a.c.Url)
	if err != nil {
		log.Fatalf("can't make GET request: %v", err)
	}

	a.db.SaveRequest(resp.StatusCode)
}
