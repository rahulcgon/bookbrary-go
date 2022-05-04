package dbconfig

import "fmt"

func GetAllUsers() {

	// Here a SQL query is used to return all
	// the data from the table user.
	result, err := DB.Query("SELECT * FROM users limit 10")

	// handle error
	if err != nil {
		panic(err)
	}

	// the result object has a method called Next,
	// which is used to iterate through all returned rows.
	for result.Next() {

		var uid int
		var name string

		// The result object provided Scan  method
		// to read row data, Scan returns error,
		// if any. Here we read id and name returned.
		err = result.Scan(&uid, &name)

		// handle error
		if err != nil {
			panic(err)
		}

		fmt.Printf("uid: %d Name: %s\n", uid, name)
	}
}