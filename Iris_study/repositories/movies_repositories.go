package repositories

import "Iris_study/datamodels"

type MoviesReponsitory interface {
	GetMovieName() string
}

type MovieManger struct {
}

func NewMovieManger() MoviesReponsitory {
	return &MovieManger{}
}
func (m *MovieManger) GetMovieName() string {
	//获取名称（模拟数据库操作）
	movir := &datamodels.Movie{
		Name: "李忠政",
	}
	return movir.Name
}
