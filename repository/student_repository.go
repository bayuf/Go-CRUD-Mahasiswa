package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/bayuf/Go-CRUD-Mahasiswa/db"
	"github.com/bayuf/Go-CRUD-Mahasiswa/model"
)

type IStudentRepository interface {
	Create(model.Student) error
	Read() ([]model.Student, error)
	Update(model.Student) error
	FindByNim(uint64) (model.Student, error)
	Delete(uint64) error
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

func (r StudentRepository) Update(m model.Student) error {
	query := "UPDATE student SET "
	args := []interface{}{}
	counter := 1

	// membuat query berdasarkan apa yang mau diubah
	if m.Name != "" {
		query += fmt.Sprintf("name=$%d, ", counter) // name=$1
		args = append(args, m.Name)                 // menambahkan ke slice args
		counter++
	}
	if m.Email != "" {
		query += fmt.Sprintf("email=$%d, ", counter) // email=$2 or $1
		args = append(args, m.Email)
		counter++
	}
	if m.Major != "" {
		query += fmt.Sprintf("major=$%d, ", counter) // major=$3 or $2 or $1
		args = append(args, m.Major)
		counter++
	}

	// hapus koma di akhir query setelah di gabung
	query = strings.TrimRight(query, ", ")

	// WHERE berdasarkan NIM
	query += fmt.Sprintf(" WHERE nim=$%d", counter)
	args = append(args, m.NIM)

	fmt.Println("QUERY:", query)
	fmt.Println("ARGS:", args)
	_, err := r.DB.Exec(query, args...)
	return err
}

func (r StudentRepository) FindByNim(req uint64) (model.Student, error) {
	query := "SELECT nim, name, email, major FROM student WHERE "
	nimReq := req
	query += fmt.Sprintf("nim=%d", nimReq)

	rows, err := r.DB.Query(query)
	if err != nil {
		return model.Student{}, err
	}

	defer rows.Close()

	var (
		nim   uint64
		name  string
		email string
		major string
	)

	for rows.Next() { // membaca baris satu per satu
		if err := rows.Scan(&nim, &name, &email, &major); err != nil { // memindahkan data ke var
			return model.Student{}, err
		}
	}
	//convert int64 -> uint
	return model.Student{
		NIM:   uint(nim),
		Name:  name,
		Email: email,
		Major: major,
	}, nil
}

func (r StudentRepository) Delete(req uint64) error {
	query := "DELETE FROM student WHERE "
	nimReq := req
	query += fmt.Sprintf("nim=%d", nimReq)

	rows, err := r.DB.Query(query)
	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}
