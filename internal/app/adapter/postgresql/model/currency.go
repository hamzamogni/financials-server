package model

type Currency struct {
	ID     uint   `gorm:"primarykey"`
	Name   string `gorm:"uniqueIndex"`
	Symbol string
}
