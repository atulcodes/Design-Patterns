package main

type Musket struct {
	Gun
}

func newMusket() *Musket {
	return &Musket{
		Gun: Gun{
			name: "Musket",
			power: 99,
		},
	}
}