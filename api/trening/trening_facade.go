package grpcserver

import (
	"context"
	"log"

	"github.com/and07/boilerplate-go/internal/app/trening/models"
	"github.com/and07/boilerplate-go/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	log.Println("userID - ", userID)
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
	log.Println("userID - ", userID)
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
	log.Println("userID - ", userID)
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
	log.Println("userID - ", userID)
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
	log.Println("userID - ", userID)
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

func NewTeningFacade(extractor extractor, authService service.Authentication, treningHandler treningHandler, logger logger) *treningFacade {
	return &treningFacade{
		extractor:      extractor,
		authService:    authService,
		treningHandler: treningHandler,
		logger:         logger,
	}
}
