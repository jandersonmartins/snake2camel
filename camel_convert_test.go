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

		checkStrings(t, got, want)
	})

	t.Run("word ending with snake char", func(t *testing.T) {
		snake := "hello_world_"
		got := ConvertToCamel(snake)
		want := "helloWorld"

		checkStrings(t, got, want)
	})

	t.Run("word started with snake char", func(t *testing.T) {
		snake := "_hello_world"
		got := ConvertToCamel(snake)
		want := "_helloWorld"

		checkStrings(t, got, want)
	})
}

func checkStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
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
