package models

import (
	"testing"
)

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
	}{
		{
			name: "valid user",
			user: User{
				Name:        "John Doe",
				Age:         30,
				PhoneNumber: "+1234567890",
				Email:       "john@example.com",
			},
			wantErr: false,
		},
		{
			name: "invalid email",
			user: User{
				Name:        "John Doe",
				Age:         30,
				PhoneNumber: "+1234567890",
				Email:       "invalid-email",
			},
			wantErr: true,
		},
		{
			name: "invalid age",
			user: User{
				Name:        "John Doe",
				Age:         -1,
				PhoneNumber: "+1234567890",
				Email:       "john@example.com",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
} 