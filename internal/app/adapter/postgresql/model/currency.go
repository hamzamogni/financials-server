package model

type Currency struct {
	ID     uint   `gorm:"primarykey;autoIncrement"`
	Name   string `gorm:"uniqueIndex"`
	Symbol string
}
