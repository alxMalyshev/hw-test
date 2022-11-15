package hw09structvalidator

import (
	//nolint
	"encoding/json"
	//nolint
	"errors"
	//nolint
	"fmt"
	//nolint
	"github.com/stretchr/testify/require"
	//nolint
	"testing"
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
		expectedErr error
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
			ErrStringLengthInvalid,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			errs := Validate(tt.in)
			if tt.expectedErr == nil {
				require.NoError(t, errs)
				return
			}
			if !errors.Is(errs, tt.expectedErr) {
				t.Fatalf("err %v is not %v", errs, tt.expectedErr)
			}
			// fmt.Println(errors.Unwrap(errors.Unwrap(errs)) == ErrStringLengthInvalid)
			// require.ErrorAs(t, errs, &validationError)
		})
	}
}
