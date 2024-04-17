package application

import "errors"

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	Id     string
	Name   string
	Price  float64
	Status string
}

//A declaração dos métodos da interface ProductInterface
//Isso significa que a struct Product implementa a interface ProductInterface
// func (p *Product) IsValid() (bool, error) {

// }

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil //retorna nil para indicar que não houve erro
	}

	return errors.New("the price must be greater than zero to enable the product")
}

// func (p *Product) Disable() error {

// }

func (p *Product) GetId() string {
	return p.Id
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
