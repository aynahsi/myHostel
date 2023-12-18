package models

type Student struct {
	StudentID int64  `json:"studentID" db:"StudentID"`
	Name      string `json:"name" db:"Name"`
	Phone     string `json:"phone" db:"Phone"`
	Address   string `json:"address" db:"Address"`
}
