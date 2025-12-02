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

func (h StudentHandler) Create() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
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

func (h StudentHandler) Update() (dto.UpdateStudentRequest, error) {

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n') // buang \n di awal
	// input NIM
	fmt.Println("Masukkan NIM Kamu:")
	fmt.Print("NIM: ")
	nimStr, _ := reader.ReadString('\n')
	nimStr = strings.TrimSpace(nimStr)
	// convert ke uint
	nimUint, err := strconv.ParseUint(nimStr, 10, 64) //merubah string ke uint
	if err != nil {
		fmt.Println("Nim harus berupa angka")
		return dto.UpdateStudentRequest{}, err
	}

	// cari nim sebelum update
	student, err := h.service.FindByNim(nimUint)
	if err != nil {
		return dto.UpdateStudentRequest{}, err
	}

	// tampilkan data sebelum di update
	fmt.Println("NIM ditemukan:")
	fmt.Printf("Nama: %s | NIM: %d | Email: %s | Major: %s\n", student.Name, student.NIM, student.Email, student.Major)

	// input Name
	fmt.Println("Masukkan data baru")
	fmt.Print("Input New Name: ")
	name, _ := reader.ReadString('\n') //membaca satu baris input
	name = strings.TrimSpace(name)

	// input email
	fmt.Print("Input New Email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	// input jurusan
	fmt.Print("Input New Major: ")
	major, _ := reader.ReadString('\n')
	major = strings.TrimSpace(major)

	// buat dto
	reqDto := dto.UpdateStudentRequest{
		Name:  &name,
		NIM:   uint(nimUint),
		Email: &email,
		Major: &major,
	}

	// panggil service
	if err := h.service.Update(reqDto); err != nil {
		fmt.Println("error:", err)
		return dto.UpdateStudentRequest{}, err
	}

	return dto.UpdateStudentRequest{}, nil
}

func (h StudentHandler) Delete() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	fmt.Print("Masukkan NIM Mahasiswa yang akan dihapus: ")
	nim, _ := reader.ReadString('\n')
	nim = strings.TrimSpace(nim)

	uintNim, _ := strconv.ParseUint(nim, 10, 64)

	student, err := h.service.FindByNim(uintNim)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Mahasiswa ditemukan:\n", student)

	for {
		fmt.Println("hapus data mahasiswa tersebut? y/n")
		var choice string
		fmt.Scan(&choice)

		if strings.TrimSpace(strings.ToLower(choice)) == "y" {
			h.service.Delete(uintNim)

			fmt.Println("data berhasil dihapus")
			break
		} else if strings.TrimSpace(strings.ToLower(choice)) == "n" {
			fmt.Println("data gagal dihapus")
			break
		} else {
			fmt.Println("invalid input (y/n)")
		}

	}

}
