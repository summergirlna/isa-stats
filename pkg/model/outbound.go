package model

import "time"

type Outbound struct {
	Airport string
	Person  *Person
	Date    time.Time
}
