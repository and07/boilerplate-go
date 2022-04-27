package grpcserver

import (
	"context"
	"net/http"

	"github.com/and07/boilerplate-go/internal/app/trening/models"
)

type extractor interface {
	ExtractGRPC(ctx context.Context) (header string, existStatus bool)
	ExtractHTTP(r *http.Request) (header string, existStatus bool)
}
type treningHandler interface {
	CreateParametersUser(ctx context.Context, request *models.CreateParametersUserRequest) (response *models.CreateParametersUserResponse, err error)
	DetailParametersUser(ctx context.Context, request *models.DetailParametersUserRequest) (response *models.DetailParametersUserResponse, err error)
	UpdateUserParams(ctx context.Context, request *models.UpdateUserParamsRequest) (response *models.UpdateUserParamsResponse, err error)
	UpdateUserImage(ctx context.Context, request *models.UpdateUserImageRequest) (response *models.UpdateUserImageResponse, err error)

	CreateTrening(ctx context.Context, request *models.CreateTreningRequest) (response *models.CreateTreningResponse, err error)
	ListTrening(ctx context.Context, request *models.ListTreningRequest) (response *models.ListTreningResponse, err error)
	DetailTrening(ctx context.Context, request *models.DetailTreningRequest) (response *models.DetailTreningResponse, err error)
	UpdateTreningStatus(ctx context.Context, request *models.UpdateTreningStatusRequest) (response *models.UpdateTreningStatusResponse, err error)
	UpdateTreningExercises(ctx context.Context, request *models.UpdateTreningExercisesRequest) (response *models.UpdateTreningExercisesResponse, err error)

	CreateExercise(ctx context.Context, request *models.CreateExerciseRequest) (response *models.CreateExerciseResponse, err error)
	ListExercise(ctx context.Context, request *models.ListExerciseRequest) (response *models.ListExerciseResponse, err error)
	DetailExercise(ctx context.Context, request *models.DetailExerciseRequest) (response *models.DetailExerciseResponse, err error)
}

type logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}
