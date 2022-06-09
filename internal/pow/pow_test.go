package pow

import "testing"

func TestPow_New(t *testing.T) {
	mockPowGenerator, err := New()
	if err != nil {
		t.Error(err)
	}

	if mockPowGenerator == nil {
		t.Errorf("expected: value of type generator to be not nil")
	}
}

func TestPow_GetPowSolution(t *testing.T) {
	mockPowGenerator, err := New()
	if err != nil {
		t.Error(err)
	}

	solution, err := mockPowGenerator.GetPowSolution()
	if err != nil {
		t.Error(err)
	}

	if len(solution) == 0 {
		t.Errorf("expected: value of solution to be not 0")
	}
}

func TestPow_Verify(t *testing.T) {
	mockPowGenerator, err := New()
	if err != nil {
		t.Error(err)
	}

	solution, err := mockPowGenerator.GetPowSolution()
	if err != nil {
		t.Error(err)
	}

	valid, err := mockPowGenerator.Verify(solution)
	if err != nil {
		t.Error(err)
	}

	if !valid {
		t.Errorf("expected: solution to be valid")
	}
}

func TestPow_Verify_Invalid(t *testing.T) {
	mockPowGenerator, err := New()
	if err != nil {
		t.Error(err)
	}

	valid, err := mockPowGenerator.Verify("INVALID_TEST_SOLUTION")
	if err == nil {
		t.Errorf("expected: err to be not nil")
	}

	if valid {
		t.Errorf("expected: solution to be invalid")
	}
}
