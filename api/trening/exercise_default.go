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
		Image:               "https://fitseven.ru/wp-content/uploads/2020/07/uprazhneniya-na-press-velosiped.jpg",
	},
	{
		Name:                "Exercise2",
		Duration:            durationpb.New(time.Duration(20) * time.Second),
		Relax:               durationpb.New(time.Duration(20) * time.Second),
		Count:               10,
		NumberOfSets:        3,
		NumberOfRepetitions: 15,
		Type:                ExerciseType_arms,
		Image:               "https://fitseven.ru/wp-content/uploads/2020/07/uprazhneniya-na-press-skruchivaniya.jpg",
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
		Image:               "https://fitseven.ru/wp-content/uploads/2020/07/uprazhneniya-na-press-planka.jpg",
	},
	{
		Name:                "Exercise5",
		Duration:            durationpb.New(time.Duration(20) * time.Second),
		Relax:               durationpb.New(time.Duration(20) * time.Second),
		Count:               10,
		NumberOfSets:        3,
		NumberOfRepetitions: 15,
		Type:                ExerciseType_other,
		Image:               "https://fitseven.ru/wp-content/uploads/2020/07/uprazhneniya-na-press-planka-na-rukah.jpg",
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
		Image:               "https://fitseven.ru/wp-content/uploads/2020/07/uprazhneniya-na-press-podyem-nog.jpg",
	},
}
