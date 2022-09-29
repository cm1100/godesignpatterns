package main

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

// factory fucntion
func NewPerson(name string, age int) Person {
	if age < 16 {
		panic("Validation")
	}
	return Person{name, age, 2}

}

// Sometimes You want to have something happen when the object is created

func main() {
	p := NewPerson("John", 22)
	p.EyeCount = 3
}
