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

type repositoryInventoryAdjustments struct {
	dbSession *mgo.Session
	database  string
}

var collectionInventoryAdjustment = "inventory_adjustments"

func NewInventoryAdjustmentsRepository(sess *mgo.Session, database string) interfaces.InventoryAdjustmentInterface {
	return &repositoryInventoryAdjustments{sess, database}

}
func (repo *repositoryInventoryAdjustments) Create(e *entity.InventoryAdjustment) (bool, error) {
	var err error
	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionInventoryAdjustment)

	if err = table.Insert(&e); err != nil {
		return false, err
	}

	return true, err
}

func (repo *repositoryInventoryAdjustments) FindByInventoryAdjustmentName(username string) (*entity.InventoryAdjustment, error) {
	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionInventoryAdjustment)

	var result entity.InventoryAdjustment
	err := table.Find(bson.M{"username": username}).One(&result)

	return &result, err
}

func (repo *repositoryInventoryAdjustments) FindById(id string) (*entity.InventoryAdjustment, error) {

	isObjectIDHex := helper.IsObjectIdHexValidation(id)

	if !isObjectIDHex {
		return nil, errors.New(helper.ErrorIsNOTObjectIdHex(id))
	}

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionInventoryAdjustment)

	var result entity.InventoryAdjustment
	err := table.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)

	return &result, err
}

func (repo *repositoryInventoryAdjustments) FindAll(searchParam request.SearchParamWithPagingInventoryAdjustmentRequest) ([]*entity.InventoryAdjustment, error, int) {

	result := []*entity.InventoryAdjustment{}

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionInventoryAdjustment)

	pipeline := []bson.M{}
	pipelineCount := []bson.M{}

	productName := helper.TrimWhiteSpace(searchParam.ProductName)
	if productName != "" {
		pipeline = append(pipeline, bson.M{"$match": bson.M{"name": productName}})
		pipelineCount = append(pipelineCount, bson.M{"$match": bson.M{"name": productName}})

	}

	productCode := helper.TrimWhiteSpace(searchParam.ProductCode)
	if productName != "" {
		pipeline = append(pipeline, bson.M{"$match": bson.M{"code": productCode}})
		pipelineCount = append(pipelineCount, bson.M{"$match": bson.M{"code": productCode}})
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

func (repo *repositoryInventoryAdjustments) FindAllWithPaging(searchParam request.SearchParamWithPagingInventoryAdjustmentRequest) ([]*entity.InventoryAdjustment, error) {

	result := []*entity.InventoryAdjustment{}

	ds := repo.dbSession.Copy()
	defer ds.Close()
	table := ds.DB(repo.database).C(collectionInventoryAdjustment)

	pipeline := []bson.M{}

	productName := helper.TrimWhiteSpace(searchParam.ProductName)
	if productName != "" {
		pipeline = append(pipeline, bson.M{"$match": bson.M{"name": productName}})
	}

	productCode := helper.TrimWhiteSpace(searchParam.ProductCode)
	if productCode != "" {
		pipeline = append(pipeline, bson.M{"$match": bson.M{"code": productCode}})
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