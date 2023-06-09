package models

import "LiuMa-backend-go/internal/database"

type CaseType int
type CaseDataType int
type CaseData int

const (
	UnknownCaseType CaseType = iota
	GET
	POST
	PUT
	DELETE
)
const (
	UnknownCaseDataType CaseDataType = iota
	Params
	FROM_DATA
	JSON
	FILE
)

const (
	UnknownCaseData CaseData = iota
	DATA_DATA
	DATA_JSON
)

type Case struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Level       string       `json:"level"`
	Type        CaseType     `json:"type"`
	Description string       `json:"description"`
	Url         string       `json:"url"`
	Headers     string       `json:"headers"`
	DataType    CaseDataType `json:"headersType"`
	Data        CaseData     `json:"data"`
	CreateTime  int64        `json:"createTime"`
	UpdateTime  int64        `json:"updateTime"`
	Status      string       `json:"status"`

	GroupID int64 `gorm:"group_id"`
}

func (Case) TableName() string {
	return "case"
}

// 获取单个用例
func (model *Case) CaseGet(id int) (case1 Case, err error) {
	err = database.DB.Where("id = ?", id).Find(&case1).Error
	return
}

// 获取全部用例
func (model *Case) CaseGetList(page int, pageSize int) (cases []Case, err error) {
	offset := pageSize * (page - 1)
	err = database.DB.Offset(offset).Limit(pageSize).Find(&cases).Error
	return
}

// 创建用例
func (model *Case) CaseAdd() (id int64, err error) {
	err = database.DB.Create(model).Error
	return
}

// 更新用例
func (model *Case) CaseUpdate(id int) (err error) {
	err = database.DB.Model(model).Where("id = ?", id).Updates(*model).Error
	return
}

// 删除用例
func (model *User) CaseDelete(id int) (err error) {
	err = database.DB.Where("id = ?", id).Delete(User{}).Error
	return
}
