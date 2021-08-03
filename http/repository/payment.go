package repository

import (
	"evermos-test/database/entity"
	"evermos-test/http/interfaces"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repositoryPayments struct {
	dbSession *mgo.Session
	database  string
}

var collectionPayment = "payments"

func NewPaymentsRepository(sess *mgo.Session, database string) interfaces.PaymentInterface {
	return &repositoryPayments{sess, database}

}

func (repo *repositoryPayments) Create(e *entity.Payment) (bool, error) {
	var err error
	ds := repo.dbSession.Copy()
	defer ds.Close()

	table := ds.DB(repo.database).C(collectionPayment)

	index := mgo.Index{
		Key:    []string{"code"},
		Unique: true,
	}
	table.EnsureIndex(index)
	e.ValidateBeforeCreate()
	if err = table.Insert(&e); err != nil {
		return false, err
	}

	return true, err
}

func (repo *repositoryPayments) UpdateStatus(paymentCode, status string)  error{
	var err error

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionPayment)

	var payment *entity.Payment
	table.Find(bson.M{"code" : paymentCode}).One(&payment)

	who := bson.M{"code" : paymentCode}

	what := bson.M{
		"status": status,
		"code" : paymentCode,
		"created_at_utc" : payment.CreatedAtUTC,
		"created_at_timezone" : payment.CreatedAtTimezone,
	}
	err = table.Update(who, what)
	if err != nil {

		return err
	}
	return nil

}