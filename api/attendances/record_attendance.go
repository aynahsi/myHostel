package attendances

import (
	"encoding/json"
	"gofr.dev/pkg/gofr"
	"gofrproject/models"
)

func RecordAttendance(ctx *gofr.Context) (interface{}, error) {
	var attendance models.Attendance
	if err := ctx.Bind(&attendance); err != nil {
		return nil, err
	}

	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO AttendanceRecords (StudentID, Date, Status) VALUES (?, ?, ?)",
		attendance.StudentID, attendance.Date, attendance.Status)
	if err != nil {
		return nil, err
	}

	jsonResponse, _ := json.Marshal(map[string]string{"message": "Attendance recorded successfully"})
	return string(jsonResponse), nil
}
