package entity

import "strings"

type Currency struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Currency) Valid() bool {
	c.Name = strings.TrimSpace(c.Name)
	return c.ID > 0 && len(c.Name) > 0
}
