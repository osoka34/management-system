package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTokens(t *testing.T) {
	type args struct {
		username string
		uid      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name:    "success",
			wantErr: false,
			args: args{
				username: "username",
				uid:      "uid",
			},
			err: nil,
		},

		{
			name:    "empty args",
			wantErr: true,
			args:    args{},
			err:     EmptyArgs,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bearer, refresh, err := GenerateTokens(tt.args.username, tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateTokens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("bearer: %s, refresh: %s", bearer, refresh)
		})
	}
}

func TestExtractAccessClaims(t *testing.T) {
	type args struct {
		username string
		uid      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name:    "success",
			wantErr: false,
			args: args{
				username: "username",
				uid:      "uid",
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bearer, refresh, err := GenerateTokens(tt.args.username, tt.args.uid)
			if err != nil {
				t.Errorf("GenerateTokens() error = %v", err)
				return
			}

			claimsA, err := ExtractAccessClaims(bearer)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractAccessClaims() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.args.username, claimsA.Username)
			assert.Equal(t, tt.args.uid, claimsA.UID)

			claimsR, err := ExtractRefreshClaims(refresh)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractRefreshClaims() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.args.username, claimsR.Username)
			assert.Equal(t, tt.args.uid, claimsR.UID)
		})
	}
}
