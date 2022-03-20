package servce

import "Iris_study/repositories"

type StudentServce interface {
	ShowStudent() string
}

type StudentServerManger struct {
	repo repositories.StudentReponsitory
}

func NewStudentServceManger(reop repositories.StudentReponsitory) StudentServce {
	return &StudentServerManger{repo: reop}
}

func (s *StudentServerManger) ShowStudent() string {
	return s.repo.GetStudent()
}
