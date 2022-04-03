package handlers

import (
	"context"

	"github.com/and07/boilerplate-go/internal/app/trening/models"
	"github.com/and07/boilerplate-go/pkg/data"
)

// TreningHandler ...
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

type service struct {
	repo   data.TreningRepository
	logger logger
}

func (s *service) CreateParametersUser(ctx context.Context, request *models.CreateParametersUserRequest) (response *models.CreateParametersUserResponse, err error) {

	if err = s.repo.CreateUserParams(ctx, &data.ParametersUser{
		Weight:        request.Weight,
		Height:        request.Height,
		Age:           request.Age,
		Gender:        request.Gender,
		Activity:      int32(request.Activity),
		Diet:          int32(request.Diet),
		DesiredWeight: request.DesiredWeight,
		Eat:           request.Eat,
		UserID:        request.UserID,
	}); err != nil {

		s.logger.Error("unable to insert user params to database", "error", err)
		return nil, err
	}
	response = &models.CreateParametersUserResponse{
		Status: true,
	}

	return
}

func (s *service) DetailParametersUser(ctx context.Context, request *models.DetailParametersUserRequest) (response *models.DetailParametersUserResponse, err error) {

	userParams, err := s.repo.GetUserParamsByID(ctx, request.UserID)
	if err != nil {
		return nil, err
	}

	s.logger.Debug("userParams %#v", userParams)

	response = &models.DetailParametersUserResponse{
		Status: true,
		Data: &models.ParametersUser{
			Weight:        userParams.Weight,
			Height:        userParams.Height,
			Age:           userParams.Age,
			Gender:        userParams.Gender,
			Activity:      models.UserActivity(userParams.Activity),
			Diet:          models.UserDiet(userParams.Diet),
			DesiredWeight: userParams.DesiredWeight,
			Eat:           userParams.Eat,
		},
	}

	return
}

func (s *service) CreateTrening(ctx context.Context, request *models.CreateTreningRequest) (response *models.CreateTreningResponse, err error) {
	return
}
func (s *service) ListTrening(ctx context.Context, request *models.ListTreningRequest) (response *models.ListTreningResponse, err error) {
	return
}
func (s *service) DetailTrening(ctx context.Context, request *models.DetailTreningRequest) (response *models.DetailTreningResponse, err error) {
	return
}

func (s *service) CreateExercise(ctx context.Context, request *models.CreateExerciseRequest) (response *models.CreateExerciseResponse, err error) {

	if err = s.repo.CreateExercise(ctx, &data.Exercise{
		UserID:              request.UserID,
		Name:                request.Name,
		Duration:            request.Duration,
		Relax:               request.Duration,
		Count:               request.Count,
		NumberOfSets:        request.NumberOfSets,
		NumberOfRepetitions: request.NumberOfRepetitions,
		Type:                int32(request.Type),
	}); err != nil {

		s.logger.Error("unable to insert exercise to database", "error", err)
		return nil, err
	}

	response = &models.CreateExerciseResponse{
		Status: true,
	}
	return
}

// ListExercise ...
func (s *service) ListExercise(ctx context.Context, request *models.ListExerciseRequest) (response *models.ListExerciseResponse, err error) {

	res, err := s.repo.ListExercise(ctx, request.UserID)
	if err != nil {
		s.logger.Error("unable to get exercise to database", "error", err)
		return nil, err
	}
	response.Data = make([]*models.Exercise, len(res))
	for i, e := range res {
		response.Data[i] = &models.Exercise{
			UID:                 e.UID,
			Name:                e.Name,
			Duration:            e.Duration,
			Relax:               e.Relax,
			Count:               e.Count,
			NumberOfSets:        e.NumberOfSets,
			NumberOfRepetitions: e.NumberOfRepetitions,
			Type:                models.ExerciseType(e.Type),
		}
	}
	response.Status = true

	return
}

func (s *service) DetailExercise(ctx context.Context, request *models.DetailExerciseRequest) (response *models.DetailExerciseResponse, err error) {
	return
}

// NewTreningHandler ...
func NewTreningHandler(repo data.TreningRepository, logger logger) TreningHandler {
	return &service{
		repo:   repo,
		logger: logger,
	}
}
