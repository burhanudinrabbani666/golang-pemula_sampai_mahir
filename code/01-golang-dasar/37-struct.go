package main

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	// 1.
	var bani Customer

	bani.Name = "Burhanudin"
	bani.Address = "Cirebon"
	bani.Age = 23

	// 2.
	nico := Customer{
		Name:    "Nico",
		Address: "Cirebon",
		Age:     24,
	}

	// 3.
	anang := Customer{"Anang", "Bekasi", 25}

}
