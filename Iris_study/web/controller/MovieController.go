package controller

import (
	"Iris_study/repositories"
	"Iris_study/servce"
	"github.com/kataras/iris/v12/mvc"
)

type MovieController struct {
}

func (c *MovieController) Get() mvc.View {
	moviesReponsitory := repositories.NewMovieManger()
	movie := servce.NewMovieServiceManger(moviesReponsitory)
	moviesresult := movie.ShowMovieName()
	return mvc.View{
		Name: "movie/index.html",
		Data: moviesresult,
	}
}
