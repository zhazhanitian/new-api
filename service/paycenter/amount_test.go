package paycenter

import "testing"

func TestYuanToFenRoundsHalfAwayFromZero(t *testing.T) {
	cases := []struct {
		yuan float64
		fen  int64
	}{
		{0.001, 0},
		{0.004, 0},
		{0.005, 1},
		{0.01, 1},
		{1, 100},
		{1.004, 100},
		{1.005, 101},
		{7.3, 730},
		{10.1, 1010},
	}
	for _, tc := range cases {
		if got := YuanToFen(tc.yuan); got != tc.fen {
			t.Fatalf("YuanToFen(%v) = %d, want %d", tc.yuan, got, tc.fen)
		}
	}
}

func TestFenRoundTrip(t *testing.T) {
	for _, fen := range []int64{1, 10, 73, 100, 730, 1010, 999999} {
		got := YuanToFen(FenToYuan(fen))
		if got != fen {
			t.Fatalf("round trip fen %d -> %d", fen, got)
		}
	}
}
