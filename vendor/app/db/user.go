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

func Login(mobile, pwd string) (*User, int64) {
	stmt, err := mydb.Prepare("SELECT user.uid, user.pwd, user.profession, user.business, user.corp FROM user WHERE user.mobile=?")
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
	return user, uid
}

func Register(mobile, pwd string) int64 {
	stmt1, err := mydb.Prepare("SELECT user.uid FROM user WHERE user.mobile=?")
	defer stmt1.Close()
	util.CheckError(err)

	rows, err := stmt1.Query(mobile)
	defer rows.Close()
	util.CheckError(err)

	if rows.Next() {
		return -1
	}

	stmt2, err := mydb.Prepare("INSERT INTO user(user.mobile, user.pwd) VALUES (?,?)")
	defer stmt2.Close()
	util.CheckError(err)

	res, err := stmt2.Exec(mobile, pwd)
	util.CheckError(err)

	lastId, err := res.LastInsertId() //uid要用int64
	util.CheckError(err)
	log.Println("lastId=", lastId)
	return lastId
}

func CompleteInfo(uid int64, profession, business, corp string) int64 {
	stmt1, err := mydb.Prepare("SELECT user.uid FROM user WHERE user.uid=?")
	defer stmt1.Close()
	util.CheckError(err)

	rows, err := stmt1.Query(uid)
	defer rows.Close()
	util.CheckError(err)

	if !rows.Next() {
		return -1
	}

	stmt2, err := mydb.Prepare("UPDATE user SET user.profession=?, user.business=?, user.corp=? WHERE user.uid=?")
	defer stmt2.Close()
	util.CheckError(err)

	res, err := stmt2.Exec(profession, business, corp, uid)
	util.CheckError(err)

	lastId, err := res.LastInsertId() //uid要用int64
	util.CheckError(err)
	log.Println("lastId=", lastId)
	return lastId
}



