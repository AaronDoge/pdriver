package main

import "fmt"

type Citizen struct {
	Name     string
	Age      int
	Race     string
	Gender   string
	Country  string
	Language string
}

type teacher struct {
	Citizen
	Subject string
	School  string
}

func (c *Citizen) Score() int {
	return len(c.Name) + c.Age + len(c.Language)
}

func (t *teacher) Score() int {
	return t.Citizen.Score() + len(t.Subject) + len(t.School)
}

func main() {
	// 初始化一个实例
	var Jhon = teacher{
		Citizen: Citizen{
			Name:     "Jhon",
			Age:      30,
			Race:     "White",
			Gender:   "Male",
			Country:  "US",
			Language: "US-EN",
		},
		Subject: "English Writing",
		School:  "M High School",
	}

	fmt.Println(Jhon.Citizen.Language)
	fmt.Println("Score is: ", Jhon.Citizen.Score())
	fmt.Println("Teacher Score is: ", Jhon.Score())


}
