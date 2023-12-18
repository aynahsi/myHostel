package students

import (
	"gofr.dev/pkg/gofr"
	"gofrproject/models"
)

func UpdateStudent(ctx *gofr.Context) (interface{}, error) {
	studentID := ctx.PathParam("studentID")
	var student models.Student
	if err := ctx.Bind(&student); err != nil {
		return nil, err
	}

	_, err := ctx.DB().ExecContext(ctx, "UPDATE Students SET Name = ?, Phone = ?, Address = ? WHERE StudentID = ?",
		student.Name, student.Phone, student.Address, studentID)
	if err != nil {
		return nil, err
	}

	return map[string]string{"message": "Student updated successfully"}, nil
}
