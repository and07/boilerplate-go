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
	Name      string        `json:"name,omitempty"  db:"name"`
	Interval  time.Duration `json:"interval,omitempty"  db:"interval"`
	Exercises ExerciseSlice `json:"exercises,omitempty"  db:"exercises"`
	UserID    string        `json:"user_id,omitempty" db:"user_id"`
	CreatedAt time.Time     `json:"createdat" db:"createdat"`
	UpdatedAt time.Time     `json:"updatedat" db:"updatedat"`
}
