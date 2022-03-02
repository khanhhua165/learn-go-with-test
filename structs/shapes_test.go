package structs

import "testing"

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, g, w float64) {
		t.Helper()
		if g != w {
			t.Errorf("got %g want %g", g, w)
		}
	}
	rec := Rectangle{10.0, 10.0}
	got := rec.Perimeter()
	want := 40.0
	checkPerimeter(t, got, want)

}

type AreaTest struct {
	shape Shape
	want  float64
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, at AreaTest) {
		t.Helper()
		g := at.shape.Area()
		if g != at.want {
			t.Errorf("%#v got %g want %g", at.shape, g, at.want)
		}
	}

	areaTests := []AreaTest{
		{shape: Rectangle{12, 6}, want: 72.0},
		{shape: Circle{10}, want: 314.1592653589793},
		{shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		checkArea(t, tt)
	}
}
