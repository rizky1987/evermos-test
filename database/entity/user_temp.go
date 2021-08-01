package entity

import "gopkg.in/mgo.v2/bson"

type UserGroupByRT struct {
	RT			int			`bson:"_id"`
	Data		[]*User		`bson:"data"`
}

// untuk kasus ini kita memakai 1 entity untuk 3 group karena type datanya sama
type UserGroupByOptionalGroup struct {
	Id			bson.ObjectId			`bson:"_id"`
	Data		[]*User					`bson:"data"`
}