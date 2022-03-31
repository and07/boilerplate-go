// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/trening/trening.proto

package grpcserver

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ParametersUserRequest) Validate() error {
	return nil
}
func (this *ParametersUserResponse) Validate() error {
	return nil
}
func (this *DetailParametersUserRequest) Validate() error {
	return nil
}
func (this *DetailParametersUserResponse) Validate() error {
	return nil
}

var _regex_Exercise_Name = regexp.MustCompile(`^[a-zA-Z0-9_.-]*$`)

func (this *Exercise) Validate() error {
	if !_regex_Exercise_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9_.-]*$"`, this.Name))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Duration != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Duration); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Duration", err)
		}
	}
	if this.Relax != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Relax); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Relax", err)
		}
	}
	return nil
}
func (this *CreateTreningRequest) Validate() error {
	if this.Interval != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Interval); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Interval", err)
		}
	}
	for _, item := range this.Exercises {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Exercises", err)
			}
		}
	}
	return nil
}
func (this *CreateTreningResponse) Validate() error {
	return nil
}
func (this *ListTreningRequest) Validate() error {
	return nil
}
func (this *ListTreningResponse) Validate() error {
	return nil
}
func (this *DetailTreningRequest) Validate() error {
	return nil
}
func (this *DetailTreningResponse) Validate() error {
	return nil
}
func (this *CreateExerciseRequest) Validate() error {
	if this.Duration != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Duration); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Duration", err)
		}
	}
	if this.Relax != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Relax); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Relax", err)
		}
	}
	return nil
}
func (this *CreateExerciseResponse) Validate() error {
	return nil
}
func (this *ListExerciseRequest) Validate() error {
	return nil
}
func (this *ListExerciseResponse) Validate() error {
	return nil
}
func (this *DetailExerciseRequest) Validate() error {
	return nil
}
func (this *DetailExerciseResponse) Validate() error {
	return nil
}
