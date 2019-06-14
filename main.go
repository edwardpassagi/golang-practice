package main

import "fmt"
import "errors"
type User struct{
	Name, Role, Email string
	Age int
}

func (u User) Salary() (int, error) {
	if u.Role == ""{
		return 0, errors.New("No Role Value")
	}
	switch u.Role {
	case "Dev": 
		return 100, nil
	case "Arc": 
		return 200, nil
	} 
	return 0, errors.New(
		fmt.Sprintf("Not able to handle the '%s' role", u.Role),
	)
}

func main() {
	user := User{Name: "Edward", Role: "Dev"}
	salary, err := user.Salary()
	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println("Salary", salary)
	}
}