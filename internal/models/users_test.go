package models

import (
	"testing"

	"chetraseng.com/internal/assert"
)

func TestUserModelExists(t *testing.T) {
	tests := []struct {
		name   string
		userID int
		want   bool
	}{
		{
			name:   "Valid ID",
			userID: 1,
			want:   true,
		},
		{
			name:   "Zero ID",
			userID: 0,
			want:   false,
		},
		{
			name:   "Non-existent ID",
			userID: 2,
			want:   false,
		},
	}

	// Skip test if short flag is provided
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db := newTestDB(t)

			m := UserModel{db}

			exists, err := m.Exists(test.userID)

			assert.Equal(t, exists, test.want)
			assert.NilError(t, err)
		})
	}

}
