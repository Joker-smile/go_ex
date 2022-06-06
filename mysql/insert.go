package mysql

import "fmt"

func Insert() {
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu002", "man", "stu02@qq.com")
	if err != nil {
		fmt.Println("exec failed1, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed2, ", err)
		return
	}

	fmt.Println("insert succ:", id)

	defer Db.Close()
}
