package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  int
}

func (p *Person) String() string {
	return fmt.Sprintf("name:%s, age:%d, sex:%d", p.Name, p.Age, p.Sex)
}

func (p *Person) Format(f fmt.State, c rune) {
	if c == 'L' {
		f.Write([]byte(p.String()))
		f.Write([]byte("Person has three fields."))
	} else {
		f.Write([]byte(fmt.Sprintln(p.String())))
	}
}

func main() {
	p := &Person{Name: "张三", Age: 19, Sex: 1}
	fmt.Printf("%L", p)
}
