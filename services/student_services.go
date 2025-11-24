package services

import (
	"errors"
	"strings"

	"github.com/bayuf/Go-CRUD-Mahasiswa/dto"
	"github.com/bayuf/Go-CRUD-Mahasiswa/model"
	"github.com/bayuf/Go-CRUD-Mahasiswa/repository"
)

type StudentService struct {
	Repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return StudentService{
		Repo: repo,
	}
}

func (s StudentService) Create(req dto.CreateStudentRequest) error {
	// validasi isian kosong
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("error: name is empty")
	}
	if req.NIM == 0 {
		return errors.New("error: nim invalid")
	}
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("error: email is empty")
	}
	if !strings.Contains(req.Email, "@") {
		return errors.New("error: email invalid format")
	}
	if strings.TrimSpace(req.Major) == "" {
		return errors.New("error: major is empty")
	}

	// konversi dto ke model
	m := model.Student{
		Name:  req.Name,
		NIM:   req.NIM,
		Email: req.Email,
		Major: req.Major,
	}

	// simpan ke db melalui repo
	if err := s.Repo.Create(m); err != nil {
		return err
	}
	return nil
}

func (s StudentService) Read() ([]model.Student, error) {
	return s.Repo.Read() // melakukan query
}
