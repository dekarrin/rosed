package gem


// data extracted from https://www.unicode.org/Public/UCD/latest/ucd/extracted/DerivedCombiningClass.txt
// on Nov 5, 2022.
func getCCC(r rune) int {
	inRange := func(start rune, end rune) bool {
		return (start <= r && r <= end) 
	}
	
	if inRange(0x0334, 0x0338) ||
	   r == 0x1cd4 ||
	   inRange(0x1CE2, 0x1CE8) ||
	   inRange(0x1CE2, 0x1CE8) ||
	   inRange(0x1CE2, 0x1CE8) ||
	   inRange(0x1CE2, 0x1CE8) ||
	   inRange(0x1CE2, 0x1CE8) ||
	   inRange(0x1CE2, 0x1CE8) ||
	   inRange(0x1CE2, 0x1CE8) ||
	   {
		return 1
	} else {
		return 0
	}
}
