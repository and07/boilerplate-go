package handlers

import (
	"context"

	"github.com/and07/boilerplate-go/internal/app/trening/models"
)

type TreningHandler interface {
	CreateParametersUser(ctx context.Context, request *models.CreateParametersUserRequest) (response *models.CreateParametersUserResponse, err error)
	DetailParametersUser(ctx context.Context, request *models.DetailParametersUserRequest) (response *models.DetailParametersUserResponse, err error)

	CreateTrening(ctx context.Context, request *models.CreateTreningRequest) (response *models.CreateTreningResponse, err error)
	ListTrening(ctx context.Context, request *models.ListTreningRequest) (response *models.ListTreningResponse, err error)
	DetailTrening(ctx context.Context, request *models.DetailTreningRequest) (response *models.DetailTreningResponse, err error)

	CreateExercise(ctx context.Context, request *models.CreateExerciseRequest) (response *models.CreateExerciseResponse, err error)
	ListExercise(ctx context.Context, request *models.ListExerciseRequest) (response *models.ListExerciseResponse, err error)
	DetailExercise(ctx context.Context, request *models.DetailExerciseRequest) (response *models.DetailExerciseResponse, err error)
}

type service struct{}

func (t *service) CreateParametersUser(ctx context.Context, request *models.CreateParametersUserRequest) (response *models.CreateParametersUserResponse, err error) {

	return
}

func (t *service) DetailParametersUser(ctx context.Context, request *models.DetailParametersUserRequest) (response *models.DetailParametersUserResponse, err error) {
	return
}

func (t *service) CreateTrening(ctx context.Context, request *models.CreateTreningRequest) (response *models.CreateTreningResponse, err error) {
	return
}
func (t *service) ListTrening(ctx context.Context, request *models.ListTreningRequest) (response *models.ListTreningResponse, err error) {
	return
}
func (t *service) DetailTrening(ctx context.Context, request *models.DetailTreningRequest) (response *models.DetailTreningResponse, err error) {
	return
}
func (t *service) CreateExercise(ctx context.Context, request *models.CreateExerciseRequest) (response *models.CreateExerciseResponse, err error) {
	return
}
func (t *service) ListExercise(ctx context.Context, request *models.ListExerciseRequest) (response *models.ListExerciseResponse, err error) {
	return
}
func (t *service) DetailExercise(ctx context.Context, request *models.DetailExerciseRequest) (response *models.DetailExerciseResponse, err error) {
	return
}

func NewTreningHandler() TreningHandler {
	return &service{}
}
