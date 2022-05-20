package main

import "fmt"

type contactInfo struct {
  email string
  zipCode int
}
type person struct {
  firstName string
  lastName string
  contactInfo
}

func (p person) print() {
  fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
  (*pointerToPerson).firstName = newFirstName
}

func main() {
  // Not recommended syntax
  //alex := person{"Alex", "Anderson"}
  //alex := person{firstName: "Alex", lastName: "Anderson"}
  //fmt.Println(alex)

  /*var alex person
  alex.firstName = "Alex"
  alex.lastName = "Anderson"
  fmt.Println(alex)
  fmt.Printf("%+v", alex)*/

  jim := person{
    firstName: "Jim",
    lastName: "Party",
    contactInfo: contactInfo{
      email: "jim@gmail.com",
      zipCode: 9400,
    },
  }

  //jimPointer := &jim
  //jimPointer.updateName("Jimmy")

  jim.updateName("Jimmy")

  jim.print()
  //fmt.Println(jim)
  //fmt.Printf("%+v", jim)
}