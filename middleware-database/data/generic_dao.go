package data

import (
	"github.com/smart-aksis/golang-middlewares/middleware-database/relational"
	"github.com/smart-aksis/golang-middlewares/middleware-rest/request_utils"
	"gorm.io/gorm"
)

type GenericDaoInterface interface {
	GetModel() (tx *gorm.DB)
}

//func Save(dao GenericDaoInterface, entity interface{}) (interface{}, error){
//	err := dao.GetModel().Save(entity).Error
//	if err != nil {
//		return nil, err
//	} else {
//		return entity, nil
//	}
//}

func Paginate(dao GenericDaoInterface, filters []request_utils.FilterField, paginationProperties request_utils.PaginationProperties, dest interface{}) error {
	var result *gorm.DB
	if len(filters) > 0 {
		result = dao.GetModel().Where(relational.GetFilter(filters...))
	} else {
		result = dao.GetModel()
	}

	var limit int
	var page int

	limit = paginationProperties.PageSize
	page = paginationProperties.PageNumber - 1

	if page < 0 {
		page = 0
	}

	if limit < 1 {
		limit = 10
	}

	return result.Limit(limit).Offset(page * limit).Find(dest).Error
}

