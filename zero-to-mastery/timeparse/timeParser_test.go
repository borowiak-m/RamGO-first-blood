package timeparse

import "testing"

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"19:00:12", true},
		{"1:3:44", true},
		{"bad", false},
		{"1:-3:44", false},
		{"0:59:59", true},
		{"", false},
		{"11:11", false},
		{"aa:bb:cc", false},
		{"5:23", false},
		{"12:12:12", true},
		{"54:03:15", false},
		{"3:76:20", false},
		{"4:5:88", false},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			t.Errorf("%v : %v, error should be nil", data.time, err)
		}
	}
}
