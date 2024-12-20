package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("add to numbers", func(t *testing.T) {
		got := add(3, 4)

		want := 7

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func ExampleAdd() {
	sum := add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
