package any_test

import (
	"testing"
	"time"

	"github.com/hyqe/any"
)

func TestTime(t *testing.T) {
	x := any.Time()
	if x.IsZero() {
		t.Fatal("generated zero time")
	}
}

func TestTimeSpan(t *testing.T) {
	t.Parallel()
	t.Run("one_option", func(t *testing.T) {
		now := time.Now()
		x := any.TimeSpan(now, now)
		if x != now {
			t.Fatalf("want: %v got: %v", now, x)
		}
	})
	t.Run("start_after_end", func(t *testing.T) {
		start := time.Time{}.Add(time.Hour * 2)
		end := time.Time{}.Add(time.Hour * 1)
		x := any.TimeSpan(start, end)
		if start.Before(x) && start.After(x) {
			t.Fatalf("want between: '%v' - '%v' got: '%v'", start, end, x)
		}
	})
	t.Run("low", func(t *testing.T) {
		start := time.Time{}
		end := time.Time{}.Add(1)
		x := any.TimeSpan(start, end)
		if start.After(x) || end.Before(x) {
			t.Fatalf("want between: '%v' - '%v' got: '%v'", start, end, x)
		}
	})
	t.Run("mid", func(t *testing.T) {
		start := time.Time{}
		end := time.Time{}.Add(92233720)
		x := any.TimeSpan(start, end)
		if start.After(x) || end.Before(x) {
			t.Fatalf("want between: '%v' - '%v' got: '%v'", start, end, x)
		}
	})
	t.Run("high", func(t *testing.T) {
		start := time.Time{}
		end := time.Time{}.Add(9223372036854775805)
		x := any.TimeSpan(start, end)
		if start.After(x) || end.Before(x) {
			t.Fatalf("want between: '%v' - '%v' got: '%v'", start, end, x)
		}
	})
}
