package mongo

import (
	"time"

	"gopkg.in/mgo.v2"
)

type Info struct {
	Hostname string
	Database string
	Username string
	Password string
}

func (i *Info) Connect() (*mgo.Session, error) {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:     []string{i.Hostname},
		Timeout:   60 * time.Second,
		Database:  i.Database,
		Username:  i.Username,
		Password:  i.Password,
		Source:    "admin",
		Mechanism: "SCRAM-SHA-1",
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)

	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	return session, err
}
