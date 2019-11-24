package entity

type HistoryRecordVo struct {
	id int64 `db:"id"`
	pId int64 `db:"pId"`
	content string `db:"content"`
}
