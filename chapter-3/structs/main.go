package main

import "fmt"

type user struct {
	name  string
	email string
	age   int
}

type admin struct {
	user
	level string
}

func newUser(name string) user {
	return user{
		name:  name,
		email: fmt.Sprintf("%s@gmail.com", name),
	}
}

func (u user) Print() {
	fmt.Printf("Name: %s, Email: %s\n", u.name, u.email)
}

func (u *user) ChangeEmail(email string) {
	u.email = email
}

func main() {
	u1 := user{
		name:  "john",
		email: "john@gmail.com",
	}
	fmt.Printf("%+v\n", u1) // {name:john email:john@gmail.com}

	u2 := newUser("jack")
	fmt.Printf("%+v\n", u2) // {name:jack email:jack@gmail.com}

	a := admin{
		user: user{
			name:  "john",
			email: "john@gmail.com",
		},
		level: "super",
	}
	fmt.Printf("%+v\n", a) // {user:{name:john email:john@gmail.com} level:super}

	u3 := newUser("amanda")
	fmt.Printf("%+v\n", u3) // {name:amanda email:amanda@gmail.com}
	fmt.Println(u3.name)    // amanda - we access the field name of the struct user
	u3.name = "ruth"        // we change the value of the field name of the struct user
	fmt.Printf("%+v\n", u3) // {name:ruth email:amanda@gmail.com}
}
