package shoe

// Implement Shoe interface
type IShoe interface {
	SetName(string)
	GetName() string
	SetPrice(float32)
	GetPrice() float32
}

// Abstract Shoe structure
type Shoe struct {
	Name string
	Price float32
}

func (s *Shoe) SetName(name string) {
	s.Name = name
}

func (s *Shoe) GetName() string {
	return s.Name
}

func (s *Shoe) SetPrice(price float32) {
	s.Price = price
}

func (s *Shoe) GetPrice() float32 {
	return s.Price
}