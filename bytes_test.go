package human_test

import (
	"testing"

	"github.com/quarkey/go-human"
)

func TestBytesToHuman(t *testing.T) {
	tests := []struct {
		b int64
		h string // wanted
	}{
		{1024, "1.00KB"},
		{10737418240, "10.00GB"},
		{1099511627776, "1.00TB"},
	}
	for _, c := range tests {
		got := human.BytesToHuman(c.b)
		if got != c.h {
			t.Errorf("BytesToHuman(%d) got == %q want, %q", c.b, got, c.h)
		}
		t.Logf("BytesToHuman(%d) got == %q want, %q", c.b, got, c.h)
	}

	got := human.BytesToHuman(0)
	if got != "0.00B" {
		t.Errorf("BytesToHuman(0) got == %q want, %q", got, "0.00B")
	}

	got = human.BytesToHuman(-10)
	if got != "0.00B" {
		t.Errorf("BytesToHuman(0) got == %q want, %q", got, "0.00B")
	}

}

func TestToBytes(t *testing.T) {
	tests := []struct {
		b float64
		h string
	}{
		{1024.00, "1.00KB"},
		{16976459532861.439453, "15.44TB"},
	}
	for _, c := range tests {
		got, err := human.ToBytes(c.h)
		if got != c.b || err != nil {
			t.Errorf("ToBytes(%s) got == %g want, %g", c.h, got, c.b)
		}
		t.Logf("ToBytes(%s) got == %g want, %g", c.h, got, c.b)
	}
}
