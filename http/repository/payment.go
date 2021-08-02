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

	who := bson.M{"code" : paymentCode}

	what := bson.M{"status": status}
	err = table.Update(who, what)
	if err != nil {

		return err
	}
	return nil

}