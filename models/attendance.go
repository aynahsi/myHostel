package models

import "time"

type Attendance struct {
	RecordID  int       `json:"recordID" db:"recordID"`
	StudentID int64     `json:"studentID" db:"studentID"`
	Date      time.Time `json:"date" db:"date"`
	Status    string    `json:"status" db:"status"`
}
