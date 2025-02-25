package main

import (
	"flag"
	"fmt"
)

type Celsius float32

type celsiusFlag struct {
	Celsius
}

func c2f(c float32) float32 {
	return c*9/5 + 32
}

func f2c(f float32) float32 {
	return (f - 32) * 5 / 9
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var val float32
	fmt.Sscanf(s, "%f%s", &val, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(val)
		return nil
	case "F", "°F":
		f.Celsius = Celsius(f2c(val))
		return nil
	}
	return fmt.Errorf("invalid temprature %q", s)
}

func (f *celsiusFlag) String() string {
	return fmt.Sprintf("%f°C", f.Celsius)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{Celsius: value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temprature")

// func main() {

// 	flag.Parse()
// 	fmt.Println(*temp)

// }
