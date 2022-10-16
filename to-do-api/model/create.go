package model

import "fmt"

func CreateTodo(name, todo string) error {
	_, err := con.Exec("INSERT INTO TODO VALUES(?, ?)", name, todo)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DeleteTodo(name string) error {
	_, err := con.Exec("DELETE FROM TODO WHERE name=?", name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}