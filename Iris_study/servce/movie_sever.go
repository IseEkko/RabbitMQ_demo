package servce

import (
	"Iris_study/repositories"
	"fmt"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiveManger struct {
	repo repositories.MoviesReponsitory
}

func NewMovieServiceManger(repo repositories.MoviesReponsitory) MovieService {
	return &MovieServiveManger{repo: repo}
}

func (m *MovieServiveManger) ShowMovieName() string {
	fmt.Println("获取用户名称：" + m.repo.GetMovieName())
	return "获取用户名称：" + m.repo.GetMovieName()
}
