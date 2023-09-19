package royaltyPoint

import "testing"

func TestPointShouldBe0WhenPriceIs80(t *testing.T) {
	expected := 0
	if actual := point(80.00); actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
