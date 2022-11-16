package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID      string `json:"id" validate:"len:36"`
		Name    string
		Age     int      `validate:"min:18|max:50"`
		Email   string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role    UserRole `validate:"in:admin,stuff"`
		RoleInt int      `validate:"in:200,100"`
		Phones  []string `validate:"len:11"`
		meta    json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr []error
	}{
		{
			User{
				ID:      "111111111111111111111111111111111111",
				Name:    "Holder",
				Age:     50,
				Email:   "test@test.ru",
				Role:    "admin,stuff",
				RoleInt: 100,
				Phones:  []string{"89876543212", "89876543212"},
				meta:    nil,
			},
			nil,
		},
		{
			User{
				ID:      "test string length with not correct length",
				Name:    "Holder",
				Age:     50,
				Email:   "test@test.ru",
				Role:    "admin,stuff",
				RoleInt: 100,
				Phones:  []string{"89876543212", "89876543212"},
				meta:    nil,
			},
			[]error{ErrStringLengthInvalid},
		},
		{
			User{
				ID:      "111111111111111111111111111111111111",
				Name:    "Holder",
				Age:     55,
				Email:   "test@test.ru",
				Role:    "admin,stuff",
				RoleInt: 100,
				Phones:  []string{"898765", "89876543212"},
				meta:    nil,
			},
			[]error{ErrIntMaxInvalid, ErrStringLengthInvalid},
		},
		{
			Token{
				[]byte{1, 2, 3},
				[]byte{1, 2, 3},
				[]byte{1, 2, 3},
			},
			nil,
		},
		{
			App{
				"1.2.3.6",
			},
			[]error{ErrStringLengthInvalid},
		},
		{
			Response{
				401,
				"Body: Unauthorized",
			},
			[]error{ErrIntInInvalid},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			var validationErrors ValidationErrors
			errs := Validate(tt.in)
			if len(tt.expectedErr) == 0 {
				require.NoError(t, errs)
				return
			}

			require.ErrorAs(t, errs, &validationErrors)
			for i, e := range validationErrors {
				require.ErrorIs(t, e.Err, tt.expectedErr[i])
			}
		})
	}
}
