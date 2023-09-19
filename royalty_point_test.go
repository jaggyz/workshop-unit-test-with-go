package royaltyPoint

import "testing"

func TestPointShouldBe_WhenPriceIs_(t *testing.T) {
	expected := 
	if actual := point(); actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
