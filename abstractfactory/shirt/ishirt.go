package shirt

// Implement Shirt interface
type IShirt interface {
	SetSize(int)
	GetSize() int
	SetPrice(float32)
	GetPrice() float32
}

// Abstract Shirt Structure
type Shirt struct {
	Size int
	Price float32
}

func (s *Shirt) SetSize(size int) {
	s.Size = size
}

func (s *Shirt) GetSize() int {
	return s.Size
}

func (s *Shirt) SetPrice(price float32) {
	s.Price = price
}

func (s *Shirt) GetPrice() float32 {
	return s.Price
}