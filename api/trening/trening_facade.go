package grpcserver

import (
	"context"
	"log"
	"time"

	"github.com/and07/boilerplate-go/internal/app/trening/models"
	"github.com/and07/boilerplate-go/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
)

type treningFacade struct {
	extractor      extractor
	authService    service.Authentication
	treningHandler treningHandler
	logger         logger
}

func (t *treningFacade) CreateParametersUser(ctx context.Context, request *CreateParametersUserRequest) (response *CreateParametersUserResponse, err error) {

	userID, err := t.userID(ctx)
	if err != nil {
		response = &CreateParametersUserResponse{
			Status:  false,
			Message: err.Error(),
		}
		return
	}

	res, err := t.treningHandler.CreateParametersUser(ctx, &models.CreateParametersUserRequest{
		Weight:        request.Weight,
		Height:        request.Height,
		Age:           request.Age,
		Activity:      models.UserActivity(request.Activity),
		Diet:          models.UserDiet(request.Diet),
		Eat:           request.Eat,
		DesiredWeight: request.DesiredWeight,
		Gender:        request.Gender,
		UserID:        userID,
	})
	if err != nil {
		response = &CreateParametersUserResponse{
			Status:  false,
			Message: "can't create user parameters",
		}
		return
	}

	response = &CreateParametersUserResponse{
		Status:  res.Status,
		Message: res.Message,
	}

	return
}

func (t *treningFacade) DetailParametersUser(ctx context.Context, request *DetailParametersUserRequest) (response *DetailParametersUserResponse, err error) {
	userID, err := t.userID(ctx)
	if err != nil {
		response = &DetailParametersUserResponse{
			Status:  false,
			Message: err.Error(),
		}
		return
	}

	res, err := t.treningHandler.DetailParametersUser(ctx, &models.DetailParametersUserRequest{
		UserID: userID,
	})
	if err != nil {
		response = &DetailParametersUserResponse{
			Status:  false,
			Message: err.Error(),
		}
		return
	}
	data := ParametersUser{
		Weight:        res.Data.Weight,
		Height:        res.Data.Height,
		Age:           res.Data.Age,
		Gender:        res.Data.Gender,
		Activity:      UserActivity(res.Data.Activity),
		Diet:          UserDiet(res.Data.Diet),
		DesiredWeight: res.Data.DesiredWeight,
		Eat:           res.Data.Eat,
		Username:      res.Data.Username,
	}

	response = &DetailParametersUserResponse{
		Status: res.Status,
		Data:   &data,
	}

	t.logger.Debug("userParams %#v", res.Data)

	return
}

func (t *treningFacade) CreateTrening(ctx context.Context, request *CreateTreningRequest) (response *CreateTreningResponse, err error) {
	userID, err := t.userID(ctx)
	if err != nil {
		response = &CreateTreningResponse{
			Status:  false,
			Message: err.Error(),
		}
		return
	}

	var exercises []*models.Exercise

	for _, e := range request.Exercises {
		exercises = append(exercises, &models.Exercise{
			Name:                e.Name,
			Duration:            e.Duration.AsDuration(),
			Relax:               e.Relax.AsDuration(),
			Count:               e.Count,
			NumberOfSets:        e.NumberOfSets,
			NumberOfRepetitions: e.NumberOfRepetitions,
			Type:                models.ExerciseType(e.Type),
			Image:               e.Image,
			Video:               e.Video,
		})
	}

	// TODO
	res, err := t.treningHandler.CreateTrening(ctx, &models.CreateTreningRequest{
		UserID:    userID,
		Name:      request.Name,
		Interval:  request.Interval.AsDuration(),
		Exercises: exercises,
	})
	if err != nil {
		response = &CreateTreningResponse{
			Status:  false,
			Message: err.Error(),
		}
		return
	}
	response = &CreateTreningResponse{
		Status:  true,
		Message: res.Message,
	}

	return
}

func (t *treningFacade) ListTrening(ctx context.Context, request *ListTreningRequest) (response *ListTreningResponse, err error) {
	userID, err := t.userID(ctx)
	if err != nil {
		response = &ListTreningResponse{
			Status:  false,
			Message: err.Error(),
		}
		return
	}

	res, err := t.treningHandler.ListTrening(ctx, &models.ListTreningRequest{
		UserID: userID,
		Status: int(request.Status.Number()),
	})
	if err != nil {
		response = &ListTreningResponse{
			Status:  false,
			Message: err.Error(),
		}
		return
	}

	var trenings []*Trening
	for _, t := range res.Data {

		var exercises []*Exercise
		for _, e := range t.Exercises {
			exercises = append(exercises, &Exercise{
				Name:                e.Name,
				Duration:            durationpb.New(e.Duration * time.Second),
				Relax:               durationpb.New(e.Relax * time.Second),
				Count:               e.Count,
				NumberOfSets:        e.NumberOfSets,
				NumberOfRepetitions: e.NumberOfRepetitions,
				Type:                ExerciseType(e.Type),
				Video:               e.Video,
				Image:               e.Image,
			})
		}

		trenings = append(trenings, &Trening{
			Uid:       t.UID,
			Name:      t.Name,
			Interval:  durationpb.New(t.Interval * time.Second),
			Exercises: exercises,
			Image:     t.Image,
		})
	}

	response = &ListTreningResponse{
		Status: res.Status,
		Data:   trenings,
	}

	return
}

func (t *treningFacade) DetailTrening(ctx context.Context, request *DetailTreningRequest) (response *DetailTreningResponse, err error) {
	userID, err := t.userID(ctx)
	if err != nil {
		response = &DetailTreningResponse{
			Status:  false,
			Message: err.Error(),
		}
		return
	}
	res, err := t.treningHandler.DetailTrening(ctx, &models.DetailTreningRequest{
		UserID: userID,
		UID:    request.Uid,
	})
	if err != nil {
		response = &DetailTreningResponse{
			Status:  false,
			Message: err.Error(),
		}
		return
	}

	var exercises []*Exercise
	for _, e := range res.Data.Exercises {
		exercises = append(exercises, &Exercise{
			Name:                e.Name,
			Duration:            durationpb.New(e.Duration * time.Second),
			Relax:               durationpb.New(e.Relax * time.Second),
			Count:               e.Count,
			NumberOfSets:        e.NumberOfSets,
			NumberOfRepetitions: e.NumberOfRepetitions,
			Type:                ExerciseType(e.Type),
			Video:               e.Video,
			Image:               e.Image,
		})
	}

	response = &DetailTreningResponse{
		Status: res.Status,
		Data: &Trening{
			Uid:       res.Data.UID,
			Name:      res.Data.Name,
			Interval:  durationpb.New(res.Data.Interval * time.Second),
			Exercises: exercises,
			Image:     res.Data.Image,
		},
	}

	return
}

func (t *treningFacade) UpdateTreningStatus(ctx context.Context, request *UpdateTreningStatusRequest) (response *UpdateTreningStatusResponse, err error) {
	//TODO
	userID, err := t.userID(ctx)
	if err != nil {
		response = &UpdateTreningStatusResponse{
			Status:  false,
			Message: err.Error(),
		}
		return
	}
	t.logger.Debug("request", request.Status)

	res, err := t.treningHandler.UpdateTreningStatus(ctx, &models.UpdateTreningStatusRequest{
		UserID: userID,
		Status: int(request.Status.Number()),
		UID:    request.Uid,
	})
	if err != nil {
		response = &UpdateTreningStatusResponse{
			Status:  false,
			Message: err.Error(),
		}
		return
	}

	response = &UpdateTreningStatusResponse{
		Status: res.Status,
	}
	return
}

func (t *treningFacade) CreateExercise(ctx context.Context, request *CreateExerciseRequest) (response *CreateExerciseResponse, err error) {
	userID, err := t.userID(ctx)
	if err != nil {
		response = &CreateExerciseResponse{
			Status:  false,
			Message: err.Error(),
		}
		return
	}

	res, err := t.treningHandler.CreateExercise(ctx, &models.CreateExerciseRequest{
		UserID:              userID,
		Name:                request.Name,
		Duration:            request.Duration.AsDuration(),
		Relax:               request.Relax.AsDuration(),
		Count:               request.Count,
		NumberOfSets:        request.NumberOfSets,
		NumberOfRepetitions: request.NumberOfRepetitions,
		Type:                models.ExerciseType(request.Type),
	})
	if err != nil {
		return
	}

	response = &CreateExerciseResponse{
		Status:  res.Status,
		Message: res.Message,
	}

	return
}

func (t *treningFacade) ListExercise(ctx context.Context, request *ListExerciseRequest) (response *ListExerciseResponse, err error) {
	userID, err := t.userID(ctx)
	if err != nil {
		response = &ListExerciseResponse{
			Status:  false,
			Message: err.Error(),
		}
		return
	}

	res, err := t.treningHandler.ListExercise(ctx, &models.ListExerciseRequest{
		UserID: userID,
	})
	if err != nil {
		return
	}
	response = &ListExerciseResponse{}
	response.Status = res.Status
	response.Data = make([]*Exercise, len(res.Data))

	for i, e := range res.Data {
		response.Data[i] = &Exercise{
			Uid:                 e.UID,
			Name:                e.Name,
			Duration:            durationpb.New(e.Duration),
			Relax:               durationpb.New(e.Relax),
			Count:               e.Count,
			NumberOfSets:        e.NumberOfSets,
			NumberOfRepetitions: e.NumberOfRepetitions,
			Type:                ExerciseType(e.Type),
		}
	}

	return
}

func (t *treningFacade) DetailExercise(ctx context.Context, request *DetailExerciseRequest) (response *DetailExerciseResponse, err error) {
	userID, err := t.userID(ctx)
	if err != nil {
		response = &DetailExerciseResponse{
			Status:  false,
			Message: err.Error(),
		}
		return
	}
	log.Println("userID - ", userID)
	return
}

func (t *treningFacade) ListExerciseDefault(ctx context.Context, request *ListExerciseRequest) (response *ListExerciseResponse, err error) {
	_, err = t.userID(ctx)
	if err != nil {
		response = &ListExerciseResponse{
			Status:  false,
			Message: err.Error(),
		}
		return
	}

	exercises := exerciseDefault

	if request.Type != "" {
		exercises = []*Exercise{}
		for _, e := range exerciseDefault {
			if request.Type == e.Type.String() {
				exercises = append(exercises, e)
			}
		}
	}

	response = &ListExerciseResponse{
		Status: true,
		Data:   exercises,
	}

	return
}

func (t *treningFacade) userID(ctx context.Context) (string, error) {
	token, isExist := t.extractor.ExtractGRPC(ctx)
	if !isExist {
		return "", status.Error(codes.Unauthenticated, "token is not exist")
	}

	userID, err := t.authService.ValidateAccessToken(token)
	if err != nil {
		return "", status.Error(codes.Unauthenticated, "Authentication failed. Invalid token")
	}
	return userID, nil
}

// NewTeningFacade ...
func NewTeningFacade(extractor extractor, authService service.Authentication, treningHandler treningHandler, logger logger) *treningFacade {
	return &treningFacade{
		extractor:      extractor,
		authService:    authService,
		treningHandler: treningHandler,
		logger:         logger,
	}
}
