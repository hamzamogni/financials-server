package model

type Currency struct {
	Symbol string `gorm:"primarykey"`
	Name   string
}
