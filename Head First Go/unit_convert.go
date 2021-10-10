package main

import "fmt"

type Liters float64

type Milliliters float64

type Gallons float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

func main() {
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())

	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())

	fmt.Println(Liters(12.09248342))
	fmt.Println(Milliliters(12.09248342))
	fmt.Println(Gallons(12.09248342))
}
