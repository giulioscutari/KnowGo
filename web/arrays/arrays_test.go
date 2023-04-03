package arrays

import "testing"

func TestSum(t *testing.T) {

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("Got %d want %d given %v", got, want, numbers)
		}
	})
}
