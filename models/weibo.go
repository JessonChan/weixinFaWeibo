package models

type WeiboUser struct {
	Id    int64
	Idstr string
}
type WeiboResponse struct {
	Id   int64
	User WeiboUser
}
