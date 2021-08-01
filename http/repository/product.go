package repository

import (
	"errors"
	"evermos-test/database/entity"
	"evermos-test/helper"
	"evermos-test/http/interfaces"
	"evermos-test/http/request"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repositoryProducts struct {
	dbSession *mgo.Session
	database  string
}

var collectionProduct = "products"

func NewProductsRepository(sess *mgo.Session, database string) interfaces.ProductInterface {
	return &repositoryProducts{sess, database}

}
func (repo *repositoryProducts) Create(e *entity.Product) (bool, error) {
	var err error
	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionProduct)

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

func (repo *repositoryProducts) Update(id string, e *entity.Product) (bool, error) {
	var err error

	isObjectIDHex := helper.IsObjectIdHexValidation(id)

	if !isObjectIDHex {
		return false, errors.New(helper.ErrorIsNOTObjectIdHex(id))
	}

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionProduct)
	err = table.Update(
		bson.M{"_id": bson.ObjectIdHex(id)},
		bson.M{"$set": &e},
	)

	if err != nil {
		return false, err
	}

	return true, err
}

func (repo *repositoryProducts) FindByProductName(name string) (*entity.Product, error) {
	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionProduct)

	var result entity.Product
	err := table.Find(bson.M{"name": name}).One(&result)

	return &result, err
}

func (repo *repositoryProducts) FindById(id string) (*entity.Product, error) {

	isObjectIDHex := helper.IsObjectIdHexValidation(id)

	if !isObjectIDHex {
		return nil, errors.New(helper.ErrorIsNOTObjectIdHex(id))
	}

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionProduct)

	var result entity.Product
	err := table.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)

	return &result, err
}

func (repo *repositoryProducts) FindAll(searchParam request.SearchParamWithPagingProductRequest) ([]*entity.Product, error, int) {

	result := []*entity.Product{}

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionProduct)

	pipeline := []bson.M{}
	pipelineCount := []bson.M{}

	block := helper.TrimWhiteSpace(searchParam.Name)
	if block != "" {
		pipeline = append(pipeline, bson.M{"$match": bson.M{"name": block}})
		pipelineCount = append(pipelineCount, bson.M{"$match": bson.M{"block": block}})
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

func (repo *repositoryProducts) FindAllWithPaging(searchParam request.SearchParamWithPagingProductRequest) ([]*entity.Product, error) {

	result := []*entity.Product{}

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionProduct)

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
