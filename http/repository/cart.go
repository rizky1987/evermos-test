package repository

import (
	"evermos-test/database/entity"
	"evermos-test/helper"
	"evermos-test/http/interfaces"
	"evermos-test/http/request"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repositoryCarts struct {
	dbSession *mgo.Session
	database  string
}

var collectionCart = "carts"

func NewCartsRepository(sess *mgo.Session, database string) interfaces.CartInterface {
	return &repositoryCarts{sess, database}

}

func (repo *repositoryCarts) Create(e *entity.Cart) (bool, error) {
	var err error
	ds := repo.dbSession.Copy()
	defer ds.Close()

	table := ds.DB(repo.database).C(collectionCart)

	if err = table.Insert(&e); err != nil {
		return false, err
	}

	return true, err
}

func (repo *repositoryCarts) Update(id *bson.ObjectId, e *entity.Cart) (bool, error) {
	var err error

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionCart)
	err = table.Update(
		bson.M{"_id": id},
		bson.M{"$set": &e},
	)

	if err != nil {
		return false, err
	}

	return true, err
}

func (repo *repositoryCarts) FindByCartName(name string) (*entity.Cart, error) {
	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionCart)

	var result entity.Cart
	err := table.Find(bson.M{"name": name}).One(&result)

	return &result, err
}

func (repo *repositoryCarts) FindById(id *bson.ObjectId) (*entity.Cart, error) {

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionCart)

	var result entity.Cart
	err := table.Find(bson.M{"_id": id}).One(&result)

	return &result, err
}

func (repo *repositoryCarts) FindAll(searchParam request.SearchParamWithPagingCartRequest) ([]*entity.Cart, error, int) {

	result := []*entity.Cart{}

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionCart)

	pipeline := []bson.M{}
	pipelineCount := []bson.M{}

	paymentCode := helper.TrimWhiteSpace(searchParam.PaymentCode)
	if paymentCode != "" {
		pipeline = append(pipeline, bson.M{"$match": bson.M{"payment_code": paymentCode}})
		pipelineCount = append(pipelineCount, bson.M{"$match": bson.M{"payment_code": paymentCode}})
	}

	skip := 0
	limit := 0

	if searchParam.Limit > 0 && searchParam.Page > 0 {
		if searchParam.Limit > 0 {
			limit = searchParam.Limit
		}

		if searchParam.Page > 1 {
			skip = limit * (searchParam.Page - 1)
		}

		if skip > 0 {
			skipQuery := bson.M{"$skip": skip}
			pipeline = append(pipeline, skipQuery)
		}

		if limit > 0 {
			limitQuery := bson.M{"$limit": limit}
			pipeline = append(pipeline, limitQuery)
		}
	}
	if err := table.Pipe(pipeline).All(&result); err != nil {
		return nil, err, 0
	}

	resultCount := &TempCount{}
	countQuery := bson.M{"$count": "count"}
	pipelineCount = append(pipelineCount, countQuery)
	if err := table.Pipe(pipelineCount).One(resultCount); err != nil {
		return nil, err, 0
	}

	return result, nil, resultCount.Count
}

func (repo *repositoryCarts) FindAllWithPaging(searchParam request.SearchParamWithPagingCartRequest) ([]*entity.Cart, error) {

	result := []*entity.Cart{}

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionCart)

	pipeline := []bson.M{}

	paymentCode := helper.TrimWhiteSpace(searchParam.PaymentCode)
	if paymentCode != "" {
		pipeline = append(pipeline, bson.M{"$match": bson.M{"payment_code": paymentCode}})
	}

	grouping := bson.M{"$group": bson.M{
		"_id":          0,
		"data":         bson.M{"$push": "$$ROOT"},
		"totalRecords": bson.M{"$sum": 1}},
	}

	project := bson.M{"$project": bson.M{
		"totalRecords": "$totalRecords",
		"data":         "$data",
	}}

	pipeline = append(pipeline, grouping)
	pipeline = append(pipeline, project)

	if err := table.Pipe(pipeline).All(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *repositoryCarts) Checkout(cartIds []*bson.ObjectId, paymentCode string) error {

	var err error

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionCart)

	who := bson.M{"_id" : bson.M{"$in" : cartIds}}
	change := bson.M{"$set": bson.M{
		"status"		: helper.CartStatusCheckout,
		"payment_code"	: paymentCode,
	}}
	_, err = table.UpdateAll(
		who,
		change,
	)





	if err != nil {

		return err
	}

	return nil
}

func (repo *repositoryCarts) FindByPaymentCode(paymentCode string) ([]*entity.Cart, error){
	var err error

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionCart)

	var cartEntities []*entity.Cart
	err = table.Find(bson.M{"payment_code" : paymentCode}).All(&cartEntities)
	if err != nil {
		return nil, err
	}

	return cartEntities, nil
}

func (repo *repositoryCarts) DeleteByPaymentCode(paymentCode string) error{
	var err error

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionCart)

	_, err = table.RemoveAll(bson.M{"payment_code": paymentCode})
	if err != nil {

		return err
	}
	return  nil

}