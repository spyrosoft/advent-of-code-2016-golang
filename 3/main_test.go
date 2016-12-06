package main

import "testing"

func TestTriangles(t *testing.T) {
	input = `    541  588  421
  827  272  126
  660  514  367
   39  703  839
  229  871    3
  237  956  841`
	main()
	if totalPossibleHorizontalTriangles != 3 {
		t.Error("Expected number of possible horizontal triangles, 3, is not:", totalPossibleHorizontalTriangles)
	}
	if totalPossibleVerticalTriangles != 6 {
		t.Error("Expected number of possible vertical triangles, 6, is not:", totalPossibleVerticalTriangles)
	}
}
