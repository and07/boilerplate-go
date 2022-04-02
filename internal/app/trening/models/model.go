package models

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

type UserDiet int32
type UserActivity int32
type CreateParametersUserRequest struct {
	Weight        int32        `protobuf:"varint,1,opt,name=weight,proto3" json:"weight,omitempty"`
	Height        int32        `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Age           int32        `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Gender        int32        `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Activity      UserActivity `protobuf:"varint,5,opt,name=activity,proto3,enum=trening.UserActivity" json:"activity,omitempty"`
	Diet          UserDiet     `protobuf:"varint,6,opt,name=diet,proto3,enum=trening.UserDiet" json:"diet,omitempty"`
	DesiredWeight int32        `protobuf:"varint,7,opt,name=desired_weight,json=desiredWeight,proto3" json:"desired_weight,omitempty"`
	Eat           int32        `protobuf:"varint,8,opt,name=eat,proto3" json:"eat,omitempty"`
	UserID        string
}

type CreateParametersUserResponse struct {
	Status  bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

type DetailParametersUserRequest struct {
	ID     string
	UserID string
}
type ParametersUser struct {
	Weight        int32        `protobuf:"varint,1,opt,name=weight,proto3" json:"weight,omitempty"`
	Height        int32        `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Age           int32        `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Gender        int32        `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Activity      UserActivity `protobuf:"varint,5,opt,name=activity,proto3,enum=trening.UserActivity" json:"activity,omitempty"`
	Diet          UserDiet     `protobuf:"varint,6,opt,name=diet,proto3,enum=trening.UserDiet" json:"diet,omitempty"`
	DesiredWeight int32        `protobuf:"varint,7,opt,name=desired_weight,json=desiredWeight,proto3" json:"desired_weight,omitempty"`
	Eat           int32        `protobuf:"varint,8,opt,name=eat,proto3" json:"eat,omitempty"`
}

type DetailParametersUserResponse struct {
	Status  bool            `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *ParametersUser `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

type Exercise struct {
	Name     string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Duration time.Time `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Relax    time.Time `protobuf:"bytes,3,opt,name=relax,proto3" json:"relax,omitempty"`
	Count    int32     `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}
type CreateTreningRequest struct {
	Name      string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Interval  time.Time   `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	Exercises []*Exercise `protobuf:"bytes,3,rep,name=exercises,proto3" json:"exercises,omitempty"`
	UserID    string
}
type CreateTreningResponse struct {
	Status  bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}
type ListTreningRequest struct {
	UserID string
}
type Trening struct {
	Name      string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Interval  time.Time   `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	Exercises []*Exercise `protobuf:"bytes,3,rep,name=exercises,proto3" json:"exercises,omitempty"`
}
type ListTreningResponse struct {
	Status  bool       `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*Trening `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}
type DetailTreningRequest struct {
	ID     string
	UserID string
}
type DetailTreningResponse struct {
	Status  bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *Trening `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}
type CreateExerciseRequest struct {
	Name     string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Duration *timestamp.Timestamp `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Relax    *timestamp.Timestamp `protobuf:"bytes,3,opt,name=relax,proto3" json:"relax,omitempty"`
	Count    int32                `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	UserID   string
}
type CreateExerciseResponse struct {
	Status  bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}
type ListExerciseRequest struct {
	UserID string
}
type ListExerciseResponse struct {
	Status  bool        `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*Exercise `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	UserID  string
}
type DetailExerciseRequest struct {
	ID     string
	UserID string
}
type DetailExerciseResponse struct {
	Status  bool      `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *Exercise `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}