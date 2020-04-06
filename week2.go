package main

import (
	"fmt"
	"math"
)

type (
	Feet         float64
	Centimeter   float64
	Seconds      float64
	Minutes      float64
	Milliseconds float64
	Celsius      float64
	Farenheit    float64
	Radian       float64
	Degree       float64
	Kilogram     float64
	Pounds       float64
)

type Converter struct{}

func (cvr Converter) CentimeterToFeet(c Centimeter) Feet {
	var result = Feet(c / 30.48)
	return result
}
func (cvr Converter) FeetToCentimeter(f Feet) Centimeter {
	var result = Centimeter(f * 30.48)
	return result
}
func (cvr Converter) MinutesToSeconds(m Minutes) Seconds {
	var result = Seconds(m * 60)
	return result
}
func (cvr Converter) SecondsToMinutes(s Seconds) Minutes {
	var result = Minutes(s / 60)
	return result
}
func (cvr Converter) CelsiusToFarenheit(ce Celsius) Farenheit {
	var result = Farenheit(ce*(9/5) + 32)
	return result
}
func (cvr Converter) FarenheitToCelsius(f Farenheit) Celsius {
	var result = Celsius(f * 3230.40)
	return result
}
func (cvr Converter) RadianToDegree(r Radian) Degree {
	var result = Degree(r * (180 / math.Pi))
	return result
}
func (cvr Converter) DegreeToRadian(d Degree) Radian {
	var result = Radian(d * 0.01745)
	return result
}
func (cvr Converter) KilogramToPounds(k Kilogram) Pounds {
	var result = Pounds(k * 2.205)
	return result
}
func (cvr Converter) PoundsToKilogram(p Pounds) Kilogram {
	var result = Kilogram(p * 0.454)
	return result
}
func (cvr Converter) SecondsToMilliseconds(s Seconds) Milliseconds {
	var result = Milliseconds(s * 1000)
	return result
}
func (cvr Converter) MillisecondsToSeconds(m Milliseconds) Seconds {
	var result = Seconds(m / 1000)
	return result
}
func main() {
	cvr := Converter{}

	fmt.Println(cvr.CentimeterToFeet(32))
	fmt.Println(cvr.FeetToCentimeter(2))
	fmt.Println(cvr.MinutesToSeconds(7))
	fmt.Println(cvr.SecondsToMinutes(8))
	fmt.Println(cvr.CelsiusToFarenheit(4))
	fmt.Println(cvr.FarenheitToCelsius(25.3))
	fmt.Println(cvr.RadianToDegree(108.3))
	fmt.Println(cvr.DegreeToRadian(77.25))
	fmt.Println(cvr.KilogramToPounds(21))
	fmt.Println(cvr.PoundsToKilogram(100))
	fmt.Println(cvr.SecondsToMilliseconds(202))
	fmt.Println(cvr.MillisecondsToSeconds(99999999))
}
