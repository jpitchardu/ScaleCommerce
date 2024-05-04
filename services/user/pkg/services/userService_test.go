package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	type args struct {
		model *UserModel
	}

	testCases := []struct {
		name string
		args args
		id   string
		err  error
	}{
		{
			name: "should create user",
			args: args{
				model: &UserModel{
					ID:       "1",
					Name:     "name",
					Email:    "email",
					Password: "password",
				},
			},
			id:  "1",
			err: nil,
		},
	}

	s := NewUserService()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			id, err := s.CreateUser(tc.args.model)

			assert.Equal(t, id, tc.id)
			assert.Equal(t, err, tc.err)

		})
	}
}
