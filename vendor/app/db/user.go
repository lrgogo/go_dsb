package db

type User struct {
	Uid        int `db:"uid" json:"uid"`
	Mobile     string `db:"mobile" json:"mobile"`
	Pwd        string `db:"pwd" json:"-"`
	Profession string `db:"profession" json:"profession"`
	Corp       string `db:"corp" json:"corp"`
	Business   string `db:"business" json:"business"`
}
