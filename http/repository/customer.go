package repository

import (
	"evermos-test/database/entity"
	"evermos-test/helper"
	"evermos-test/http/interfaces"
	"evermos-test/http/request"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repositoryCustomers struct {
	dbSession *mgo.Session
	database  string
}

var collectionCustomer = "customers"

func NewCustomersRepository(sess *mgo.Session, database string) interfaces.CustomerInterface {
	return &repositoryCustomers{sess, database}

}
func (repo *repositoryCustomers) Create(e *entity.Customer) (bool, error) {
	var err error
	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionCustomer)

	index := mgo.Index{
		Key:    []string{"code"},
		Unique: true,
	}
	table.EnsureIndex(index)

	if err = table.Insert(&e); err != nil {
		return false, err
	}

	return true, err
}

func (repo *repositoryCustomers) Update(id *bson.ObjectId, e *entity.Customer) (bool, error) {
	var err error

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionCustomer)
	err = table.Update(
		bson.M{"_id": id},
		bson.M{"$set": &e},
	)

	if err != nil {
		return false, err
	}

	return true, err
}

func (repo *repositoryCustomers) FindByCustomerName(name string) (*entity.Customer, error) {
	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionCustomer)

	var result entity.Customer
	err := table.Find(bson.M{"name": name}).One(&result)

	return &result, err
}

func (repo *repositoryCustomers) FindById(id *bson.ObjectId) (*entity.Customer, error) {

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionCustomer)

	var result entity.Customer
	err := table.Find(bson.M{"_id": id}).One(&result)

	return &result, err
}

func (repo *repositoryCustomers) FindAll(searchParam request.SearchParamWithPagingCustomerRequest) ([]*entity.Customer, error, int) {

	result := []*entity.Customer{}

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionCustomer)

	pipeline := []bson.M{}
	pipelineCount := []bson.M{}

	name := helper.TrimWhiteSpace(searchParam.Name)
	if name != "" {
		pipeline = append(pipeline, bson.M{"$match": bson.M{"name": name}})
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

func (repo *repositoryCustomers) FindAllWithPaging(searchParam request.SearchParamWithPagingCustomerRequest) ([]*entity.Customer, error) {

	result := []*entity.Customer{}

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionCustomer)

	pipeline := []bson.M{}

	name := helper.TrimWhiteSpace(searchParam.Name)
	if name != "" {
		pipeline = append(pipeline, bson.M{"$match": bson.M{"name": name}})
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
