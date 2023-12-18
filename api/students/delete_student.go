package students

import (
	"encoding/json"
	"gofr.dev/pkg/gofr"
)

// DeleteStudent handles the deletion of a student record
func DeleteStudent(ctx *gofr.Context) (interface{}, error) {
	studentID := ctx.PathParam("studentID")

	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM Students WHERE StudentID = ?", studentID)
	if err != nil {
		return nil, err
	}

	jsonResponse, _ := json.Marshal(map[string]string{"message": "Student deleted successfully"})
	return string(jsonResponse), nil
}
