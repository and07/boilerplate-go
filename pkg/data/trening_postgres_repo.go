package data

import (
	"context"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// TreningRepository ...
type TreningRepository interface {
	CreateUserParams(ctx context.Context, userParams *ParametersUser) error
	GetUserParamsByID(ctx context.Context, userID string) (*ParametersUser, error)
	CreateExercise(ctx context.Context, exercise *Exercise) error
	ListExercise(ctx context.Context, userID string) (res []Exercise, err error)
}

type treningPostgresRepository struct {
	db     *sqlx.DB
	logger hclog.Logger
}

// NewTreningPostgresRepository returns a new TreningPostgresRepository instance
func NewTreningPostgresRepository(db *sqlx.DB, logger hclog.Logger) TreningRepository {

	// creation of trening table.
	if db != nil {
		db.MustExec(treningUserParamsSchema)
		db.MustExec(treningExercise)
		//db.MustExec(trening)
	}

	return &treningPostgresRepository{db, logger}
}

// CreateCreateUserParams inserts the given user params into the database
func (repo *treningPostgresRepository) CreateUserParams(ctx context.Context, userParams *ParametersUser) error {
	userParams.UID = uuid.NewV4().String()
	userParams.CreatedAt = time.Now()
	userParams.UpdatedAt = time.Now()

	repo.logger.Info("creating user params", hclog.Fmt("%#v", userParams))
	query := `insert into trening_users_params 
		(uid, user_id, weight, height, age, gender, activity, diet, desired_weight, eat, createdat, updatedat) 
	values 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);
	`
	_, err := repo.db.ExecContext(ctx, query, userParams.UID, userParams.UserID, userParams.Weight, userParams.Height, userParams.Age, userParams.Gender, userParams.Activity, userParams.Diet, userParams.DesiredWeight, userParams.Eat, userParams.CreatedAt, userParams.UpdatedAt)
	return err
}

// GetUserParamsByID retrieves the user object having the given user_id, else returns error
func (repo *treningPostgresRepository) GetUserParamsByID(ctx context.Context, userID string) (*ParametersUser, error) {
	repo.logger.Debug("querying for user params with user_id", userID)
	query := "select * from trening_users_params where user_id = $1"
	var userParams ParametersUser
	if err := repo.db.GetContext(ctx, &userParams, query, userID); err != nil {
		return nil, err
	}
	repo.logger.Debug("read users params", hclog.Fmt("%#v", userParams))
	return &userParams, nil
}

// CreateExercise inserts the given exercise into the database
func (repo *treningPostgresRepository) CreateExercise(ctx context.Context, exercise *Exercise) error {
	exercise.UID = uuid.NewV4().String()
	exercise.CreatedAt = time.Now()
	exercise.UpdatedAt = time.Now()

	repo.logger.Info("creating exercise", hclog.Fmt("%#v", exercise))
	query := `insert into trening_users_params 
		(uid, user_id, name, duration, relax, count, number_of_sets, number_of_repetitions, type, createdat, updatedat) 
	values 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
	`
	_, err := repo.db.ExecContext(ctx, query, exercise.UID, exercise.UserID, exercise.Name, exercise.Duration, exercise.Relax, exercise.Count, exercise.NumberOfSets, exercise.NumberOfRepetitions, exercise.Type, exercise.CreatedAt, exercise.UpdatedAt)
	return err
}

func (repo *treningPostgresRepository) ListExercise(ctx context.Context, userID string) (res []Exercise, err error) {
	repo.logger.Info("list exercise")
	query := `select * from users where user_id = $1`
	res = []Exercise{}
	err = repo.db.Select(&res, query, userID)
	if err != nil {
		return
	}

	return
}
