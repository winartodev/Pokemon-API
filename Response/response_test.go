package response

import (
	pokemon "Pokemon-API/Pokemon"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResponseError(t *testing.T) {
	type args struct {
		result *BodyError
	}
	tests := []struct {
		name string
		args args
		want *BodyError
	}{
		{
			name: "Response Body Error",
			args: args{
				&BodyError{
					Code:   http.StatusBadRequest,
					Status: http.StatusText(http.StatusBadRequest),
					Error:  "dial tcp 127.0.0.1:3306: connect: connection refused",
				},
			},
			want: &BodyError{
				Code:   http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Error:  "dial tcp 127.0.0.1:3306: connect: connection refused",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := NewResponseError(tc.args.result)
			assert.Equal(t, got, tc.want)
		})
	}
}

func TestNewResponseSuccess(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want *BodySuccess
	}{
		{
			name: "Response Body Error",
			args: args{
				data: `{"id":1,"name":"Pikachu","species":"Mouse"}`,
			},
			want: &BodySuccess{
				Code:    http.StatusOK,
				Status:  http.StatusText(http.StatusOK),
				Message: "Success",
				Data:    `{"id":1,"name":"Pikachu","species":"Mouse"}`,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := NewResponseSuccess(tc.args.data)
			assert.Equal(t, got, tc.want)
		})
	}
}

func TestNewPokemonCreateSuccess(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want *BodySuccess
	}{
		{
			name: "New Pokemon Create Success",
			args: args{
				id: 1,
			},
			want: &BodySuccess{
				Code:    http.StatusOK,
				Status:  http.StatusText(http.StatusOK),
				Message: "Success",
				Data:    map[string]string{"status": "id pokemon 1 create"},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := NewPokemonCreateSuccess(tc.args.id)
			assert.Equal(t, got, tc.want)
		})
	}
}

func TestRespondWithJSON(t *testing.T) {
	type args struct {
		w      http.ResponseWriter
		result interface{}
		code   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Respond With JSON Body Success",
			args: args{
				w: httptest.NewRecorder(),
				result: BodySuccess{
					Code:    http.StatusOK,
					Status:  http.StatusText(http.StatusOK),
					Message: "Success",
					Data: pokemon.Entity{
						ID:      1,
						Name:    "Pikachu",
						Species: "Mouse Pokemon",
					},
				},
				code: http.StatusOK,
			},
		},
		{
			name: "Test Respond With JSON Body Error",
			args: args{
				w: httptest.NewRecorder(),
				result: BodyError{
					Code:   http.StatusBadRequest,
					Status: http.StatusText(http.StatusBadRequest),
					Error:  errors.New("id pokemon 1 not found").Error(),
				},
				code: http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RespondWithJSON(tt.args.w, tt.args.result, tt.args.code)
		})
	}
}
