package students

import (
	"encoding/json"
	"gofr.dev/pkg/gofr"
	"gofrproject/models"
)

func GetAllStudents(ctx *gofr.Context) (interface{}, error) {
	rows, err := ctx.DB().QueryContext(ctx, "SELECT StudentID, Name, Phone, Address FROM Students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var studentList []models.Student

	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.StudentID, &student.Name, &student.Phone, &student.Address)
		if err != nil {
			return nil, err
		}
		studentList = append(studentList, student)
	}

	jsonResponse, err := json.Marshal(studentList)
	if err != nil {
		return nil, err
	}
	return string(jsonResponse), nil
}
