package sayhello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello World!"
	got := Hello()

	if want != got {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
