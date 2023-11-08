package shape

import (
	"testing"
)

func TestTriangleArea(t *testing.T) {
	// Create a triangle with known values
	t1 := NewTriangle(3.0, 4.0, 5.0)

	// Calculate the expected area
	expectedArea := 0.5 * 3.0 * 4.0

	// Call the Area method and compare the result with the expected value
	if actualArea := t1.Area(); actualArea != expectedArea {
		t.Errorf("Expected area %f, but got %f", expectedArea, actualArea)
	}
}

func TestTrianglePerimeter(t *testing.T) {
	// Create a triangle with known values
	t1 := NewTriangle(3.0, 4.0, 5.0)

	// Calculate the expected perimeter
	expectedPerimeter := 3.0 + 4.0 + 5.0

	// Call the Perimeter method and compare the result with the expected value
	if actualPerimeter := t1.Perimeter(); actualPerimeter != expectedPerimeter {
		t.Errorf("Expected perimeter %f, but got %f", expectedPerimeter, actualPerimeter)
	}
}
