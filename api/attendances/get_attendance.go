package attendances

import (
	"encoding/json"
	"gofr.dev/pkg/gofr"
	"gofrproject/models"
)

func GetAttendance(ctx *gofr.Context) (interface{}, error) {
	studentID := ctx.PathParam("studentID")
	rows, err := ctx.DB().QueryContext(ctx, "SELECT RecordID, StudentID, Date, Status FROM AttendanceRecords WHERE StudentID = ?", studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendances []models.Attendance
	for rows.Next() {
		var attendance models.Attendance
		err := rows.Scan(&attendance.RecordID, &attendance.StudentID, &attendance.Date, &attendance.Status)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, attendance)
	}

	jsonResponse, _ := json.Marshal(attendances)
	return string(jsonResponse), nil
}
