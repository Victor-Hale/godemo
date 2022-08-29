package models

import (
	"fmt"
	"time"
)
type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {

	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

type Information struct {
	Id         int       `json:"id" orm:"column(id)"`
	Photo    string    `json:"photo" orm:"column(photo)"`
	Anthor   string    `json:"anthor" orm:"column(anthor)"`
	Port   string    `json:"port" orm:"column(port)"`
	Title   string    `json:"title" orm:"column(title)"`
	Teleurl   string    `json:"teleurl" orm:"column(teleurl)"`
	Create_at LocalTime `gorm:"column:create_at;default:null" json:"create_at"`
	Update_at LocalTime `gorm:"column:update_at;default:null" json:"update_at"`
}

