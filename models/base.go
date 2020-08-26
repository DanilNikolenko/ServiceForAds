package models

import "time"

type Base struct {
	CreatedAt time.Time `orm:"column(created_at);type(timestamp without time zone);auto_now_add" json:"-"`
}
