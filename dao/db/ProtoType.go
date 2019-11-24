package db

import (
	"github.com/jmoiron/sqlx"
	"work/src/web/entity"
)

//插入样机数据
func InsertProtoType(prototype *entity.ProtoTypeVo) (protoTypeId int64, err error) {
	insertSql := "insert into t_prototype_info (name, sn, deal_person) values (?, ?, ?)"
	result, err := DB.Exec(insertSql, prototype.Name, prototype.Sn, prototype.DealPerson)

	if err !=nil {
		return
	}

	protoTypeId, err = result.LastInsertId()
	return
}

//获取单个样机
func GetProtoTypeById(id int64) (prototype *entity.ProtoTypeVo, err error) {
	prototype = &entity.ProtoTypeVo{}
	sqlStr := "select name, sn, deal_person from t_prototype_info where id = ?"
	err = DB.Get(prototype, sqlStr, id)
	return
}

//获取多个样机
func GetProtoTypes(prototypeIds []int64) (prototype []*entity.ProtoTypeVo, err error) {
	sqlstr, args, err := sqlx.In("select * from t_prototype_info where id in(?)", prototypeIds)
	if err != nil {
		return
	}

	err = DB.Select(&prototype, sqlstr, args...)
	return
}

//查询所有样机数据
func GetProtoTypeList() (prototype []*entity.ProtoTypeVo, err error) {
	sqlStr := "select * from t_prototype_info"
	err = DB.Select(&prototype, sqlStr)
	return
}
