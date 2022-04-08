package grpcserver

import (
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
)

var exerciseDefault = []*Exercise{
	{
		Name:                "Exercise1",
		Duration:            durationpb.New(time.Duration(20) * time.Second),
		Relax:               durationpb.New(time.Duration(20) * time.Second),
		Count:               10,
		NumberOfSets:        3,
		NumberOfRepetitions: 15,
		Type:                ExerciseType_other,
	},
	{
		Name:                "Exercise2",
		Duration:            durationpb.New(time.Duration(20) * time.Second),
		Relax:               durationpb.New(time.Duration(20) * time.Second),
		Count:               10,
		NumberOfSets:        3,
		NumberOfRepetitions: 15,
		Type:                ExerciseType_arms,
	},
	{
		Name:                "Exercise3",
		Duration:            durationpb.New(time.Duration(20) * time.Second),
		Relax:               durationpb.New(time.Duration(20) * time.Second),
		Count:               10,
		NumberOfSets:        3,
		NumberOfRepetitions: 15,
		Type:                ExerciseType_other,
	},
	{
		Name:                "Exercise4",
		Duration:            durationpb.New(time.Duration(20) * time.Second),
		Relax:               durationpb.New(time.Duration(20) * time.Second),
		Count:               10,
		NumberOfSets:        3,
		NumberOfRepetitions: 15,
		Type:                ExerciseType_legs,
	},
	{
		Name:                "Exercise5",
		Duration:            durationpb.New(time.Duration(20) * time.Second),
		Relax:               durationpb.New(time.Duration(20) * time.Second),
		Count:               10,
		NumberOfSets:        3,
		NumberOfRepetitions: 15,
		Type:                ExerciseType_other,
	},
	{
		Name:                "Exercise6",
		Duration:            durationpb.New(time.Duration(20) * time.Second),
		Relax:               durationpb.New(time.Duration(20) * time.Second),
		Count:               10,
		NumberOfSets:        3,
		NumberOfRepetitions: 15,
		Type:                ExerciseType_arms,
	},
	{
		Name:                "Exercise7",
		Duration:            durationpb.New(time.Duration(20) * time.Second),
		Relax:               durationpb.New(time.Duration(20) * time.Second),
		Count:               10,
		NumberOfSets:        3,
		NumberOfRepetitions: 15,
		Type:                ExerciseType_legs,
	},
}
