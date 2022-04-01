package grpcserver

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type treningFacade struct {
	extractor extractor
}

func (t *treningFacade) CreateParametersUser(ctx context.Context, request *CreateParametersUserRequest) (response *CreateParametersUserResponse, err error) {
	token, isExist := t.extractor.ExtractGRPC(ctx)
	if !isExist {
		return nil, status.Error(codes.Unauthenticated, "token is not exist")
	}

	log.Println("treningFacade -", token)

	response = &CreateParametersUserResponse{
		Status:  false,
		Message: "treningFacade",
	}

	return
}
func (t *treningFacade) DetailParametersUser(ctx context.Context, request *DetailParametersUserRequest) (response *DetailParametersUserResponse, err error) {
	return
}
func (t *treningFacade) CreateTrening(ctx context.Context, request *CreateTreningRequest) (response *CreateTreningResponse, err error) {
	return
}
func (t *treningFacade) ListTrening(ctx context.Context, request *ListTreningRequest) (response *ListTreningResponse, err error) {
	return
}
func (t *treningFacade) DetailTrening(ctx context.Context, request *DetailTreningRequest) (response *DetailTreningResponse, err error) {
	return
}
func (t *treningFacade) CreateExercise(ctx context.Context, request *CreateExerciseRequest) (response *CreateExerciseResponse, err error) {
	return
}
func (t *treningFacade) ListExercise(ctx context.Context, request *ListExerciseRequest) (response *ListExerciseResponse, err error) {
	return
}
func (t *treningFacade) DetailExercise(ctx context.Context, request *DetailExerciseRequest) (response *DetailExerciseResponse, err error) {
	return
}

func NewtTeningFacade(extractor extractor) *treningFacade {
	return &treningFacade{
		extractor: extractor,
	}
}
