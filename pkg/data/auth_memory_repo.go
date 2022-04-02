package data

import (
	"context"
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/hashicorp/go-hclog"
	uuid "github.com/satori/go.uuid"
)

// MemoryRepository has the implementation of the memory methods.
type MemoryRepository struct {
	logger           hclog.Logger
	userData         map[string]*User
	verificationData map[string]*VerificationData
	mx               sync.RWMutex
}

// NewAuthMemoryRepository returns a new PostgresRepository instance
func NewAuthMemoryRepository(logger hclog.Logger) *MemoryRepository {
	return &MemoryRepository{
		logger:           logger,
		userData:         make(map[string]*User),
		verificationData: make(map[string]*VerificationData),
	}
}

// Create inserts the given user into the database
func (repo *MemoryRepository) Create(ctx context.Context, user *User) error {
	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	repo.logger.Info("creating user", hclog.Fmt("%#v", user))

	repo.mx.Lock()
	repo.userData[user.Email] = user
	repo.mx.Unlock()

	return nil
}

// GetUserByEmail retrieves the user object having the given email, else returns error
func (repo *MemoryRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	repo.logger.Debug("querying for user with email", email)

	repo.mx.RLock()
	defer repo.mx.RUnlock()
	user, ok := repo.userData[email]
	if !ok {
		return nil, errors.New("no rows in result set")
	}

	repo.logger.Debug("read users", hclog.Fmt("%#v", user))

	return user, nil
}

// GetUserByID retrieves the user object having the given ID, else returns error
func (repo *MemoryRepository) GetUserByID(ctx context.Context, userID string) (*User, error) {
	repo.logger.Debug("querying for user with id", userID)

	repo.mx.RLock()
	defer repo.mx.RUnlock()
	for _, usr := range repo.userData {
		if usr.ID == userID {
			return usr, nil
		}
	}

	return nil, errors.New("not found user")
}

// UpdateUsername updates the username of the given user
func (repo *MemoryRepository) UpdateUsername(ctx context.Context, user *User) error {
	user.UpdatedAt = time.Now()

	repo.mx.Lock()
	repo.userData[user.Email] = user
	repo.mx.Unlock()

	return nil
}

// UpdateUserVerificationStatus updates user verification status to true
func (repo *MemoryRepository) UpdateUserVerificationStatus(ctx context.Context, email string, status bool) error {

	repo.mx.Lock()
	user := repo.userData[email]
	user.IsVerified = status

	repo.userData[user.Email] = user
	repo.mx.Unlock()

	return nil
}

// UpdatePassword updates the user password
func (repo *MemoryRepository) UpdatePassword(ctx context.Context, userID string, password string, tokenHash string) error {

	repo.mx.Lock()
	defer repo.mx.Unlock()
	var user *User
	for _, usr := range repo.userData {
		if usr.ID == userID {
			user = usr
			break
		}
	}

	user.Password = password
	user.TokenHash = tokenHash

	return nil
}

// StoreVerificationData adds a mail verification data to db
func (repo *MemoryRepository) StoreVerificationData(ctx context.Context, verificationData *VerificationData) error {

	repo.logger.Debug("verificationData ", hclog.Fmt("%#v", verificationData))

	repo.mx.Lock()
	repo.verificationData[verificationData.Email+strconv.Itoa(int(verificationData.Type))] = verificationData
	repo.mx.Unlock()

	return nil
}

// GetVerificationData retrieves the stored verification code.
func (repo *MemoryRepository) GetVerificationData(ctx context.Context, email string, verificationDataType VerificationDataType) (*VerificationData, error) {
	repo.mx.RLock()
	verificationData := repo.verificationData[email+strconv.Itoa(int(verificationDataType))]
	repo.mx.RUnlock()

	return verificationData, nil
}

// DeleteVerificationData deletes a used verification data
func (repo *MemoryRepository) DeleteVerificationData(ctx context.Context, email string, verificationDataType VerificationDataType) error {
	repo.mx.Lock()
	delete(repo.verificationData, email+strconv.Itoa(int(verificationDataType)))
	repo.mx.Unlock()

	return nil
}
