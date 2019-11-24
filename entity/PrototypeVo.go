package entity

//定义样机信息结构体
type ProtoTypeVo struct {
	Id         int64  `db:"id"`
	Name       string `db:"name"`
	Sn         string  `db:"sn"`
	DealPerson string `db:"deal_person"`
}
