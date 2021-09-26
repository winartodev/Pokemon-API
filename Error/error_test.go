package errorhandler

import (
	response "Pokemon-API/Response"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetErrorCode(t *testing.T) {
	type args struct {
		code int
		err  error
	}
	tests := []struct {
		name string
		args args
		want *response.BodyError
	}{
		// error status not found.
		{
			name: "Get Status Not Found",
			args: args{
				code: http.StatusNotFound,
				err:  errors.New("id pokemon 2 not found"),
			},
			want: &response.BodyError{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Error:  errors.New("id pokemon 2 not found").Error(),
			},
		},
		// error status bad request
		{
			name: "Get Status Bad Request",
			args: args{
				code: http.StatusBadRequest,
				err:  errors.New("duplicate entry '1' for key 'PRIMARY"),
			},
			want: &response.BodyError{
				Code:   http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Error:  errors.New("duplicate entry '1' for key 'PRIMARY").Error(),
			},
		},
		// error internal server error
		{
			name: "Get Status Internal Server Error",
			args: args{
				code: http.StatusInternalServerError,
				err:  errors.New("Internal Server Error"),
			},
			want: &response.BodyError{
				Code:   http.StatusInternalServerError,
				Status: http.StatusText(http.StatusInternalServerError),
				Error:  errors.New("Internal Server Error").Error(),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GetErrorCode(tc.args.code, tc.args.err)
			assert.Equal(t, got, tc.want)
		})
	}
}
