package attendances

import (
	"encoding/json"
	"gofr.dev/pkg/gofr"
)

func DeleteAttendance(ctx *gofr.Context) (interface{}, error) {
	recordID := ctx.PathParam("recordID")

	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM AttendanceRecords WHERE RecordID = ?", recordID)
	if err != nil {
		return nil, err
	}

	jsonResponse, _ := json.Marshal(map[string]string{"message": "Attendance deleted successfully"})
	return string(jsonResponse), nil
}
