package domain

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Currency struct {
	Symbol string
	Name   string
}

func NewCurrency(symbol string, name string) *Currency {
	return &Currency{
		Symbol: cases.Upper(language.English).String(symbol),
		Name:   cases.Title(language.Und).String(name),
	}
}
