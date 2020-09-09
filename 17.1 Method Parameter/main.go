package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	var s1 = student{"Khaiz Tammam", 21}

	fmt.Println("s1 before", s1.name)

	s1.changeName1("Dika Nafaro")
	fmt.Println("s1 after change name 1", s1.name) //Outuput akan tetap khaiz tamam(hanya berubah di changename1) dikarenakan nilai pada referencenya tidak berubah

	s1.changeName2("Syila Andini")
	fmt.Println("s1 after change name 2", s1.name)
}

func (s student) changeName1(name string) {
	fmt.Println("--> on Change Name 1, name changed to", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("--> on Change Name 1, name changed to", name)
	s.name = name
}
