package handler

import (
	"fmt"

	"github.com/bayuf/Go-CRUD-Mahasiswa/dto"
	"github.com/bayuf/Go-CRUD-Mahasiswa/services"
)

type StudentHandler struct {
	service services.StudentService
}

func NewStudentHandler(service services.StudentService) StudentHandler {
	return StudentHandler{
		service: service,
	}
}

func (h StudentHandler) Create(req dto.CreateStudentRequest) {
	if err := h.service.Create(req); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Student added")
}
