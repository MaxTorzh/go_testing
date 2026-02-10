package task9

import (
	"errors"
	"strings"
	"testing"
)

func TestParsePositiveNumber_ConcreteErrors(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  error 
		wantCont string 
	}{
		{
			name:     "пустая строка - ErrInvalidInput",
			input:    "",
			wantErr:  ErrInvalidInput,
			wantCont: "empty string",
		},
		{
			name:     "не число - ErrInvalidInput",
			input:    "abc",
			wantErr:  ErrInvalidInput,
			wantCont: "invalid input",
		},
		{
			name:     "отрицательное число - ErrNegativeNumber",
			input:    "-10",
			wantErr:  ErrNegativeNumber,
			wantCont: "negative number",
		},
		{
			name:     "слишком большое число - ErrTooLarge",
			input:    "2000",
			wantErr:  ErrTooLarge,
			wantCont: "> 1000",
		},
		{
			name:     "успешный парсинг",
			input:    "42",
			wantErr:  nil,
			wantCont: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParsePositiveNumber(tt.input)

			if tt.wantErr == nil {
				if err != nil {
					t.Errorf("ожидали успех, получили ошибку: %v", err)
				}
				return
			}

			if err == nil {
				t.Fatal("ожидали ошибку, получили nil")
			}

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("ошибка = %v, ожидали %v", err, tt.wantErr)
			}

			if tt.wantCont != "" && !strings.Contains(err.Error(), tt.wantCont) {
				t.Errorf("текст ошибки = %q, должен содержать %q", err.Error(), tt.wantCont)
			}
		})
	}
}

func TestSafeDivide_ConcreteError(t *testing.T) {
	result, err := SafeDivide(10, 2)
	if err != nil {
		t.Fatalf("неожиданная ошибка: %v", err)
	}
	if result != 5 {
		t.Errorf("результат = %d, ожидали 5", result)
	}

	_, err = SafeDivide(10, 0)
	if err == nil {
		t.Fatal("ожидали ошибку деления на ноль")
	}

	if !errors.Is(err, ErrDivByZero) {
		t.Errorf("ошибка = %v, ожидали %v", err, ErrDivByZero)
	}

	if err.Error() != "division by zero" {
		t.Errorf("текст ошибки = %q, ожидали %q", err.Error(), "division by zero")
	}
}

func TestValidateAge_ErrorWrapping(t *testing.T) {
	tests := []struct {
		name    string
		age     int
		wantErr error
	}{
		{
			name:    "отрицательный возраст - ErrInvalidInput",
			age:     -5,
			wantErr: ErrInvalidInput,
		},
		{
			name:    "слишком большой возраст - ErrTooLarge",
			age:     150,
			wantErr: ErrTooLarge,
		},
		{
			name:    "возраст меньше 18 - не наша кастомная ошибка",
			age:     15,
			wantErr: nil,
		},
		{
			name:    "валидный возраст",
			age:     25,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAge(tt.age)

			if tt.wantErr == nil && err == nil {
				return
			}

			if tt.wantErr == nil && err != nil {
				if errors.Is(err, ErrInvalidInput) || errors.Is(err, ErrTooLarge) {
					t.Errorf("получили кастомную ошибку %v, а ожидали простую", err)
				}
				return
			}

			if err == nil {
				t.Fatal("ожидали ошибку, получили nil")
			}

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("ошибка = %v, ожидали %v", err, tt.wantErr)
			}
		})
	}
}

func TestErrorEquality(t *testing.T) {
	err1 := errors.New("something went wrong")
	err2 := errors.New("something went wrong")
	
	if err1 == err2 {
		t.Error("две разные ошибки с одинаковым текстом не должны быть равны")
	}

	if ErrInvalidInput != ErrInvalidInput {
		t.Error("одна и та же кастомная ошибка должна быть равна сама себе")
	}

	err := ErrInvalidInput
	if !errors.Is(err, ErrInvalidInput) {
		t.Error("errors.Is должна находить обернутую ошибку")
	}
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}