package main

type Person struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func main() {
	p1 := new(Person)

	p1.Name = "Zhanserik"
	p1.Age = 19
}
