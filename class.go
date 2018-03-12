package main

import "fmt"

type Person struct {
	name string
	sex	bool
	age int
    address string
}

func (person *Person) setName(_name string) {
    person.name = _name
}

func (person *Person) getName() string {
    return person.name
}


func main(){
	
	person := &Person{"Test",true,30,"Shenzhen,China"}

	fmt.Println(person.name);

	person.name = "Neo"
	fmt.Println(person.name);

	person.setName("Netkiller")
	fmt.Println(person.getName());

}