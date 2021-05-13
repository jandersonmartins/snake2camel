package snake2camel

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

	t.Run("convert a large text", func(t *testing.T) {
		text := `
			function foo() {
				const user_name = "example"
				const user_age = 10
			}
		`

		got := ConvertToCamel(text)
		want := `
			function foo() {
				const userName = "example"
				const userAge = 10
			}
		`

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
		ConvertToCamel(`
			lorem ipsum
			foo_bar example
		`)
	}
}
