package repository

import (
	"database/sql"

	"github.com/bayuf/Go-CRUD-Mahasiswa/db"
	"github.com/bayuf/Go-CRUD-Mahasiswa/model"
)

type StudentRepository struct {
	DB *sql.DB
}

func NewStudentRepository() StudentRepository {
	return StudentRepository{DB: db.DB}
}

func (r StudentRepository) Create(m model.Student) error {
	query := `
			INSERT INTO student (name, nim, email, major)
			VALUES ($1, $2, $3, $4)		
	`
	_, err := r.DB.Exec(query, m.Name, m.NIM, m.Email, m.Major) // exec digunakan untuk queri tanpa hasil row
	if err != nil {
		return err
	}

	return nil
}
