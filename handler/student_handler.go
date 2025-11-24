package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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
	reader := bufio.NewReader(os.Stdin)

	// input Name
	fmt.Println("Masukkan data diri")
	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n') //membaca satu baris input
	name = strings.TrimSpace(name)

	// input NIM
	fmt.Print("NIM: ")
	nimStr, _ := reader.ReadString('\n')
	nimStr = strings.TrimSpace(nimStr)
	// convert ke uint
	nimUint, err := strconv.ParseUint(nimStr, 10, 64) //merubah string ke uint
	if err != nil {
		fmt.Println("Nim harus berupa angka")
		return
	}

	// input email
	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	// input jurusan
	fmt.Print("Jurusan: ")
	major, _ := reader.ReadString('\n')
	major = strings.TrimSpace(major)

	// buat dto
	reqDto := dto.CreateStudentRequest{
		Name:  name,
		NIM:   uint(nimUint),
		Email: email,
		Major: major,
	}

	// panggil service
	if err := h.service.Create(reqDto); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Student added")
}

func (h StudentHandler) Read() {
	datas, err := h.service.Read()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("===== DATA MAHASISWA =====")
	for _, data := range datas {
		fmt.Printf(" %s | %d | %s | %s\n", data.Name, data.NIM, data.Email, data.Major)
	}

}

// func (h StudentHandler) Create(req dto.CreateStudentRequest) {
// 	if err := h.service.Create(req); err != nil {
// 		fmt.Println("error:", err)
// 		return
// 	}

// 	fmt.Println("Student added")
// }
