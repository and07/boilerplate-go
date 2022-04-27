package grpcserver

import (
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
)

var exerciseDefault = []*Exercise{
	{
		Name:                "Exercise1",
		Duration:            durationpb.New(time.Duration(5) * time.Second),
		Relax:               durationpb.New(time.Duration(5) * time.Second),
		Count:               10,
		NumberOfSets:        3,
		NumberOfRepetitions: 15,
		Type:                ExerciseType_other,
		Image:               "https://fitseven.ru/wp-content/uploads/2020/07/uprazhneniya-na-press-velosiped.jpg",
		Description:         "la blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla bla",
		Technique: `
		Stand with your hands aside.
		Pick up dumbbells.
		Lift the dumbbells ahead
		Lower the dumbbells`,
		Mistake: "Don’t lift the dumbbells too high",
		Weight:  5000,
	},
	{
		Name:                "Exercise2",
		Duration:            durationpb.New(time.Duration(5) * time.Second),
		Relax:               durationpb.New(time.Duration(5) * time.Second),
		Count:               10,
		NumberOfSets:        3,
		NumberOfRepetitions: 15,
		Type:                ExerciseType_arms,
		Image:               "https://fitseven.ru/wp-content/uploads/2020/07/uprazhneniya-na-press-skruchivaniya.jpg",
		Description:         "bla bla bla bla bla blabla bla bla bla bla babla bla",
		Technique: `
		Stand with your hands aside.
		Lower the dumbbells`,
		Mistake: "Don’t lift the dumbbells too high",
		Weight:  5000,
	},
	{
		Name:                "Exercise3",
		Duration:            durationpb.New(time.Duration(5) * time.Second),
		Relax:               durationpb.New(time.Duration(5) * time.Second),
		Count:               10,
		NumberOfSets:        3,
		NumberOfRepetitions: 15,
		Type:                ExerciseType_other,
		Description:         "bla bla bla bla bla blabla bla bla bla bla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla bla",
		Technique: `
		Stand with your hands aside.
		Pick up dumbbells.
		Lift the dumbbells ahead
		Lower the dumbbells`,
		Image:  "https://filzor.ru/wp-content/uploads/2021/07/685dcbc62da32d989482066ebc60de3f-768x401.jpg",
		Weight: 0,
	},
	{
		Name:                "Exercise4",
		Duration:            durationpb.New(time.Duration(5) * time.Second),
		Relax:               durationpb.New(time.Duration(5) * time.Second),
		Count:               10,
		NumberOfSets:        3,
		NumberOfRepetitions: 15,
		Type:                ExerciseType_legs,
		Image:               "https://fitseven.ru/wp-content/uploads/2020/07/uprazhneniya-na-press-planka.jpg",
		Description:         "b blabla blabla bla",
		Mistake:             "Don’t lift the dumbbells too high",
		Weight:              5000,
	},
	{
		Name:                "Exercise5",
		Duration:            durationpb.New(time.Duration(20) * time.Second),
		Relax:               durationpb.New(time.Duration(5) * time.Second),
		Count:               10,
		NumberOfSets:        3,
		NumberOfRepetitions: 15,
		Type:                ExerciseType_other,
		Image:               "https://fitseven.ru/wp-content/uploads/2020/07/uprazhneniya-na-press-planka-na-rukah.jpg",
		Description:         "bla bla bla bla bla blabla abla blabla blabla bla",
		Technique: `
		Stand with your hands aside.
		Pick up dumbbells.
		Lift the dumbbells ahead
		Lower the dumbbells`,
		Mistake: "Don’t lift the dumbbells too high",
		Weight:  0,
	},
	{
		Name:                "Exercise6",
		Duration:            durationpb.New(time.Duration(20) * time.Second),
		Relax:               durationpb.New(time.Duration(5) * time.Second),
		Count:               10,
		NumberOfSets:        3,
		NumberOfRepetitions: 15,
		Type:                ExerciseType_back,
		Description:         "bla bla bla bla bla blabla bla bla bla bla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla bla",
		Technique: `
		Stand with your hands aside.
		Lift the dumbbells ahead
		Lower the dumbbells`,
		Mistake: "Don’t lift the dumbbells too high",
		Image:   "https://cdn-st1.rtr-vesti.ru/vh/pictures/xw/317/436/2.jpg",
		Weight:  5000,
	},
	{
		Name:                "Exercise7",
		Duration:            durationpb.New(time.Duration(20) * time.Second),
		Relax:               durationpb.New(time.Duration(5) * time.Second),
		Count:               10,
		NumberOfSets:        3,
		NumberOfRepetitions: 15,
		Type:                ExerciseType_legs,
		Image:               "https://fitseven.ru/wp-content/uploads/2020/07/uprazhneniya-na-press-podyem-nog.jpg",
		Description:         "bla bla bla bla bla blabla blabla blabla blabla blabla blabla blabla blabla bla",
		Technique: `
		Stand with your hands aside.
		Lower the dumbbells`,
		Mistake: "Don’t lift the dumbbells too high",
		Weight:  0,
	},
}
