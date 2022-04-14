package data

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// ParametersUser ...
type ParametersUser struct {
	UID           string    `json:"uid" db:"uid"`
	UserID        string    `json:"user_id,omitempty" db:"user_id"`
	UserName      string    `json:"username,omitempty" db:"username"`
	Weight        int32     `json:"weight,omitempty" db:"weight"`
	Height        int32     `json:"height,omitempty" db:"height"`
	Age           int32     `json:"age,omitempty" db:"age"`
	Gender        int32     `json:"gender,omitempty" db:"gender"`
	Activity      int32     `json:"activity,omitempty" db:"activity"`
	Diet          int32     `json:"diet,omitempty" db:"diet"`
	DesiredWeight int32     `json:"desired_weight,omitempty" db:"desired_weight"`
	Eat           int32     `json:"eat,omitempty" db:"eat"`
	CreatedAt     time.Time `json:"createdat" db:"createdat"`
	UpdatedAt     time.Time `json:"updatedat" db:"updatedat"`
}

type Exercise struct {
	UID                 string        `json:"uid" db:"uid"`
	UserID              string        `json:"user_id,omitempty" db:"user_id"`
	Name                string        `json:"name,omitempty" db:"name"`
	Duration            time.Duration `json:"duration,omitempty" db:"duration"`
	Relax               time.Duration `json:"relax,omitempty" db:"relax"`
	Count               int32         `json:"count,omitempty" db:"count"`
	NumberOfSets        int32         `json:"number_of_sets,omitempty" db:"number_of_sets"`
	NumberOfRepetitions int32         `json:"number_of_repetitions,omitempty" db:"number_of_repetitions"`
	Type                int32         `json:"type,omitempty" db:"type"`
	Image               string        `json:"image,omitempty" db:"image"`
	Video               string        `json:"video,omitempty" db:"video"`
	Description         string        `json:"description,omitempty" db:"description"`
	Technique           string        `json:"technique,omitempty" db:"technique"`
	Mistake             string        `json:"mistake,omitempty" db:"mistake"`
	Weight              int32         `json:"weight,omitempty" db:"weight"`
	CreatedAt           time.Time     `json:"createdat" db:"createdat"`
	UpdatedAt           time.Time     `json:"updatedat" db:"updatedat"`
}

type ExerciseSlice []Exercise

// Make the Exercise struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a ExerciseSlice) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Make the Exercise struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (s *ExerciseSlice) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, s)
	case string:
		return json.Unmarshal([]byte(v), s)
	}
	return errors.New("type assertion failed")
}

type Trening struct {
	UID       string        `json:"uid" db:"uid"`
	UserID    string        `json:"user_id,omitempty" db:"user_id"`
	Name      string        `json:"name,omitempty"  db:"name"`
	Exercises ExerciseSlice `json:"exercises,omitempty"  db:"exercises"`
	Interval  time.Duration `json:"interval,omitempty"  db:"interval"`
	Type      int           `json:"type,omitempty" db:"type"`
	Status    int           `json:"status,omitempty" db:"status"`
	CreatedAt time.Time     `json:"createdat" db:"createdat"`
	UpdatedAt time.Time     `json:"updatedat" db:"updatedat"`
}
