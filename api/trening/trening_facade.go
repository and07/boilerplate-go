package grpcserver

import (
	"context"
	"encoding/json"
	"errors"
	fmt "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/and07/boilerplate-go/internal/app/trening/models"
	"github.com/and07/boilerplate-go/pkg/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type treningFacade struct {
	extractor      extractor
	jwtManager     *token.JWTManager
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

		if err.Error() == "sql: no rows in result set" {
			return nil, status.Error(codes.NotFound, "NotFound")
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
		Uid:           res.Data.UID,
		Image:         res.Data.Image,
	}

	response = &DetailParametersUserResponse{
		Status: res.Status,
		Data:   &data,
	}

	t.logger.Debug("userParams %#v", res.Data)

	return
}

func (t *treningFacade) UpdateUserParams(ctx context.Context, request *UpdateUserParamsRequest) (response *UpdateUserParamsResponse, err error) {
	userID, err := t.userID(ctx)
	if err != nil {
		response = &UpdateUserParamsResponse{
			Status:  false,
			Message: err.Error(),
		}
		return
	}

	res, err := t.treningHandler.UpdateUserParams(ctx, &models.UpdateUserParamsRequest{
		Weight:        request.Weight,
		Height:        request.Height,
		Age:           request.Age,
		Activity:      models.UserActivity(request.Activity),
		Diet:          models.UserDiet(request.Diet),
		Eat:           request.Eat,
		DesiredWeight: request.DesiredWeight,
		Gender:        request.Gender,
		UserID:        userID,
		UID:           request.Uid,
	})
	if err != nil {
		response = &UpdateUserParamsResponse{
			Status:  false,
			Message: "can't create user parameters",
		}
		return
	}

	response = &UpdateUserParamsResponse{
		Status:  res.Status,
		Message: res.Message,
	}

	return
}

func (t *treningFacade) UploadImageUser(u TreningService_UploadImageUserServer) error {
	//TODO

	return errors.New("not implemented")
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
			Description:         e.Description,
			Technique:           e.Technique,
			Mistake:             e.Mistake,
			Weight:              e.Weight,
		})
	}

	// TODO
	res, err := t.treningHandler.CreateTrening(ctx, &models.CreateTreningRequest{
		UserID:    userID,
		Name:      request.Name,
		Interval:  request.Interval.AsDuration(),
		Exercises: exercises,
		Date:      request.Date.AsTime(),
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
				Description:         e.Description,
				Technique:           e.Technique,
				Mistake:             e.Mistake,
				Weight:              e.Weight,
			})
		}

		trenings = append(trenings, &Trening{
			Uid:       t.UID,
			Name:      t.Name,
			Interval:  durationpb.New(t.Interval * time.Second),
			Exercises: exercises,
			Image:     t.Image,
			Date:      timestamppb.New(t.Date),
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
			Description:         e.Description,
			Technique:           e.Technique,
			Mistake:             e.Mistake,
			Weight:              e.Weight,
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
			Date:      timestamppb.New(res.Data.Date),
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
			Weight:              e.Weight,
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

func (t *treningFacade) UpdateTreningExercises(ctx context.Context, request *UpdateTreningExercisesRequest) (response *UpdateTreningExercisesResponse, err error) {
	userID, err := t.userID(ctx)
	if err != nil {
		response = &UpdateTreningExercisesResponse{
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
			Description:         e.Description,
			Technique:           e.Technique,
			Mistake:             e.Mistake,
			Weight:              e.Weight,
		})
	}

	res, err := t.treningHandler.UpdateTreningExercises(ctx, &models.UpdateTreningExercisesRequest{
		UserID:    userID,
		Exercises: exercises,
		UID:       request.Uid,
	})
	if err != nil {
		response = &UpdateTreningExercisesResponse{
			Status:  false,
			Message: err.Error(),
		}
		return
	}

	response = &UpdateTreningExercisesResponse{
		Status: res.Status,
	}
	return
}

func (t *treningFacade) BinaryFileUpload(w http.ResponseWriter, r *http.Request, params map[string]string) {

	token, err := t.extractor.ExtractHTTP(r)
	if err != nil {
		t.logger.Debug("token not exist %#v", err)
		http.Error(w, "token not exist", http.StatusUnauthorized)
		return
	}

	claims, err := t.jwtManager.ValidateAccessToken(token)
	if err != nil {
		http.Error(w, "Authentication failed. Invalid token", http.StatusUnauthorized)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	f, header, err := r.FormFile("attachment")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get file 'attachment': %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer f.Close()
	filename := header.Filename

	// TODO
	// Now do something with the io.Reader in `f`, i.e. read it into a buffer or stream it to a gRPC client side stream.
	// Also `header` will contain the filename, size etc of the original file.
	//

	b, err := ioutil.ReadAll(f)
	if err != nil {
		//TODO
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	t.logger.Debug("filename %s   header %#v ", filename, header.Header)
	res, err := t.treningHandler.UpdateUserImage(context.Background(), &models.UpdateUserImageRequest{
		UserID:   claims.UserID,
		Body:     b,
		FileName: filename,
	})
	if err != nil {
		//TODO
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)

}

func (t *treningFacade) userID(ctx context.Context) (string, error) {
	claims, ok := ctx.Value("claim").(*token.AccessTokenCustomClaims)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "Authentication failed. Invalid token")
	}
	return claims.UserID, nil
}

// NewTeningFacade ...
func NewTeningFacade(extractor extractor, jwtManager *token.JWTManager, treningHandler treningHandler, logger logger) *treningFacade {
	return &treningFacade{
		extractor:      extractor,
		jwtManager:     jwtManager,
		treningHandler: treningHandler,
		logger:         logger,
	}
}
