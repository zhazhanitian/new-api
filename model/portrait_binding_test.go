package model

import "testing"

func TestAggregatePortraitStatus(t *testing.T) {
	cases := []struct {
		name     string
		statuses []string
		want     string
	}{
		{"empty", nil, ""},
		{"all active", []string{"Active", "Active"}, "Active"},
		{"one failed wins", []string{"Active", "Failed"}, "Failed"},
		{"failed beats processing", []string{"Processing", "Failed"}, "Failed"},
		{"active and processing", []string{"Active", "Processing"}, "Processing"},
		{"active and submitted", []string{"Active", "Submitted"}, "Processing"},
		{"all submitted", []string{"Submitted", "Submitted"}, "Submitted"},
		{"submitted and processing", []string{"Submitted", "Processing"}, "Processing"},
		{"single active", []string{"Active"}, "Active"},
	}
	for _, tc := range cases {
		if got := AggregatePortraitStatus(tc.statuses); got != tc.want {
			t.Errorf("%s: got %q, want %q", tc.name, got, tc.want)
		}
	}
}
