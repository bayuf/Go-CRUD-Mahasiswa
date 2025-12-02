package services

import (
	"errors"
	"strings"

	"github.com/bayuf/Go-CRUD-Mahasiswa/dto"
	"github.com/bayuf/Go-CRUD-Mahasiswa/model"
	"github.com/bayuf/Go-CRUD-Mahasiswa/repository"
)

type StudentService struct {
	Repo repository.IStudentRepository
}

func NewStudentService(repo repository.IStudentRepository) StudentService {
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

func (s StudentService) Update(req dto.UpdateStudentRequest) error {
	// validasi isian kosong
	if req.NIM == 0 {
		return errors.New("error: nim invalid")
	}

	if req.Name == nil || req.Email == nil || req.Major == nil {
		return errors.New("error: tidak ada data yang dirubah")
	}

	// konversi dto ke model
	m := model.Student{
		NIM:   req.NIM,
		Name:  *req.Name,
		Email: *req.Email,
		Major: *req.Major,
	}

	// update ke db melalui repo
	if err := s.Repo.Update(m); err != nil {
		return err
	}

	return nil
}

func (s StudentService) FindByNim(req uint64) (model.Student, error) {
	if req == 0 {
		return model.Student{}, errors.New("error: nim invalid")
	}

	student, err := s.Repo.FindByNim(req)
	if err != nil {
		return model.Student{}, err
	}

	return model.Student{
		NIM:   student.NIM,
		Name:  student.Name,
		Email: student.Email,
		Major: student.Major,
	}, nil
}

func (s StudentService) Delete(req uint64) error {
	if req == 0 {
		return errors.New("error: nim invalid")
	}

	if err := s.Repo.Delete(req); err != nil {
		return err
	}

	return nil
}
