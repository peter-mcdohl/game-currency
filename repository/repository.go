package repository

type Repository interface {
	Insert(data interface{}) error
	FindAll(data interface{}) error
	FindByID(data interface{}, id int) error
	FindByField(data interface{}, fieldName string, fieldValue interface{}) error
	FindByConditionStruct(data interface{}, conds interface{}) error
}
