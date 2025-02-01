
package hello

import "testing"

func TestGreet(t *testing.T) {
    expected := "Hello, World!"
    if result := Greet("World"); result != expected {
        t.Errorf("Expected '%s', got '%s'", expected, result)
    }
}
