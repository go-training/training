package main

import (
	"fmt"
	"strings"
)

type searchOpts struct {
	username string
	email    string
}

func getUserListSQL(username, email string) string {
	sql := "select * from user"
	where := []string{}

	if username != "" {
		where = append(where, fmt.Sprintf("username = '%s'", username))
	}

	if email != "" {
		where = append(where, fmt.Sprintf("email = '%s'", email))
	}

	return sql + " where " + strings.Join(where, " or ")
}

func getUserListOptsSQL(opts searchOpts) string {
	sql := "select * from user"
	where := []string{}

	if opts.username != "" {
		where = append(where, fmt.Sprintf("username = '%s'", opts.username))
	}

	if opts.email != "" {
		where = append(where, fmt.Sprintf("email = '%s'", opts.email))
	}

	return sql + " where " + strings.Join(where, " or ")
}

func main() {
	fmt.Println(getUserListSQL("appleboy", ""))
	fmt.Println(getUserListSQL("appleboy", "appleboy.tw@gmail.com"))

	fmt.Println(getUserListOptsSQL(searchOpts{
		username: "appleboy",
	}))

	fmt.Println(getUserListOptsSQL(searchOpts{
		username: "appleboy",
		email:    "test@gmail.com",
	}))
}
