package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("<b>User name is: %s. He is %d and he has money"+
		" equal: %d</b>", u.Name, u.Age, u.Money)
}
func (u *User) setNewName(newName string) {
	u.Name = newName
}
func HomePage(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8}
	//bob.setNewName("Jonatan")
	//fmt.Fprintf(w, bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}
func AboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About")
}
func handleRequest() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about/", AboutPage)
	http.ListenAndServe(":8089", nil)
}
func main() {
	handleRequest()
}
