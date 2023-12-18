package students

import (
	"gofr.dev/pkg/gofr"
	"gofrproject/models"
)

func AddStudent(ctx *gofr.Context) (interface{}, error) {
	var student models.Student
	if err := ctx.Bind(&student); err != nil {
		return nil, err
	}

	// Inserting the student into the database
	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO Students (StudentID, Name, Phone, Address) VALUES (?, ?, ?, ?)",
		student.StudentID, student.Name, student.Phone, student.Address)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message":    "Student added successfully",
		"student_id": student.StudentID,
	}, nil
}
