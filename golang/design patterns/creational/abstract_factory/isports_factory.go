package abstractfactory

import "fmt"

type Brand int

const (
	adidas Brand = iota
	nike
)

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func (b *Brand) String() string {
	return [...]string{"Adidas", "Nike"}[*b]
}

func GetSportsFactory(brand Brand) (ISportsFactory, error) {
	switch brand {
	case adidas:
		return &Adidas, nil
	case nike:
		return &nike, nil
	default:
		return nil, fmt.Errorf("Brand doesnt exist")
	}
}
