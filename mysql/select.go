package mysql

import "fmt"

func Select() {
	var person []Person
	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 2)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("select succ:", person)
	defer Db.Close()
}
