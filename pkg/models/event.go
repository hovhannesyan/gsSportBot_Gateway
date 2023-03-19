package models

type Event struct {
	Id           int64
	DisciplineId int64
	Place        string
	Date         string
	StartTime    string
	EndTime      string
	Price        string
	Limit        string
	Description  string
}
