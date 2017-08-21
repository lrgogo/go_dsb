package db

import (
	"log"
	"app/util"
)

type User struct {
	Uid        int64
	Mobile     string
	Profession string
	Corp       string
	Business   string
}

func login(mobile, pwd string) (*User, int64) {
	stmt, err := db.Prepare("SELECT user.uid, user.pwd, user.profession, user.business, user.corp FROM user WHERE user.mobile=?")
	defer stmt.Close()
	util.CheckError(err)

	rows, err := stmt.Query(mobile)
	defer rows.Close()
	util.CheckError(err)

	if !rows.Next() {
		return nil, -1
	}
	var uid int64
	var password, profession, business, corp string
	err = rows.Scan(&uid, &password, &profession, &business, &corp)
	util.CheckError(err)

	if password != pwd {
		return nil, -2
	}
	user := &User{
		Uid:        uid,
		Mobile:     mobile,
		Profession: profession,
		Corp:       corp,
		Business:   business,
	}
	return user, 0
}

func register(mobile, pwd string) int64 {
	stmt, err := db.Prepare("INSERT INTO user(user.mobile, user.pwd) VALUES (?,?)")
	defer stmt.Close()
	util.CheckError(err)

	res, err := stmt.Exec(mobile, pwd)
	util.CheckError(err)

	lastId, err := res.LastInsertId() //uid要用int64
	util.CheckError(err)
	log.Println("lastId=", lastId)
	return lastId
}

func completeInfo(uid int64, profession, business, corp string) int64 {
	stmt, err := db.Prepare("SELECT user.uid FROM user WHERE user.uid=?")
	defer stmt.Close()
	util.CheckError(err)

	rows, err := stmt.Query(uid)
	defer rows.Close()
	util.CheckError(err)

	if !rows.Next() {
		return -1
	}

	stmt, err = db.Prepare("INSERT INTO user(user.profession, user.business, user.corp) VALUES (?,?,?)")
	defer stmt.Close()
	util.CheckError(err)

	res, err := stmt.Exec(profession, business, corp)
	util.CheckError(err)

	lastId, err := res.LastInsertId() //uid要用int64
	util.CheckError(err)
	log.Println("lastId=", lastId)
	return lastId
}



