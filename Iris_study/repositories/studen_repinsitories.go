package repositories

import (
	"Iris_study/datamodels"
	"fmt"
)

//从数据库中获取信息
type StudentReponsitory interface {
	GetStudent() string
}

type StudentManger struct {
}

func (s *StudentManger) GetStudent() string {
	stu := &datamodels.Student{
		Id:   1,
		Name: "李忠政",
		Age:  21,
		Sex:  "男",
	}
	return fmt.Sprintf("编号：%d,姓名：%s", stu.Id, stu.Name)
}
func NewStundetManger() StudentReponsitory {
	return &StudentManger{}
}
