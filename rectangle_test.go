package shape

import (
	"testing"
)

func TestRectangleArea(t *testing.T) {
	// Create a rectangle with known values
	r1 := NewRectangle(3.0, 4.0)

	// Calculate the expected area
	expectedArea := 3.0 * 4.0

	// Call the Area method and compare the result with the expected value
	if actualArea := r1.Area(); actualArea != expectedArea {
		t.Errorf("Expected area %f, but got %f", expectedArea, actualArea)
	}
}

func TestRectanglePerimeter(t *testing.T) {
	// Create a rectangle with known values
	r1 := NewRectangle(3.0, 4.0)

	// Calculate the expected perimeter
	expectedPerimeter := 2 * (3.0 + 4.0)

	// Call the Perimeter method and compare the result with the expected value
	if actualPerimeter := r1.Perimeter(); actualPerimeter != expectedPerimeter {
		t.Errorf("Expected perimeter %f, but got %f", expectedPerimeter, actualPerimeter)
	}
}
