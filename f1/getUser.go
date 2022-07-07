package f1

import (
	"database/sql"
	"github.com/pkg/errors"
)

type User struct {
	Id   string
	Name string
}

func getUser(db *sql.DB) ([]User, error) {
	var u User
	var us []User
	// 创建sql
	sql := "select id,name from t1 where id = xxx"

	// 拿到结果
	rows, err := db.Query(sql)

	// 查询出错直接给上级
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&u.Id, &u.Name)
		us = append(us, u)
	}

	return us, errors.Wrap(rows.Err(), "no rows return")
}
