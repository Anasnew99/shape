package shape

import (
	"testing"
)

func TestCircleArea(t *testing.T) {
	// Create a circle with a known radius
	c1 := NewCircle(5.0)

	// Calculate the expected area
	expectedArea := 3.14 * 5.0 * 5.0

	// Call the Area method and compare the result with the expected value
	if actualArea := c1.Area(); actualArea != expectedArea {
		t.Errorf("Expected area %f, but got %f", expectedArea, actualArea)
	}
}

func TestCirclePerimeter(t *testing.T) {
	// Create a circle with a known radius
	c1 := NewCircle(5.0)

	// Calculate the expected perimeter
	expectedPerimeter := 2 * 3.14 * 5.0

	// Call the Perimeter method and compare the result with the expected value
	if actualPerimeter := c1.Perimeter(); actualPerimeter != expectedPerimeter {
		t.Errorf("Expected perimeter %f, but got %f", expectedPerimeter, actualPerimeter)
	}
}
