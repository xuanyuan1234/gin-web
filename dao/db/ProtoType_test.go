package db

import (
	"testing"
	"work/src/web/entity"
)

func init() {
	dns := "root:root1234@tcp(localhost:3306)/tag?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestInsertProtoType(t *testing.T) {
	prototype := &entity.ProtoTypeVo{}
	prototype.Id = 2
	prototype.Name = "lisi"
	prototype.Sn = "123"
	prototype.DealPerson = "lisi 123"

	protoTypeId, err := InsertProtoType(prototype)
	if err != nil {
		panic(err)
	}

	t.Logf("protoTypeId: %d", protoTypeId)
}

func TestGetProtoTypeById(t *testing.T) {
	prototype, err := GetProtoTypeById(1)
	if err != nil {
		panic(err)
	}

	t.Logf("prototype:%#v", prototype)
}

func TestGetProtoTypes(t *testing.T) {
	var prototypeIds []int64
	prototypeIds = append(prototypeIds, 1, 2, 3)
	prototypeList, err := GetProtoTypes(prototypeIds)
	if err != nil {
		panic(err)
	}

	for _, v := range prototypeList {
		t.Logf("prototype:%#v", v)
	}
}

func TestGetProtoTypeList(t *testing.T) {
	prototypeList, err := GetProtoTypeList()
	if err != nil {
		return
	}

	for _,v := range prototypeList {
		t.Logf("id:%d name:%#v", v.Id, v.Name)
	}
}
