package maximum_perimeter_triangle

import "testing"

func TestMaxPerimeterTriangle(t *testing.T) {
	if a, b, c := maximum_perimeter_isosceles_triangle([]int{1, 1, 1, 3, 3}); a != 1 || b != 3 || c != 3 {
		t.Errorf("%d %d %d\n", a, b, c)
	}
	if a, b, c := maximum_perimeter_isosceles_triangle([]int{1, 1, 1, 3, 3, 5}); a != 3 || b != 3 || c != 5 {
		t.Errorf("%d %d %d\n", a, b, c)
	}
	if a, b, c := maximum_perimeter_isosceles_triangle([]int{1, 1, 1, 3, 4, 5}); a != 1 || b != 1 || c != 1 {
		t.Errorf("%d %d %d\n", a, b, c)
	}
	if a, b, c := maximum_perimeter_isosceles_triangle([]int{1, 1, 2, 3, 4, 5}); a != -1 || b != -1 || c != -1 {
		t.Errorf("%d %d %d\n", a, b, c)
	}

	if a, b, c := maximum_perimeter_triangle([]int{76361, 69365, 78109, 44475, 13538, 5084, 1371, 6453, 83726, 94527, 401820417, 13839263, 768933, 22392673, 153481704, 88347, 44593, 5285782, 191555, 58624780, 248338678, 26826, 70078, 97005, 26139, 25664, 68595, 83568, 36232035, 3267536, 78837, 22863, 36133, 2409, 44739, 2018202, 35222, 650159152, 1249262, 31720, 66102, 39532, 94856906, 77216, 8553388, 480240, 288620, 59954, 36981}); a != 88347 || b != 94527 || c != 97005 {
		t.Errorf("%d %d %d \n", a, b, c)
	}
	if a, b, c := maximum_perimeter_triangle([]int{1, 1, 1, 3, 3}); a != 1 || b != 3 || c != 3 {
		t.Errorf("%d %d %d\n", a, b, c)
	}
	if a, b, c := maximum_perimeter_triangle([]int{1, 1, 1, 3, 3, 5}); a != 3 || b != 3 || c != 5 {
		t.Errorf("%d %d %d\n", a, b, c)
	}
	if a, b, c := maximum_perimeter_triangle([]int{1, 1, 1, 3, 4, 5}); a != 3 || b != 4 || c != 5 {
		t.Errorf("%d %d %d\n", a, b, c)
	}
	if a, b, c := maximum_perimeter_triangle([]int{1, 1, 2, 3, 4, 5}); a != 3 || b != 4 || c != 5 {
		t.Errorf("%d %d %d\n", a, b, c)
	}
	if a, b, c := maximum_perimeter_triangle([]int{1, 2, 3}); a != -1 || b != -1 || c != -1 {
		t.Errorf("%d %d %d\n", a, b, c)
	}
}
