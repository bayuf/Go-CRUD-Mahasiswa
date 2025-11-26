package dto

type CreateStudentRequest struct {
	Name  string
	NIM   uint
	Email string
	Major string
}

type UpdateStudentRequest struct {
	NIM   uint
	Name  *string
	Email *string
	Major *string
}
