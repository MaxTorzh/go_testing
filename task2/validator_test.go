package task2

import (
	"errors"
	"testing"
)

func TestValidateNumber(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		wantErr error
	}{
		{
			name:    "positive number in the range",
			input:   42,
			wantErr: nil,
		},
		{
			name:    "border number",
			input:   1,
			wantErr: nil,
		},
		{
			name:    "border number",
			input:   100,
			wantErr: nil,
		},
		{
			name:    "negative number",
			input:   -5,
			wantErr: ErrNegativeNumber,
		},
		{
			name:    "zero",
			input:   0,
			wantErr: ErrZeroNumber,
		},
		{
			name:    "too large number",
			input:   101,
			wantErr: ErrTooLarge,
		},
		{
			name:    "huge number",
			input:   1000,
			wantErr: ErrTooLarge,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateNumber(tt.input)

			if tt.wantErr == nil {
				if err != nil {
					t.Errorf("ValidateNumber(%d) ends with error: %v, but expected success", tt.input, err)
				}
				return
			}

			if err == nil {
				t.Errorf("ValidateNumber(%d) ends without error, but expected: %v", tt.input, tt.wantErr)
				return
			}

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("ValidateNumber(%d) returns incorrect error.\nReceived: %v\nExpected: %v", tt.input, err, tt.wantErr)
			}
			t.Logf("Expected error was received: %v", err)
		})
	}
}
