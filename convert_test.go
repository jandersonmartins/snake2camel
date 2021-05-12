package s2c

import (
	"fmt"
	"testing"
)

func TestConvertToCamel(t *testing.T) {
	t.Run("basic convertion", func(t *testing.T) {
		snake := "hello_world"
		got := ConvertToCamel(snake)
		want := "helloWorld"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func ExampleConvertToCamel() {
	snake := "hello_world"
	got := ConvertToCamel(snake)
	fmt.Println(got)
	//Output: helloWorld
}

func BenchmarkConvertToCamel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertToCamel("hello_world")
	}
}
