package main

import (
	"fmt"

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
	req := dto.CreateStudentRequest{}

	for {
		fmt.Println("===== MENU MAHASISWA =====")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Lihat Semua Mahasiswa")
		fmt.Println("3. Update Data Mahasiswa")
		fmt.Println("4. Delete Data Mahasiswa")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih Menu: ")
		menu := 0
		fmt.Scan(&menu)

		switch menu {
		case 1:
			// melakukan Create
			handler.Create(req)
		case 2:
			// menampilkan semua list mahasiswa
			handler.Read()
		case 3:
			// Update
		case 4:
			// Delete
		case 5:
			return
		}
	}
}
