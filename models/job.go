package models

// Job data structure
type Job struct {
	// ID       primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	JobTitle string `json:"jobTitle" bson:"jobTitle"`
	Company  string `json:"company" bson:"company"`
	Location string `json:"location" bson:"location"`
	Remote   string `json:"remote" bson:"remote"`
	Job      string `json:"job" bson:"job"`
}
