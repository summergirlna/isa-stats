package model

import "time"

// Inbound 入国者
type Inbound struct {
	Airport string
	Person  *Person
	Date    time.Time
}
