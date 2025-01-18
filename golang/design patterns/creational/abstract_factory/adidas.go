package abstractfactory

type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{}
}
