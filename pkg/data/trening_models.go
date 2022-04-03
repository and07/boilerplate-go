package data

import "time"

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
	UID                 string    `json:"uid" db:"uid"`
	UserID              string    `json:"user_id,omitempty" db:"user_id"`
	Name                string    `json:"name,omitempty" db:"name"`
	Duration            time.Time `json:"duration,omitempty" db:"duration"`
	Relax               time.Time `json:"relax,omitempty" db:"relax"`
	Count               int32     `json:"count,omitempty" db:"count"`
	NumberOfSets        int32     `json:"number_of_sets,omitempty" db:"number_of_sets"`
	NumberOfRepetitions int32     `json:"number_of_repetitions,omitempty" db:"number_of_repetitions"`
	Type                int32     `json:"type,omitempty" db:"type"`
	CreatedAt           time.Time `json:"createdat" db:"createdat"`
	UpdatedAt           time.Time `json:"updatedat" db:"updatedat"`
}
