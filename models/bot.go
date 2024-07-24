package models

type Bot struct {
	Id int    `gorm:"primaryKey" json:"id"`
	Ip string `json:"ip"`
}
