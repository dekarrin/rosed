package gem

// This file implements algorithms described in the Unicode Standard
// section 3.11 "Normalization Forms", version 15.0.0 as taken on
// November 5th, 2022. Specifically it is going to implement normalization
// of strings to NFC (Canonical recomposed form).

// NOTE: unicode standard 3.11 "Normalization Forms" explicitly calls out
// that collation is v different from normalization and that those algorithms
// should be used if collation is being done (which is the only reason we
// are doing normalization to begin with
//
// However, collation as defined in Unicode UTS #10 is quite a bit more complex
// than normalization and it appears to use certain normalization forms of
// subsequences of strings being collated for the cases of some characters, so
// it seems we'll need to implement composition and decomposition at some
// point anyways so may as well start with the 'lazy' approach of just
// converting the two strings to be collated to NFC and comparing those.

// Applies canonical decomposition to every rune in the in string until there are no more changes.
func getNFD(runes []rune) []rune {
	if len(runes) < 1 {
		return nil
	}
	
	var decomposed []rune
	
	changed := true
	for changed {
		decomposed = make([]rune, 0)
		changed = false
		for _, ch := range runes {
			decompType, decompMapping := getDecompositionMapping(ch)
			
			if decompType != "" {
				// only do canonical decomps
				decompMapping = []rune{ch}
			}
			
			// now check if we do anything different
			if len(decompMapping) > 1 || decompMapping[0] != ch {
				changed = true
			}
			
			// and keep the decomposed mapping
			decomposed = append(decomposed, decompMapping...)
		}
		runes = decomposed
	}
	
	return decomposed
}
