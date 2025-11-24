package main

import (
	"github.com/bayuf/Go-CRUD-Mahasiswa/db"
	"github.com/bayuf/Go-CRUD-Mahasiswa/dto"
	"github.com/bayuf/Go-CRUD-Mahasiswa/handler"
	"github.com/bayuf/Go-CRUD-Mahasiswa/repository"
	"github.com/bayuf/Go-CRUD-Mahasiswa/services"
)

func main() {
	// connect to DB
	db.Connect()

	// init
	repo := repository.NewStudentRepository()
	svc := services.NewStudentService(repo)
	handler := handler.NewStudentHandler(svc)

	// queri ke DB
	req := dto.CreateStudentRequest{
		Name:  "Bayu Firmansyah",
		NIM:   143218041,
		Email: "bayufirmansyah203@gmail.com",
		Major: "Teknik Informatika",
	}

	handler.Create(req)

}
