package students

import (
	"gofr.dev/pkg/gofr"
	"gofrproject/models"
)

func GetStudent(ctx *gofr.Context) (interface{}, error) {
	studentID := ctx.PathParam("studentID")
	var student models.Student

	err := ctx.DB().QueryRowContext(ctx, "SELECT StudentID, Name, Phone, Address FROM Students WHERE StudentID = ?", studentID).
		Scan(&student.StudentID, &student.Name, &student.Phone, &student.Address)
	if err != nil {
		return nil, err
	}

	return student, nil
}
