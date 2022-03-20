package controller

import (
	"Iris_study/repositories"
	"Iris_study/servce"
	"github.com/kataras/iris/v12/mvc"
)

type StudentController struct {
}

func (s *StudentController) Get() mvc.View {
	StudentRespons := repositories.NewStundetManger()
	student := servce.NewStudentServceManger(StudentRespons)
	studentResult := student.ShowStudent()
	return mvc.View{
		Name: "student/index.html",
		Data: studentResult,
	}
}
