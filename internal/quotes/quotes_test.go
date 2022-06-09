package quotes

import "testing"

func TestQuotes_New(t *testing.T) {
	mockQuotesPicker := New(Quotes)

	if mockQuotesPicker == nil {
		t.Errorf("expected: value of type quotesPicker to be not nil")
	}
}

func TestQuotes_PickQuoteIndex(t *testing.T) {
	tests := []struct {
		mockQuotesPicker  Picker
		shouldReturnError bool
	}{
		{
			mockQuotesPicker:  New(Quotes),
			shouldReturnError: false,
		},
		{
			mockQuotesPicker:  New(make([]Quote, 0)),
			shouldReturnError: true,
		},
	}

	for _, test := range tests {
		idx, err := test.mockQuotesPicker.PickQuoteIndex()
		if err != nil && !test.shouldReturnError {
			t.Errorf("expected error to be nil: got %s", err.Error())
		}

		if idx < 0 || (idx > int64(len(Quotes)-1)) {
			t.Errorf("expected idx to be between o and %d: got %d", len(Quotes)-1, idx)
		}
	}
}

func TestQuotes_GetQuote(t *testing.T) {
	tests := []struct {
		mockQuotesPicker  Picker
		shouldReturnError bool
	}{
		{
			mockQuotesPicker:  New(Quotes),
			shouldReturnError: false,
		},
		{
			mockQuotesPicker:  New(Quotes),
			shouldReturnError: true,
		},
	}

	for _, test := range tests {
		idx, err := test.mockQuotesPicker.PickQuoteIndex()
		if err != nil {
			t.Errorf("expected error to be nil: got %s", err.Error())
		}

		if idx < 0 || (idx > int64(len(Quotes)-1)) {
			t.Errorf("expected idx to be between o and %d: got %d", len(Quotes)-1, idx)
		}

		if test.shouldReturnError {
			idx = 0
		}

		quote, err := test.mockQuotesPicker.GetQuote(idx)
		if err != nil {
			t.Errorf("expected error to be nil: got %s", err.Error())
		}

		if len(quote.Quote) == 0 {
			t.Errorf("expected quote to be not empty")
		}
	}
}
