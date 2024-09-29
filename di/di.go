package di

import (
	_ "github.com/lib/pq" // <------------ here
	"go.uber.org/dig"
)

func NewDiContainer() (*dig.Container, error) {
	c := dig.New()
	providers := []interface {
	}{}
	for _, provider := range providers {
		if err := c.Provide(provider); err != nil {
			return nil, err
		}
	}
	return c, nil
}
