package repository

import (
	"database/sql"

	"github.com/bayuf/Go-CRUD-Mahasiswa/db"
	"github.com/bayuf/Go-CRUD-Mahasiswa/model"
)

type IStudentRepository interface {
	Create(m model.Student) error
	Read() ([]model.Student, error)
}

type StudentRepository struct {
	DB *sql.DB
}

func NewStudentRepository() IStudentRepository {
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

func (r StudentRepository) Read() ([]model.Student, error) {
	var listAll []model.Student
	query := `SELECT nim, name, email, major FROM student`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() { // membaca baris satu per satu
		var (
			nim   int64
			name  string
			email string
			major string
		)

		if err := rows.Scan(&nim, &name, &email, &major); err != nil { // memindahkan data ke var
			return nil, err
		}

		//convert int64 -> uint
		student := model.Student{
			NIM:   uint(nim),
			Name:  name,
			Email: email,
			Major: major,
		}

		listAll = append(listAll, student)
	}

	return listAll, nil
}
