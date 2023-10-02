package models

type JobBoard struct {
	Title    string  `json:"title"`
	Location string  `json:"location"`
	JobType  JobType `json:"jobtype"`
}
