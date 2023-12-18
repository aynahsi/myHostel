package attendances

import (
	"encoding/json"
	"gofr.dev/pkg/gofr"
	"gofrproject/models"
)

func UpdateAttendance(ctx *gofr.Context) (interface{}, error) {
	recordID := ctx.PathParam("recordID")
	var attendance models.Attendance
	if err := ctx.Bind(&attendance); err != nil {
		return nil, err
	}

	_, err := ctx.DB().ExecContext(ctx, "UPDATE AttendanceRecords SET StudentID = ?, Date = ?, Status = ? WHERE RecordID = ?",
		attendance.StudentID, attendance.Date, attendance.Status, recordID)
	if err != nil {
		return nil, err
	}

	jsonResponse, _ := json.Marshal(map[string]string{"message": "Attendance updated successfully"})
	return string(jsonResponse), nil
}
