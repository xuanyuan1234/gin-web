package service

import (
	"work/src/web/dao/db"
	"work/src/web/entity"
)

func GetAllProtoType() (prototypeList []*entity.ProtoTypeVo, err error) {
	prototypeList, err = db.GetProtoTypeList()
	if err != nil {
		return
	}

	return
}
