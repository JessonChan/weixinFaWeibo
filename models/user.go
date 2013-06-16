package models

import (
	"database/sql"
	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id             int `PK`
	WeixinUserName string
	Token          string
}
type AccessToken struct {
	Access_token string
	Expires_in   int
	Uid          string
}

func InsertUser(u *User) error {
	u1:=GetUser(u.WeixinUserName)
    if u!=nil{
		u.Id=u1.Id
	}

	db := GetDB()
	return db.Save(u)
}

func GetUser(name string) (u User) {
	db := GetDB()
	db.Where("weixin_user_name='" + name + "'").Find(&u)
	return
}
func GetToken(name string) string {
	u := GetUser(name)
	return u.Token
}

func GetDB() beedb.Model {
	db, e := sql.Open("mysql", "root:123@/golang?charset=utf8")
	if e != nil {
		println(e.Error())
	}
	orm := beedb.New(db)
	orm.SetTable("weixin")
	return orm
}
