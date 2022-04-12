package handlers

import (
	"context"
	"time"

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
	UpdateTreningStatus(ctx context.Context, request *models.UpdateTreningStatusRequest) (response *models.UpdateTreningStatusResponse, err error)

	CreateExercise(ctx context.Context, request *models.CreateExerciseRequest) (response *models.CreateExerciseResponse, err error)
	ListExercise(ctx context.Context, request *models.ListExerciseRequest) (response *models.ListExerciseResponse, err error)
	DetailExercise(ctx context.Context, request *models.DetailExerciseRequest) (response *models.DetailExerciseResponse, err error)
}

type service struct {
	repo   data.TreningRepository
	logger logger
}

func (s *service) CreateParametersUser(ctx context.Context, request *models.CreateParametersUserRequest) (response *models.CreateParametersUserResponse, err error) {

	usr, err := s.repo.GetUserByID(ctx, request.UserID)
	if err != nil {
		s.logger.Error("unable to insert user to database", "error", err)
		return nil, err
	}

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
		UserName:      usr.Username,
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
			Username:      userParams.UserName,
		},
	}

	return
}

func (s *service) CreateTrening(ctx context.Context, request *models.CreateTreningRequest) (response *models.CreateTreningResponse, err error) {

	var exercises data.ExerciseSlice

	for _, e := range request.Exercises {
		exercises = append(exercises, data.Exercise{
			UserID:              request.UserID,
			Name:                e.Name,
			Duration:            e.Duration / time.Second,
			Relax:               e.Relax / time.Second,
			Count:               e.Count,
			NumberOfSets:        e.NumberOfSets,
			NumberOfRepetitions: e.NumberOfRepetitions,
			Type:                int32(e.Type),
			Image:               e.Image,
			Video:               e.Video,
		})
	}

	err = s.repo.CreateTrening(ctx, &data.Trening{
		Name:      request.Name,
		UserID:    request.UserID,
		Interval:  request.Interval / time.Second,
		Exercises: exercises,
	})
	if err != nil {
		return nil, err
	}

	response = &models.CreateTreningResponse{
		Status: true,
	}

	return
}
func (s *service) UpdateTrening(ctx context.Context, request *models.UpdateTreningRequest) (response *models.UpdateTreningResponse, err error) {
	return
}
func (s *service) ListTrening(ctx context.Context, request *models.ListTreningRequest) (response *models.ListTreningResponse, err error) {

	res, err := s.repo.ListTrening(ctx, request.UserID, request.Status)
	if err != nil {
		return nil, err
	}
	var trenings []*models.Trening
	for _, t := range res {

		var exercises []*models.Exercise
		for _, e := range t.Exercises {
			exercises = append(exercises, &models.Exercise{
				Name:                e.Name,
				Duration:            e.Duration,
				Relax:               e.Relax,
				Count:               e.Count,
				NumberOfSets:        e.NumberOfSets,
				NumberOfRepetitions: e.NumberOfRepetitions,
				Type:                models.ExerciseType(e.Type),
				Image:               e.Image,
				Video:               e.Video,
			})
		}

		trenings = append(trenings, &models.Trening{
			UID:       t.UID,
			Name:      t.Name,
			Interval:  t.Interval,
			Exercises: exercises,
		})
	}

	response = &models.ListTreningResponse{
		Status: true,
		Data:   trenings,
	}

	return
}

func (s *service) UpdateTreningStatus(ctx context.Context, request *models.UpdateTreningStatusRequest) (response *models.UpdateTreningStatusResponse, err error) {

	if err = s.repo.UpdateTreningStatus(ctx, &data.Trening{UID: request.UID, UserID: request.UserID, Status: request.Status}); err != nil {
		s.logger.Error("unable to update trening status to database", "error", err)
		return nil, err
	}

	s.logger.Debug("request", request)

	response = &models.UpdateTreningStatusResponse{
		Status: true,
	}

	return
}

func (s *service) DetailTrening(ctx context.Context, request *models.DetailTreningRequest) (response *models.DetailTreningResponse, err error) {

	res, err := s.repo.DetailTrening(ctx, request.UserID, request.UID)
	if err != nil {
		s.logger.Error("unable to get trening to database", "error", err)
		return nil, err
	}

	var exercises []*models.Exercise
	for _, e := range res.Exercises {
		exercises = append(exercises, &models.Exercise{
			Name:                e.Name,
			Duration:            e.Duration,
			Relax:               e.Relax,
			Count:               e.Count,
			NumberOfSets:        e.NumberOfSets,
			NumberOfRepetitions: e.NumberOfRepetitions,
			Type:                models.ExerciseType(e.Type),
			Image:               e.Image,
			Video:               e.Video,
		})
	}

	response = &models.DetailTreningResponse{
		Status: true,
		Data: &models.Trening{
			UID:       res.UID,
			Name:      res.Name,
			Interval:  res.Interval,
			Exercises: exercises,
		},
	}

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
	response = &models.ListExerciseResponse{}

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
