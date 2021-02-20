package gem

// Values are taken from UAX #29 Table 2,
// "Grapheme_Cluster_Break Property Values".
//
// All values are included except those listed as obsolete and unused in the
// table.

// isCbPrepend returns whether the test of Grapheme_Cluster_Break = Prepend
// evaluates to true.
//
// Data for this function was obtained from
// http://www.unicode.org/Public/UCD/latest/ucd/auxiliary/GraphemeBreakProperty.txt
// on Feb 20, 2021, and is used under the terms of the Unicode License,
// available at the following address: https://www.unicode.org/license.html
func isCbPrepend(r rune) bool {
	return (0x0600 <= r && r <= 0x0605) ||
		(0x06DD == r) ||
		(0x070F == r) ||
		(0x08E2 == r) ||
		(0x0D4E == r) ||
		(0x110BD == r) ||
		(0x110CD == r) ||
		(0x111C2 <= r && r <= 0x111C3) ||
		(0x1193F == r) ||
		(0x11941 == r) ||
		(0x11A3A == r) ||
		(0x11A84 <= r && r <= 0x11A89) ||
		(0x11D46 == r)
}

// isCbCR returns whether the test of Grapheme_Cluster_Break = CR
// evaluates to true.
//
// Data for this function was obtained from
// http://www.unicode.org/Public/UCD/latest/ucd/auxiliary/GraphemeBreakProperty.txt
// on Feb 20, 2021, and is used under the terms of the Unicode License,
// available at the following address: https://www.unicode.org/license.html
func isCbCR(r rune) bool {
	return (0x000D == r)
}

// isCbLF returns whether the test of Grapheme_Cluster_Break = LF
// evaluates to true.
//
// Data for this function was obtained from
// http://www.unicode.org/Public/UCD/latest/ucd/auxiliary/GraphemeBreakProperty.txt
// on Feb 20, 2021, and is used under the terms of the Unicode License,
// available at the following address: https://www.unicode.org/license.html
func isCbLF(r rune) bool {
	return (0x000A == r)
}

// isCbControl returns whether the test of Grapheme_Cluster_Break = Control
// evaluates to true.
//
// Data for this function was obtained from
// http://www.unicode.org/Public/UCD/latest/ucd/auxiliary/GraphemeBreakProperty.txt
// on Feb 20, 2021, and is used under the terms of the Unicode License,
// available at the following address: https://www.unicode.org/license.html
func isCbControl(r rune) bool {
	return (0x0000 <= r && r <= 0x0009) ||
		(0x000B <= r && r <= 0x000C) ||
		(0x000E <= r && r <= 0x001F) ||
		(0x007F <= r && r <= 0x009F) ||
		(0x00AD == r) ||
		(0x061C == r) ||
		(0x180E == r) ||
		(0x200B == r) ||
		(0x200E <= r && r <= 0x200F) ||
		(0x2028 == r) ||
		(0x2029 == r) ||
		(0x202A <= r && r <= 0x202E) ||
		(0x2060 <= r && r <= 0x2064) ||
		(0x2065 == r) ||
		(0x2066 <= r && r <= 0x206F) ||
		(0xFEFF == r) ||
		(0xFFF0 <= r && r <= 0xFFF8) ||
		(0xFFF9 <= r && r <= 0xFFFB) ||
		(0x13430 <= r && r <= 0x13438) ||
		(0x1BCA0 <= r && r <= 0x1BCA3) ||
		(0x1D173 <= r && r <= 0x1D17A) ||
		(0xE0000 == r) ||
		(0xE0001 == r) ||
		(0xE0002 <= r && r <= 0xE001F) ||
		(0xE0080 <= r && r <= 0xE00FF) ||
		(0xE01F0 <= r && r <= 0xE0FFF)
}

// isCbExtend returns whether the test of Grapheme_Cluster_Break = Extend
// evaluates to true.
//
// Data for this function was obtained from
// http://www.unicode.org/Public/UCD/latest/ucd/auxiliary/GraphemeBreakProperty.txt
// on Feb 20, 2021, and is used under the terms of the Unicode License,
// available at the following address: https://www.unicode.org/license.html
func isCbExtend(r rune) bool {
	return (0x0300 <= r && r <= 0x036F) ||
		(0x0483 <= r && r <= 0x0487) ||
		(0x0488 <= r && r <= 0x0489) ||
		(0x0591 <= r && r <= 0x05BD) ||
		(0x05BF == r) ||
		(0x05C1 <= r && r <= 0x05C2) ||
		(0x05C4 <= r && r <= 0x05C5) ||
		(0x05C7 == r) ||
		(0x0610 <= r && r <= 0x061A) ||
		(0x064B <= r && r <= 0x065F) ||
		(0x0670 == r) ||
		(0x06D6 <= r && r <= 0x06DC) ||
		(0x06DF <= r && r <= 0x06E4) ||
		(0x06E7 <= r && r <= 0x06E8) ||
		(0x06EA <= r && r <= 0x06ED) ||
		(0x0711 == r) ||
		(0x0730 <= r && r <= 0x074A) ||
		(0x07A6 <= r && r <= 0x07B0) ||
		(0x07EB <= r && r <= 0x07F3) ||
		(0x07FD == r) ||
		(0x0816 <= r && r <= 0x0819) ||
		(0x081B <= r && r <= 0x0823) ||
		(0x0825 <= r && r <= 0x0827) ||
		(0x0829 <= r && r <= 0x082D) ||
		(0x0859 <= r && r <= 0x085B) ||
		(0x08D3 <= r && r <= 0x08E1) ||
		(0x08E3 <= r && r <= 0x0902) ||
		(0x093A == r) ||
		(0x093C == r) ||
		(0x0941 <= r && r <= 0x0948) ||
		(0x094D == r) ||
		(0x0951 <= r && r <= 0x0957) ||
		(0x0962 <= r && r <= 0x0963) ||
		(0x0981 == r) ||
		(0x09BC == r) ||
		(0x09BE == r) ||
		(0x09C1 <= r && r <= 0x09C4) ||
		(0x09CD == r) ||
		(0x09D7 == r) ||
		(0x09E2 <= r && r <= 0x09E3) ||
		(0x09FE == r) ||
		(0x0A01 <= r && r <= 0x0A02) ||
		(0x0A3C == r) ||
		(0x0A41 <= r && r <= 0x0A42) ||
		(0x0A47 <= r && r <= 0x0A48) ||
		(0x0A4B <= r && r <= 0x0A4D) ||
		(0x0A51 == r) ||
		(0x0A70 <= r && r <= 0x0A71) ||
		(0x0A75 == r) ||
		(0x0A81 <= r && r <= 0x0A82) ||
		(0x0ABC == r) ||
		(0x0AC1 <= r && r <= 0x0AC5) ||
		(0x0AC7 <= r && r <= 0x0AC8) ||
		(0x0ACD == r) ||
		(0x0AE2 <= r && r <= 0x0AE3) ||
		(0x0AFA <= r && r <= 0x0AFF) ||
		(0x0B01 == r) ||
		(0x0B3C == r) ||
		(0x0B3E == r) ||
		(0x0B3F == r) ||
		(0x0B41 <= r && r <= 0x0B44) ||
		(0x0B4D == r) ||
		(0x0B55 <= r && r <= 0x0B56) ||
		(0x0B57 == r) ||
		(0x0B62 <= r && r <= 0x0B63) ||
		(0x0B82 == r) ||
		(0x0BBE == r) ||
		(0x0BC0 == r) ||
		(0x0BCD == r) ||
		(0x0BD7 == r) ||
		(0x0C00 == r) ||
		(0x0C04 == r) ||
		(0x0C3E <= r && r <= 0x0C40) ||
		(0x0C46 <= r && r <= 0x0C48) ||
		(0x0C4A <= r && r <= 0x0C4D) ||
		(0x0C55 <= r && r <= 0x0C56) ||
		(0x0C62 <= r && r <= 0x0C63) ||
		(0x0C81 == r) ||
		(0x0CBC == r) ||
		(0x0CBF == r) ||
		(0x0CC2 == r) ||
		(0x0CC6 == r) ||
		(0x0CCC <= r && r <= 0x0CCD) ||
		(0x0CD5 <= r && r <= 0x0CD6) ||
		(0x0CE2 <= r && r <= 0x0CE3) ||
		(0x0D00 <= r && r <= 0x0D01) ||
		(0x0D3B <= r && r <= 0x0D3C) ||
		(0x0D3E == r) ||
		(0x0D41 <= r && r <= 0x0D44) ||
		(0x0D4D == r) ||
		(0x0D57 == r) ||
		(0x0D62 <= r && r <= 0x0D63) ||
		(0x0D81 == r) ||
		(0x0DCA == r) ||
		(0x0DCF == r) ||
		(0x0DD2 <= r && r <= 0x0DD4) ||
		(0x0DD6 == r) ||
		(0x0DDF == r) ||
		(0x0E31 == r) ||
		(0x0E34 <= r && r <= 0x0E3A) ||
		(0x0E47 <= r && r <= 0x0E4E) ||
		(0x0EB1 == r) ||
		(0x0EB4 <= r && r <= 0x0EBC) ||
		(0x0EC8 <= r && r <= 0x0ECD) ||
		(0x0F18 <= r && r <= 0x0F19) ||
		(0x0F35 == r) ||
		(0x0F37 == r) ||
		(0x0F39 == r) ||
		(0x0F71 <= r && r <= 0x0F7E) ||
		(0x0F80 <= r && r <= 0x0F84) ||
		(0x0F86 <= r && r <= 0x0F87) ||
		(0x0F8D <= r && r <= 0x0F97) ||
		(0x0F99 <= r && r <= 0x0FBC) ||
		(0x0FC6 == r) ||
		(0x102D <= r && r <= 0x1030) ||
		(0x1032 <= r && r <= 0x1037) ||
		(0x1039 <= r && r <= 0x103A) ||
		(0x103D <= r && r <= 0x103E) ||
		(0x1058 <= r && r <= 0x1059) ||
		(0x105E <= r && r <= 0x1060) ||
		(0x1071 <= r && r <= 0x1074) ||
		(0x1082 == r) ||
		(0x1085 <= r && r <= 0x1086) ||
		(0x108D == r) ||
		(0x109D == r) ||
		(0x135D <= r && r <= 0x135F) ||
		(0x1712 <= r && r <= 0x1714) ||
		(0x1732 <= r && r <= 0x1734) ||
		(0x1752 <= r && r <= 0x1753) ||
		(0x1772 <= r && r <= 0x1773) ||
		(0x17B4 <= r && r <= 0x17B5) ||
		(0x17B7 <= r && r <= 0x17BD) ||
		(0x17C6 == r) ||
		(0x17C9 <= r && r <= 0x17D3) ||
		(0x17DD == r) ||
		(0x180B <= r && r <= 0x180D) ||
		(0x1885 <= r && r <= 0x1886) ||
		(0x18A9 == r) ||
		(0x1920 <= r && r <= 0x1922) ||
		(0x1927 <= r && r <= 0x1928) ||
		(0x1932 == r) ||
		(0x1939 <= r && r <= 0x193B) ||
		(0x1A17 <= r && r <= 0x1A18) ||
		(0x1A1B == r) ||
		(0x1A56 == r) ||
		(0x1A58 <= r && r <= 0x1A5E) ||
		(0x1A60 == r) ||
		(0x1A62 == r) ||
		(0x1A65 <= r && r <= 0x1A6C) ||
		(0x1A73 <= r && r <= 0x1A7C) ||
		(0x1A7F == r) ||
		(0x1AB0 <= r && r <= 0x1ABD) ||
		(0x1ABE == r) ||
		(0x1ABF <= r && r <= 0x1AC0) ||
		(0x1B00 <= r && r <= 0x1B03) ||
		(0x1B34 == r) ||
		(0x1B35 == r) ||
		(0x1B36 <= r && r <= 0x1B3A) ||
		(0x1B3C == r) ||
		(0x1B42 == r) ||
		(0x1B6B <= r && r <= 0x1B73) ||
		(0x1B80 <= r && r <= 0x1B81) ||
		(0x1BA2 <= r && r <= 0x1BA5) ||
		(0x1BA8 <= r && r <= 0x1BA9) ||
		(0x1BAB <= r && r <= 0x1BAD) ||
		(0x1BE6 == r) ||
		(0x1BE8 <= r && r <= 0x1BE9) ||
		(0x1BED == r) ||
		(0x1BEF <= r && r <= 0x1BF1) ||
		(0x1C2C <= r && r <= 0x1C33) ||
		(0x1C36 <= r && r <= 0x1C37) ||
		(0x1CD0 <= r && r <= 0x1CD2) ||
		(0x1CD4 <= r && r <= 0x1CE0) ||
		(0x1CE2 <= r && r <= 0x1CE8) ||
		(0x1CED == r) ||
		(0x1CF4 == r) ||
		(0x1CF8 <= r && r <= 0x1CF9) ||
		(0x1DC0 <= r && r <= 0x1DF9) ||
		(0x1DFB <= r && r <= 0x1DFF) ||
		(0x200C == r) ||
		(0x20D0 <= r && r <= 0x20DC) ||
		(0x20DD <= r && r <= 0x20E0) ||
		(0x20E1 == r) ||
		(0x20E2 <= r && r <= 0x20E4) ||
		(0x20E5 <= r && r <= 0x20F0) ||
		(0x2CEF <= r && r <= 0x2CF1) ||
		(0x2D7F == r) ||
		(0x2DE0 <= r && r <= 0x2DFF) ||
		(0x302A <= r && r <= 0x302D) ||
		(0x302E <= r && r <= 0x302F) ||
		(0x3099 <= r && r <= 0x309A) ||
		(0xA66F == r) ||
		(0xA670 <= r && r <= 0xA672) ||
		(0xA674 <= r && r <= 0xA67D) ||
		(0xA69E <= r && r <= 0xA69F) ||
		(0xA6F0 <= r && r <= 0xA6F1) ||
		(0xA802 == r) ||
		(0xA806 == r) ||
		(0xA80B == r) ||
		(0xA825 <= r && r <= 0xA826) ||
		(0xA82C == r) ||
		(0xA8C4 <= r && r <= 0xA8C5) ||
		(0xA8E0 <= r && r <= 0xA8F1) ||
		(0xA8FF == r) ||
		(0xA926 <= r && r <= 0xA92D) ||
		(0xA947 <= r && r <= 0xA951) ||
		(0xA980 <= r && r <= 0xA982) ||
		(0xA9B3 == r) ||
		(0xA9B6 <= r && r <= 0xA9B9) ||
		(0xA9BC <= r && r <= 0xA9BD) ||
		(0xA9E5 == r) ||
		(0xAA29 <= r && r <= 0xAA2E) ||
		(0xAA31 <= r && r <= 0xAA32) ||
		(0xAA35 <= r && r <= 0xAA36) ||
		(0xAA43 == r) ||
		(0xAA4C == r) ||
		(0xAA7C == r) ||
		(0xAAB0 == r) ||
		(0xAAB2 <= r && r <= 0xAAB4) ||
		(0xAAB7 <= r && r <= 0xAAB8) ||
		(0xAABE <= r && r <= 0xAABF) ||
		(0xAAC1 == r) ||
		(0xAAEC <= r && r <= 0xAAED) ||
		(0xAAF6 == r) ||
		(0xABE5 == r) ||
		(0xABE8 == r) ||
		(0xABED == r) ||
		(0xFB1E == r) ||
		(0xFE00 <= r && r <= 0xFE0F) ||
		(0xFE20 <= r && r <= 0xFE2F) ||
		(0xFF9E <= r && r <= 0xFF9F) ||
		(0x101FD == r) ||
		(0x102E0 == r) ||
		(0x10376 <= r && r <= 0x1037A) ||
		(0x10A01 <= r && r <= 0x10A03) ||
		(0x10A05 <= r && r <= 0x10A06) ||
		(0x10A0C <= r && r <= 0x10A0F) ||
		(0x10A38 <= r && r <= 0x10A3A) ||
		(0x10A3F == r) ||
		(0x10AE5 <= r && r <= 0x10AE6) ||
		(0x10D24 <= r && r <= 0x10D27) ||
		(0x10EAB <= r && r <= 0x10EAC) ||
		(0x10F46 <= r && r <= 0x10F50) ||
		(0x11001 == r) ||
		(0x11038 <= r && r <= 0x11046) ||
		(0x1107F <= r && r <= 0x11081) ||
		(0x110B3 <= r && r <= 0x110B6) ||
		(0x110B9 <= r && r <= 0x110BA) ||
		(0x11100 <= r && r <= 0x11102) ||
		(0x11127 <= r && r <= 0x1112B) ||
		(0x1112D <= r && r <= 0x11134) ||
		(0x11173 == r) ||
		(0x11180 <= r && r <= 0x11181) ||
		(0x111B6 <= r && r <= 0x111BE) ||
		(0x111C9 <= r && r <= 0x111CC) ||
		(0x111CF == r) ||
		(0x1122F <= r && r <= 0x11231) ||
		(0x11234 == r) ||
		(0x11236 <= r && r <= 0x11237) ||
		(0x1123E == r) ||
		(0x112DF == r) ||
		(0x112E3 <= r && r <= 0x112EA) ||
		(0x11300 <= r && r <= 0x11301) ||
		(0x1133B <= r && r <= 0x1133C) ||
		(0x1133E == r) ||
		(0x11340 == r) ||
		(0x11357 == r) ||
		(0x11366 <= r && r <= 0x1136C) ||
		(0x11370 <= r && r <= 0x11374) ||
		(0x11438 <= r && r <= 0x1143F) ||
		(0x11442 <= r && r <= 0x11444) ||
		(0x11446 == r) ||
		(0x1145E == r) ||
		(0x114B0 == r) ||
		(0x114B3 <= r && r <= 0x114B8) ||
		(0x114BA == r) ||
		(0x114BD == r) ||
		(0x114BF <= r && r <= 0x114C0) ||
		(0x114C2 <= r && r <= 0x114C3) ||
		(0x115AF == r) ||
		(0x115B2 <= r && r <= 0x115B5) ||
		(0x115BC <= r && r <= 0x115BD) ||
		(0x115BF <= r && r <= 0x115C0) ||
		(0x115DC <= r && r <= 0x115DD) ||
		(0x11633 <= r && r <= 0x1163A) ||
		(0x1163D == r) ||
		(0x1163F <= r && r <= 0x11640) ||
		(0x116AB == r) ||
		(0x116AD == r) ||
		(0x116B0 <= r && r <= 0x116B5) ||
		(0x116B7 == r) ||
		(0x1171D <= r && r <= 0x1171F) ||
		(0x11722 <= r && r <= 0x11725) ||
		(0x11727 <= r && r <= 0x1172B) ||
		(0x1182F <= r && r <= 0x11837) ||
		(0x11839 <= r && r <= 0x1183A) ||
		(0x11930 == r) ||
		(0x1193B <= r && r <= 0x1193C) ||
		(0x1193E == r) ||
		(0x11943 == r) ||
		(0x119D4 <= r && r <= 0x119D7) ||
		(0x119DA <= r && r <= 0x119DB) ||
		(0x119E0 == r) ||
		(0x11A01 <= r && r <= 0x11A0A) ||
		(0x11A33 <= r && r <= 0x11A38) ||
		(0x11A3B <= r && r <= 0x11A3E) ||
		(0x11A47 == r) ||
		(0x11A51 <= r && r <= 0x11A56) ||
		(0x11A59 <= r && r <= 0x11A5B) ||
		(0x11A8A <= r && r <= 0x11A96) ||
		(0x11A98 <= r && r <= 0x11A99) ||
		(0x11C30 <= r && r <= 0x11C36) ||
		(0x11C38 <= r && r <= 0x11C3D) ||
		(0x11C3F == r) ||
		(0x11C92 <= r && r <= 0x11CA7) ||
		(0x11CAA <= r && r <= 0x11CB0) ||
		(0x11CB2 <= r && r <= 0x11CB3) ||
		(0x11CB5 <= r && r <= 0x11CB6) ||
		(0x11D31 <= r && r <= 0x11D36) ||
		(0x11D3A == r) ||
		(0x11D3C <= r && r <= 0x11D3D) ||
		(0x11D3F <= r && r <= 0x11D45) ||
		(0x11D47 == r) ||
		(0x11D90 <= r && r <= 0x11D91) ||
		(0x11D95 == r) ||
		(0x11D97 == r) ||
		(0x11EF3 <= r && r <= 0x11EF4) ||
		(0x16AF0 <= r && r <= 0x16AF4) ||
		(0x16B30 <= r && r <= 0x16B36) ||
		(0x16F4F == r) ||
		(0x16F8F <= r && r <= 0x16F92) ||
		(0x16FE4 == r) ||
		(0x1BC9D <= r && r <= 0x1BC9E) ||
		(0x1D165 == r) ||
		(0x1D167 <= r && r <= 0x1D169) ||
		(0x1D16E <= r && r <= 0x1D172) ||
		(0x1D17B <= r && r <= 0x1D182) ||
		(0x1D185 <= r && r <= 0x1D18B) ||
		(0x1D1AA <= r && r <= 0x1D1AD) ||
		(0x1D242 <= r && r <= 0x1D244) ||
		(0x1DA00 <= r && r <= 0x1DA36) ||
		(0x1DA3B <= r && r <= 0x1DA6C) ||
		(0x1DA75 == r) ||
		(0x1DA84 == r) ||
		(0x1DA9B <= r && r <= 0x1DA9F) ||
		(0x1DAA1 <= r && r <= 0x1DAAF) ||
		(0x1E000 <= r && r <= 0x1E006) ||
		(0x1E008 <= r && r <= 0x1E018) ||
		(0x1E01B <= r && r <= 0x1E021) ||
		(0x1E023 <= r && r <= 0x1E024) ||
		(0x1E026 <= r && r <= 0x1E02A) ||
		(0x1E130 <= r && r <= 0x1E136) ||
		(0x1E2EC <= r && r <= 0x1E2EF) ||
		(0x1E8D0 <= r && r <= 0x1E8D6) ||
		(0x1E944 <= r && r <= 0x1E94A) ||
		(0x1F3FB <= r && r <= 0x1F3FF) ||
		(0xE0020 <= r && r <= 0xE007F) ||
		(0xE0100 <= r && r <= 0xE01EF)
}

// isCbRegionalIndicator returns whether the test of
// Grapheme_Cluster_Break = Regional_Indicator evaluates to true.
//
// Data for this function was obtained from
// http://www.unicode.org/Public/UCD/latest/ucd/auxiliary/GraphemeBreakProperty.txt
// on Feb 20, 2021, and is used under the terms of the Unicode License,
// available at the following address: https://www.unicode.org/license.html
func isCbRegionalIndicator(r rune) bool {
	return (0x1f1e6 <= r && r <= 0x1f1ff)
}

// isCbSpacingMark returns whether the test of Grapheme_Cluster_Break = SpacingMark
// evaluates to true.
//
// Data for this function was obtained from
// http://www.unicode.org/Public/UCD/latest/ucd/auxiliary/GraphemeBreakProperty.txt
// on Feb 20, 2021, and is used under the terms of the Unicode License,
// available at the following address: https://www.unicode.org/license.html
func isCbSpacingMark(r rune) bool {
	return (0x0903 == r) ||
		(0x093B == r) ||
		(0x093E <= r && r <= 0x0940) ||
		(0x0949 <= r && r <= 0x094C) ||
		(0x094E <= r && r <= 0x094F) ||
		(0x0982 <= r && r <= 0x0983) ||
		(0x09BF <= r && r <= 0x09C0) ||
		(0x09C7 <= r && r <= 0x09C8) ||
		(0x09CB <= r && r <= 0x09CC) ||
		(0x0A03 == r) ||
		(0x0A3E <= r && r <= 0x0A40) ||
		(0x0A83 == r) ||
		(0x0ABE <= r && r <= 0x0AC0) ||
		(0x0AC9 == r) ||
		(0x0ACB <= r && r <= 0x0ACC) ||
		(0x0B02 <= r && r <= 0x0B03) ||
		(0x0B40 == r) ||
		(0x0B47 <= r && r <= 0x0B48) ||
		(0x0B4B <= r && r <= 0x0B4C) ||
		(0x0BBF == r) ||
		(0x0BC1 <= r && r <= 0x0BC2) ||
		(0x0BC6 <= r && r <= 0x0BC8) ||
		(0x0BCA <= r && r <= 0x0BCC) ||
		(0x0C01 <= r && r <= 0x0C03) ||
		(0x0C41 <= r && r <= 0x0C44) ||
		(0x0C82 <= r && r <= 0x0C83) ||
		(0x0CBE == r) ||
		(0x0CC0 <= r && r <= 0x0CC1) ||
		(0x0CC3 <= r && r <= 0x0CC4) ||
		(0x0CC7 <= r && r <= 0x0CC8) ||
		(0x0CCA <= r && r <= 0x0CCB) ||
		(0x0D02 <= r && r <= 0x0D03) ||
		(0x0D3F <= r && r <= 0x0D40) ||
		(0x0D46 <= r && r <= 0x0D48) ||
		(0x0D4A <= r && r <= 0x0D4C) ||
		(0x0D82 <= r && r <= 0x0D83) ||
		(0x0DD0 <= r && r <= 0x0DD1) ||
		(0x0DD8 <= r && r <= 0x0DDE) ||
		(0x0DF2 <= r && r <= 0x0DF3) ||
		(0x0E33 == r) ||
		(0x0EB3 == r) ||
		(0x0F3E <= r && r <= 0x0F3F) ||
		(0x0F7F == r) ||
		(0x1031 == r) ||
		(0x103B <= r && r <= 0x103C) ||
		(0x1056 <= r && r <= 0x1057) ||
		(0x1084 == r) ||
		(0x17B6 == r) ||
		(0x17BE <= r && r <= 0x17C5) ||
		(0x17C7 <= r && r <= 0x17C8) ||
		(0x1923 <= r && r <= 0x1926) ||
		(0x1929 <= r && r <= 0x192B) ||
		(0x1930 <= r && r <= 0x1931) ||
		(0x1933 <= r && r <= 0x1938) ||
		(0x1A19 <= r && r <= 0x1A1A) ||
		(0x1A55 == r) ||
		(0x1A57 == r) ||
		(0x1A6D <= r && r <= 0x1A72) ||
		(0x1B04 == r) ||
		(0x1B3B == r) ||
		(0x1B3D <= r && r <= 0x1B41) ||
		(0x1B43 <= r && r <= 0x1B44) ||
		(0x1B82 == r) ||
		(0x1BA1 == r) ||
		(0x1BA6 <= r && r <= 0x1BA7) ||
		(0x1BAA == r) ||
		(0x1BE7 == r) ||
		(0x1BEA <= r && r <= 0x1BEC) ||
		(0x1BEE == r) ||
		(0x1BF2 <= r && r <= 0x1BF3) ||
		(0x1C24 <= r && r <= 0x1C2B) ||
		(0x1C34 <= r && r <= 0x1C35) ||
		(0x1CE1 == r) ||
		(0x1CF7 == r) ||
		(0xA823 <= r && r <= 0xA824) ||
		(0xA827 == r) ||
		(0xA880 <= r && r <= 0xA881) ||
		(0xA8B4 <= r && r <= 0xA8C3) ||
		(0xA952 <= r && r <= 0xA953) ||
		(0xA983 == r) ||
		(0xA9B4 <= r && r <= 0xA9B5) ||
		(0xA9BA <= r && r <= 0xA9BB) ||
		(0xA9BE <= r && r <= 0xA9C0) ||
		(0xAA2F <= r && r <= 0xAA30) ||
		(0xAA33 <= r && r <= 0xAA34) ||
		(0xAA4D == r) ||
		(0xAAEB == r) ||
		(0xAAEE <= r && r <= 0xAAEF) ||
		(0xAAF5 == r) ||
		(0xABE3 <= r && r <= 0xABE4) ||
		(0xABE6 <= r && r <= 0xABE7) ||
		(0xABE9 <= r && r <= 0xABEA) ||
		(0xABEC == r) ||
		(0x11000 == r) ||
		(0x11002 == r) ||
		(0x11082 == r) ||
		(0x110B0 <= r && r <= 0x110B2) ||
		(0x110B7 <= r && r <= 0x110B8) ||
		(0x1112C == r) ||
		(0x11145 <= r && r <= 0x11146) ||
		(0x11182 == r) ||
		(0x111B3 <= r && r <= 0x111B5) ||
		(0x111BF <= r && r <= 0x111C0) ||
		(0x111CE == r) ||
		(0x1122C <= r && r <= 0x1122E) ||
		(0x11232 <= r && r <= 0x11233) ||
		(0x11235 == r) ||
		(0x112E0 <= r && r <= 0x112E2) ||
		(0x11302 <= r && r <= 0x11303) ||
		(0x1133F == r) ||
		(0x11341 <= r && r <= 0x11344) ||
		(0x11347 <= r && r <= 0x11348) ||
		(0x1134B <= r && r <= 0x1134D) ||
		(0x11362 <= r && r <= 0x11363) ||
		(0x11435 <= r && r <= 0x11437) ||
		(0x11440 <= r && r <= 0x11441) ||
		(0x11445 == r) ||
		(0x114B1 <= r && r <= 0x114B2) ||
		(0x114B9 == r) ||
		(0x114BB <= r && r <= 0x114BC) ||
		(0x114BE == r) ||
		(0x114C1 == r) ||
		(0x115B0 <= r && r <= 0x115B1) ||
		(0x115B8 <= r && r <= 0x115BB) ||
		(0x115BE == r) ||
		(0x11630 <= r && r <= 0x11632) ||
		(0x1163B <= r && r <= 0x1163C) ||
		(0x1163E == r) ||
		(0x116AC == r) ||
		(0x116AE <= r && r <= 0x116AF) ||
		(0x116B6 == r) ||
		(0x11720 <= r && r <= 0x11721) ||
		(0x11726 == r) ||
		(0x1182C <= r && r <= 0x1182E) ||
		(0x11838 == r) ||
		(0x11931 <= r && r <= 0x11935) ||
		(0x11937 <= r && r <= 0x11938) ||
		(0x1193D == r) ||
		(0x11940 == r) ||
		(0x11942 == r) ||
		(0x119D1 <= r && r <= 0x119D3) ||
		(0x119DC <= r && r <= 0x119DF) ||
		(0x119E4 == r) ||
		(0x11A39 == r) ||
		(0x11A57 <= r && r <= 0x11A58) ||
		(0x11A97 == r) ||
		(0x11C2F == r) ||
		(0x11C3E == r) ||
		(0x11CA9 == r) ||
		(0x11CB1 == r) ||
		(0x11CB4 == r) ||
		(0x11D8A <= r && r <= 0x11D8E) ||
		(0x11D93 <= r && r <= 0x11D94) ||
		(0x11D96 == r) ||
		(0x11EF5 <= r && r <= 0x11EF6) ||
		(0x16F51 <= r && r <= 0x16F87) ||
		(0x16FF0 <= r && r <= 0x16FF1) ||
		(0x1D166 == r) ||
		(0x1D16D == r)
}

// isCbL returns whether the test of Grapheme_Cluster_Break = L
// evaluates to true.
//
// Data for this function was obtained from
// http://www.unicode.org/Public/UCD/latest/ucd/auxiliary/GraphemeBreakProperty.txt
// on Feb 20, 2021, and is used under the terms of the Unicode License,
// available at the following address: https://www.unicode.org/license.html
func isCbL(r rune) bool {
	return (0x1100 <= r && r <= 0x115F) ||
		(0xA960 <= r && r <= 0xA97C)
}

// isCbV returns whether the test of Grapheme_Cluster_Break = V
// evaluates to true.
//
// Data for this function was obtained from
// http://www.unicode.org/Public/UCD/latest/ucd/auxiliary/GraphemeBreakProperty.txt
// on Feb 20, 2021, and is used under the terms of the Unicode License,
// available at the following address: https://www.unicode.org/license.html
func isCbV(r rune) bool {
	return (0x1160 <= r && r <= 0x11A7) ||
		(0xD7B0 <= r && r <= 0xD7C6)
}

// isCbT returns whether the test of Grapheme_Cluster_Break = T
// evaluates to true.
//
// Data for this function was obtained from
// http://www.unicode.org/Public/UCD/latest/ucd/auxiliary/GraphemeBreakProperty.txt
// on Feb 20, 2021, and is used under the terms of the Unicode License,
// available at the following address: https://www.unicode.org/license.html
func isCbT(r rune) bool {
	return (0x11A8 <= r && r <= 0x11FF) ||
		(0xD7CB <= r && r <= 0xD7FB)
}

// isCbLV returns whether the test of Grapheme_Cluster_Break = LV
// evaluates to true.
//
// Data for this function was obtained from
// http://www.unicode.org/Public/UCD/latest/ucd/auxiliary/GraphemeBreakProperty.txt
// on Feb 20, 2021, and is used under the terms of the Unicode License,
// available at the following address: https://www.unicode.org/license.html
func isCbLV(r rune) bool {
	return (0xAC00 == r) ||
		(0xAC1C == r) ||
		(0xAC38 == r) ||
		(0xAC54 == r) ||
		(0xAC70 == r) ||
		(0xAC8C == r) ||
		(0xACA8 == r) ||
		(0xACC4 == r) ||
		(0xACE0 == r) ||
		(0xACFC == r) ||
		(0xAD18 == r) ||
		(0xAD34 == r) ||
		(0xAD50 == r) ||
		(0xAD6C == r) ||
		(0xAD88 == r) ||
		(0xADA4 == r) ||
		(0xADC0 == r) ||
		(0xADDC == r) ||
		(0xADF8 == r) ||
		(0xAE14 == r) ||
		(0xAE30 == r) ||
		(0xAE4C == r) ||
		(0xAE68 == r) ||
		(0xAE84 == r) ||
		(0xAEA0 == r) ||
		(0xAEBC == r) ||
		(0xAED8 == r) ||
		(0xAEF4 == r) ||
		(0xAF10 == r) ||
		(0xAF2C == r) ||
		(0xAF48 == r) ||
		(0xAF64 == r) ||
		(0xAF80 == r) ||
		(0xAF9C == r) ||
		(0xAFB8 == r) ||
		(0xAFD4 == r) ||
		(0xAFF0 == r) ||
		(0xB00C == r) ||
		(0xB028 == r) ||
		(0xB044 == r) ||
		(0xB060 == r) ||
		(0xB07C == r) ||
		(0xB098 == r) ||
		(0xB0B4 == r) ||
		(0xB0D0 == r) ||
		(0xB0EC == r) ||
		(0xB108 == r) ||
		(0xB124 == r) ||
		(0xB140 == r) ||
		(0xB15C == r) ||
		(0xB178 == r) ||
		(0xB194 == r) ||
		(0xB1B0 == r) ||
		(0xB1CC == r) ||
		(0xB1E8 == r) ||
		(0xB204 == r) ||
		(0xB220 == r) ||
		(0xB23C == r) ||
		(0xB258 == r) ||
		(0xB274 == r) ||
		(0xB290 == r) ||
		(0xB2AC == r) ||
		(0xB2C8 == r) ||
		(0xB2E4 == r) ||
		(0xB300 == r) ||
		(0xB31C == r) ||
		(0xB338 == r) ||
		(0xB354 == r) ||
		(0xB370 == r) ||
		(0xB38C == r) ||
		(0xB3A8 == r) ||
		(0xB3C4 == r) ||
		(0xB3E0 == r) ||
		(0xB3FC == r) ||
		(0xB418 == r) ||
		(0xB434 == r) ||
		(0xB450 == r) ||
		(0xB46C == r) ||
		(0xB488 == r) ||
		(0xB4A4 == r) ||
		(0xB4C0 == r) ||
		(0xB4DC == r) ||
		(0xB4F8 == r) ||
		(0xB514 == r) ||
		(0xB530 == r) ||
		(0xB54C == r) ||
		(0xB568 == r) ||
		(0xB584 == r) ||
		(0xB5A0 == r) ||
		(0xB5BC == r) ||
		(0xB5D8 == r) ||
		(0xB5F4 == r) ||
		(0xB610 == r) ||
		(0xB62C == r) ||
		(0xB648 == r) ||
		(0xB664 == r) ||
		(0xB680 == r) ||
		(0xB69C == r) ||
		(0xB6B8 == r) ||
		(0xB6D4 == r) ||
		(0xB6F0 == r) ||
		(0xB70C == r) ||
		(0xB728 == r) ||
		(0xB744 == r) ||
		(0xB760 == r) ||
		(0xB77C == r) ||
		(0xB798 == r) ||
		(0xB7B4 == r) ||
		(0xB7D0 == r) ||
		(0xB7EC == r) ||
		(0xB808 == r) ||
		(0xB824 == r) ||
		(0xB840 == r) ||
		(0xB85C == r) ||
		(0xB878 == r) ||
		(0xB894 == r) ||
		(0xB8B0 == r) ||
		(0xB8CC == r) ||
		(0xB8E8 == r) ||
		(0xB904 == r) ||
		(0xB920 == r) ||
		(0xB93C == r) ||
		(0xB958 == r) ||
		(0xB974 == r) ||
		(0xB990 == r) ||
		(0xB9AC == r) ||
		(0xB9C8 == r) ||
		(0xB9E4 == r) ||
		(0xBA00 == r) ||
		(0xBA1C == r) ||
		(0xBA38 == r) ||
		(0xBA54 == r) ||
		(0xBA70 == r) ||
		(0xBA8C == r) ||
		(0xBAA8 == r) ||
		(0xBAC4 == r) ||
		(0xBAE0 == r) ||
		(0xBAFC == r) ||
		(0xBB18 == r) ||
		(0xBB34 == r) ||
		(0xBB50 == r) ||
		(0xBB6C == r) ||
		(0xBB88 == r) ||
		(0xBBA4 == r) ||
		(0xBBC0 == r) ||
		(0xBBDC == r) ||
		(0xBBF8 == r) ||
		(0xBC14 == r) ||
		(0xBC30 == r) ||
		(0xBC4C == r) ||
		(0xBC68 == r) ||
		(0xBC84 == r) ||
		(0xBCA0 == r) ||
		(0xBCBC == r) ||
		(0xBCD8 == r) ||
		(0xBCF4 == r) ||
		(0xBD10 == r) ||
		(0xBD2C == r) ||
		(0xBD48 == r) ||
		(0xBD64 == r) ||
		(0xBD80 == r) ||
		(0xBD9C == r) ||
		(0xBDB8 == r) ||
		(0xBDD4 == r) ||
		(0xBDF0 == r) ||
		(0xBE0C == r) ||
		(0xBE28 == r) ||
		(0xBE44 == r) ||
		(0xBE60 == r) ||
		(0xBE7C == r) ||
		(0xBE98 == r) ||
		(0xBEB4 == r) ||
		(0xBED0 == r) ||
		(0xBEEC == r) ||
		(0xBF08 == r) ||
		(0xBF24 == r) ||
		(0xBF40 == r) ||
		(0xBF5C == r) ||
		(0xBF78 == r) ||
		(0xBF94 == r) ||
		(0xBFB0 == r) ||
		(0xBFCC == r) ||
		(0xBFE8 == r) ||
		(0xC004 == r) ||
		(0xC020 == r) ||
		(0xC03C == r) ||
		(0xC058 == r) ||
		(0xC074 == r) ||
		(0xC090 == r) ||
		(0xC0AC == r) ||
		(0xC0C8 == r) ||
		(0xC0E4 == r) ||
		(0xC100 == r) ||
		(0xC11C == r) ||
		(0xC138 == r) ||
		(0xC154 == r) ||
		(0xC170 == r) ||
		(0xC18C == r) ||
		(0xC1A8 == r) ||
		(0xC1C4 == r) ||
		(0xC1E0 == r) ||
		(0xC1FC == r) ||
		(0xC218 == r) ||
		(0xC234 == r) ||
		(0xC250 == r) ||
		(0xC26C == r) ||
		(0xC288 == r) ||
		(0xC2A4 == r) ||
		(0xC2C0 == r) ||
		(0xC2DC == r) ||
		(0xC2F8 == r) ||
		(0xC314 == r) ||
		(0xC330 == r) ||
		(0xC34C == r) ||
		(0xC368 == r) ||
		(0xC384 == r) ||
		(0xC3A0 == r) ||
		(0xC3BC == r) ||
		(0xC3D8 == r) ||
		(0xC3F4 == r) ||
		(0xC410 == r) ||
		(0xC42C == r) ||
		(0xC448 == r) ||
		(0xC464 == r) ||
		(0xC480 == r) ||
		(0xC49C == r) ||
		(0xC4B8 == r) ||
		(0xC4D4 == r) ||
		(0xC4F0 == r) ||
		(0xC50C == r) ||
		(0xC528 == r) ||
		(0xC544 == r) ||
		(0xC560 == r) ||
		(0xC57C == r) ||
		(0xC598 == r) ||
		(0xC5B4 == r) ||
		(0xC5D0 == r) ||
		(0xC5EC == r) ||
		(0xC608 == r) ||
		(0xC624 == r) ||
		(0xC640 == r) ||
		(0xC65C == r) ||
		(0xC678 == r) ||
		(0xC694 == r) ||
		(0xC6B0 == r) ||
		(0xC6CC == r) ||
		(0xC6E8 == r) ||
		(0xC704 == r) ||
		(0xC720 == r) ||
		(0xC73C == r) ||
		(0xC758 == r) ||
		(0xC774 == r) ||
		(0xC790 == r) ||
		(0xC7AC == r) ||
		(0xC7C8 == r) ||
		(0xC7E4 == r) ||
		(0xC800 == r) ||
		(0xC81C == r) ||
		(0xC838 == r) ||
		(0xC854 == r) ||
		(0xC870 == r) ||
		(0xC88C == r) ||
		(0xC8A8 == r) ||
		(0xC8C4 == r) ||
		(0xC8E0 == r) ||
		(0xC8FC == r) ||
		(0xC918 == r) ||
		(0xC934 == r) ||
		(0xC950 == r) ||
		(0xC96C == r) ||
		(0xC988 == r) ||
		(0xC9A4 == r) ||
		(0xC9C0 == r) ||
		(0xC9DC == r) ||
		(0xC9F8 == r) ||
		(0xCA14 == r) ||
		(0xCA30 == r) ||
		(0xCA4C == r) ||
		(0xCA68 == r) ||
		(0xCA84 == r) ||
		(0xCAA0 == r) ||
		(0xCABC == r) ||
		(0xCAD8 == r) ||
		(0xCAF4 == r) ||
		(0xCB10 == r) ||
		(0xCB2C == r) ||
		(0xCB48 == r) ||
		(0xCB64 == r) ||
		(0xCB80 == r) ||
		(0xCB9C == r) ||
		(0xCBB8 == r) ||
		(0xCBD4 == r) ||
		(0xCBF0 == r) ||
		(0xCC0C == r) ||
		(0xCC28 == r) ||
		(0xCC44 == r) ||
		(0xCC60 == r) ||
		(0xCC7C == r) ||
		(0xCC98 == r) ||
		(0xCCB4 == r) ||
		(0xCCD0 == r) ||
		(0xCCEC == r) ||
		(0xCD08 == r) ||
		(0xCD24 == r) ||
		(0xCD40 == r) ||
		(0xCD5C == r) ||
		(0xCD78 == r) ||
		(0xCD94 == r) ||
		(0xCDB0 == r) ||
		(0xCDCC == r) ||
		(0xCDE8 == r) ||
		(0xCE04 == r) ||
		(0xCE20 == r) ||
		(0xCE3C == r) ||
		(0xCE58 == r) ||
		(0xCE74 == r) ||
		(0xCE90 == r) ||
		(0xCEAC == r) ||
		(0xCEC8 == r) ||
		(0xCEE4 == r) ||
		(0xCF00 == r) ||
		(0xCF1C == r) ||
		(0xCF38 == r) ||
		(0xCF54 == r) ||
		(0xCF70 == r) ||
		(0xCF8C == r) ||
		(0xCFA8 == r) ||
		(0xCFC4 == r) ||
		(0xCFE0 == r) ||
		(0xCFFC == r) ||
		(0xD018 == r) ||
		(0xD034 == r) ||
		(0xD050 == r) ||
		(0xD06C == r) ||
		(0xD088 == r) ||
		(0xD0A4 == r) ||
		(0xD0C0 == r) ||
		(0xD0DC == r) ||
		(0xD0F8 == r) ||
		(0xD114 == r) ||
		(0xD130 == r) ||
		(0xD14C == r) ||
		(0xD168 == r) ||
		(0xD184 == r) ||
		(0xD1A0 == r) ||
		(0xD1BC == r) ||
		(0xD1D8 == r) ||
		(0xD1F4 == r) ||
		(0xD210 == r) ||
		(0xD22C == r) ||
		(0xD248 == r) ||
		(0xD264 == r) ||
		(0xD280 == r) ||
		(0xD29C == r) ||
		(0xD2B8 == r) ||
		(0xD2D4 == r) ||
		(0xD2F0 == r) ||
		(0xD30C == r) ||
		(0xD328 == r) ||
		(0xD344 == r) ||
		(0xD360 == r) ||
		(0xD37C == r) ||
		(0xD398 == r) ||
		(0xD3B4 == r) ||
		(0xD3D0 == r) ||
		(0xD3EC == r) ||
		(0xD408 == r) ||
		(0xD424 == r) ||
		(0xD440 == r) ||
		(0xD45C == r) ||
		(0xD478 == r) ||
		(0xD494 == r) ||
		(0xD4B0 == r) ||
		(0xD4CC == r) ||
		(0xD4E8 == r) ||
		(0xD504 == r) ||
		(0xD520 == r) ||
		(0xD53C == r) ||
		(0xD558 == r) ||
		(0xD574 == r) ||
		(0xD590 == r) ||
		(0xD5AC == r) ||
		(0xD5C8 == r) ||
		(0xD5E4 == r) ||
		(0xD600 == r) ||
		(0xD61C == r) ||
		(0xD638 == r) ||
		(0xD654 == r) ||
		(0xD670 == r) ||
		(0xD68C == r) ||
		(0xD6A8 == r) ||
		(0xD6C4 == r) ||
		(0xD6E0 == r) ||
		(0xD6FC == r) ||
		(0xD718 == r) ||
		(0xD734 == r) ||
		(0xD750 == r) ||
		(0xD76C == r) ||
		(0xD788 == r)
}

// isCbLVT returns whether the test of Grapheme_Cluster_Break = LVT
// evaluates to true.
//
// Data for this function was obtained from
// http://www.unicode.org/Public/UCD/latest/ucd/auxiliary/GraphemeBreakProperty.txt
// on Feb 20, 2021, and is used under the terms of the Unicode License,
// available at the following address: https://www.unicode.org/license.html
func isCbLVT(r rune) bool {
	return (0xAC01 <= r && r <= 0xAC1B) ||
		(0xAC1D <= r && r <= 0xAC37) ||
		(0xAC39 <= r && r <= 0xAC53) ||
		(0xAC55 <= r && r <= 0xAC6F) ||
		(0xAC71 <= r && r <= 0xAC8B) ||
		(0xAC8D <= r && r <= 0xACA7) ||
		(0xACA9 <= r && r <= 0xACC3) ||
		(0xACC5 <= r && r <= 0xACDF) ||
		(0xACE1 <= r && r <= 0xACFB) ||
		(0xACFD <= r && r <= 0xAD17) ||
		(0xAD19 <= r && r <= 0xAD33) ||
		(0xAD35 <= r && r <= 0xAD4F) ||
		(0xAD51 <= r && r <= 0xAD6B) ||
		(0xAD6D <= r && r <= 0xAD87) ||
		(0xAD89 <= r && r <= 0xADA3) ||
		(0xADA5 <= r && r <= 0xADBF) ||
		(0xADC1 <= r && r <= 0xADDB) ||
		(0xADDD <= r && r <= 0xADF7) ||
		(0xADF9 <= r && r <= 0xAE13) ||
		(0xAE15 <= r && r <= 0xAE2F) ||
		(0xAE31 <= r && r <= 0xAE4B) ||
		(0xAE4D <= r && r <= 0xAE67) ||
		(0xAE69 <= r && r <= 0xAE83) ||
		(0xAE85 <= r && r <= 0xAE9F) ||
		(0xAEA1 <= r && r <= 0xAEBB) ||
		(0xAEBD <= r && r <= 0xAED7) ||
		(0xAED9 <= r && r <= 0xAEF3) ||
		(0xAEF5 <= r && r <= 0xAF0F) ||
		(0xAF11 <= r && r <= 0xAF2B) ||
		(0xAF2D <= r && r <= 0xAF47) ||
		(0xAF49 <= r && r <= 0xAF63) ||
		(0xAF65 <= r && r <= 0xAF7F) ||
		(0xAF81 <= r && r <= 0xAF9B) ||
		(0xAF9D <= r && r <= 0xAFB7) ||
		(0xAFB9 <= r && r <= 0xAFD3) ||
		(0xAFD5 <= r && r <= 0xAFEF) ||
		(0xAFF1 <= r && r <= 0xB00B) ||
		(0xB00D <= r && r <= 0xB027) ||
		(0xB029 <= r && r <= 0xB043) ||
		(0xB045 <= r && r <= 0xB05F) ||
		(0xB061 <= r && r <= 0xB07B) ||
		(0xB07D <= r && r <= 0xB097) ||
		(0xB099 <= r && r <= 0xB0B3) ||
		(0xB0B5 <= r && r <= 0xB0CF) ||
		(0xB0D1 <= r && r <= 0xB0EB) ||
		(0xB0ED <= r && r <= 0xB107) ||
		(0xB109 <= r && r <= 0xB123) ||
		(0xB125 <= r && r <= 0xB13F) ||
		(0xB141 <= r && r <= 0xB15B) ||
		(0xB15D <= r && r <= 0xB177) ||
		(0xB179 <= r && r <= 0xB193) ||
		(0xB195 <= r && r <= 0xB1AF) ||
		(0xB1B1 <= r && r <= 0xB1CB) ||
		(0xB1CD <= r && r <= 0xB1E7) ||
		(0xB1E9 <= r && r <= 0xB203) ||
		(0xB205 <= r && r <= 0xB21F) ||
		(0xB221 <= r && r <= 0xB23B) ||
		(0xB23D <= r && r <= 0xB257) ||
		(0xB259 <= r && r <= 0xB273) ||
		(0xB275 <= r && r <= 0xB28F) ||
		(0xB291 <= r && r <= 0xB2AB) ||
		(0xB2AD <= r && r <= 0xB2C7) ||
		(0xB2C9 <= r && r <= 0xB2E3) ||
		(0xB2E5 <= r && r <= 0xB2FF) ||
		(0xB301 <= r && r <= 0xB31B) ||
		(0xB31D <= r && r <= 0xB337) ||
		(0xB339 <= r && r <= 0xB353) ||
		(0xB355 <= r && r <= 0xB36F) ||
		(0xB371 <= r && r <= 0xB38B) ||
		(0xB38D <= r && r <= 0xB3A7) ||
		(0xB3A9 <= r && r <= 0xB3C3) ||
		(0xB3C5 <= r && r <= 0xB3DF) ||
		(0xB3E1 <= r && r <= 0xB3FB) ||
		(0xB3FD <= r && r <= 0xB417) ||
		(0xB419 <= r && r <= 0xB433) ||
		(0xB435 <= r && r <= 0xB44F) ||
		(0xB451 <= r && r <= 0xB46B) ||
		(0xB46D <= r && r <= 0xB487) ||
		(0xB489 <= r && r <= 0xB4A3) ||
		(0xB4A5 <= r && r <= 0xB4BF) ||
		(0xB4C1 <= r && r <= 0xB4DB) ||
		(0xB4DD <= r && r <= 0xB4F7) ||
		(0xB4F9 <= r && r <= 0xB513) ||
		(0xB515 <= r && r <= 0xB52F) ||
		(0xB531 <= r && r <= 0xB54B) ||
		(0xB54D <= r && r <= 0xB567) ||
		(0xB569 <= r && r <= 0xB583) ||
		(0xB585 <= r && r <= 0xB59F) ||
		(0xB5A1 <= r && r <= 0xB5BB) ||
		(0xB5BD <= r && r <= 0xB5D7) ||
		(0xB5D9 <= r && r <= 0xB5F3) ||
		(0xB5F5 <= r && r <= 0xB60F) ||
		(0xB611 <= r && r <= 0xB62B) ||
		(0xB62D <= r && r <= 0xB647) ||
		(0xB649 <= r && r <= 0xB663) ||
		(0xB665 <= r && r <= 0xB67F) ||
		(0xB681 <= r && r <= 0xB69B) ||
		(0xB69D <= r && r <= 0xB6B7) ||
		(0xB6B9 <= r && r <= 0xB6D3) ||
		(0xB6D5 <= r && r <= 0xB6EF) ||
		(0xB6F1 <= r && r <= 0xB70B) ||
		(0xB70D <= r && r <= 0xB727) ||
		(0xB729 <= r && r <= 0xB743) ||
		(0xB745 <= r && r <= 0xB75F) ||
		(0xB761 <= r && r <= 0xB77B) ||
		(0xB77D <= r && r <= 0xB797) ||
		(0xB799 <= r && r <= 0xB7B3) ||
		(0xB7B5 <= r && r <= 0xB7CF) ||
		(0xB7D1 <= r && r <= 0xB7EB) ||
		(0xB7ED <= r && r <= 0xB807) ||
		(0xB809 <= r && r <= 0xB823) ||
		(0xB825 <= r && r <= 0xB83F) ||
		(0xB841 <= r && r <= 0xB85B) ||
		(0xB85D <= r && r <= 0xB877) ||
		(0xB879 <= r && r <= 0xB893) ||
		(0xB895 <= r && r <= 0xB8AF) ||
		(0xB8B1 <= r && r <= 0xB8CB) ||
		(0xB8CD <= r && r <= 0xB8E7) ||
		(0xB8E9 <= r && r <= 0xB903) ||
		(0xB905 <= r && r <= 0xB91F) ||
		(0xB921 <= r && r <= 0xB93B) ||
		(0xB93D <= r && r <= 0xB957) ||
		(0xB959 <= r && r <= 0xB973) ||
		(0xB975 <= r && r <= 0xB98F) ||
		(0xB991 <= r && r <= 0xB9AB) ||
		(0xB9AD <= r && r <= 0xB9C7) ||
		(0xB9C9 <= r && r <= 0xB9E3) ||
		(0xB9E5 <= r && r <= 0xB9FF) ||
		(0xBA01 <= r && r <= 0xBA1B) ||
		(0xBA1D <= r && r <= 0xBA37) ||
		(0xBA39 <= r && r <= 0xBA53) ||
		(0xBA55 <= r && r <= 0xBA6F) ||
		(0xBA71 <= r && r <= 0xBA8B) ||
		(0xBA8D <= r && r <= 0xBAA7) ||
		(0xBAA9 <= r && r <= 0xBAC3) ||
		(0xBAC5 <= r && r <= 0xBADF) ||
		(0xBAE1 <= r && r <= 0xBAFB) ||
		(0xBAFD <= r && r <= 0xBB17) ||
		(0xBB19 <= r && r <= 0xBB33) ||
		(0xBB35 <= r && r <= 0xBB4F) ||
		(0xBB51 <= r && r <= 0xBB6B) ||
		(0xBB6D <= r && r <= 0xBB87) ||
		(0xBB89 <= r && r <= 0xBBA3) ||
		(0xBBA5 <= r && r <= 0xBBBF) ||
		(0xBBC1 <= r && r <= 0xBBDB) ||
		(0xBBDD <= r && r <= 0xBBF7) ||
		(0xBBF9 <= r && r <= 0xBC13) ||
		(0xBC15 <= r && r <= 0xBC2F) ||
		(0xBC31 <= r && r <= 0xBC4B) ||
		(0xBC4D <= r && r <= 0xBC67) ||
		(0xBC69 <= r && r <= 0xBC83) ||
		(0xBC85 <= r && r <= 0xBC9F) ||
		(0xBCA1 <= r && r <= 0xBCBB) ||
		(0xBCBD <= r && r <= 0xBCD7) ||
		(0xBCD9 <= r && r <= 0xBCF3) ||
		(0xBCF5 <= r && r <= 0xBD0F) ||
		(0xBD11 <= r && r <= 0xBD2B) ||
		(0xBD2D <= r && r <= 0xBD47) ||
		(0xBD49 <= r && r <= 0xBD63) ||
		(0xBD65 <= r && r <= 0xBD7F) ||
		(0xBD81 <= r && r <= 0xBD9B) ||
		(0xBD9D <= r && r <= 0xBDB7) ||
		(0xBDB9 <= r && r <= 0xBDD3) ||
		(0xBDD5 <= r && r <= 0xBDEF) ||
		(0xBDF1 <= r && r <= 0xBE0B) ||
		(0xBE0D <= r && r <= 0xBE27) ||
		(0xBE29 <= r && r <= 0xBE43) ||
		(0xBE45 <= r && r <= 0xBE5F) ||
		(0xBE61 <= r && r <= 0xBE7B) ||
		(0xBE7D <= r && r <= 0xBE97) ||
		(0xBE99 <= r && r <= 0xBEB3) ||
		(0xBEB5 <= r && r <= 0xBECF) ||
		(0xBED1 <= r && r <= 0xBEEB) ||
		(0xBEED <= r && r <= 0xBF07) ||
		(0xBF09 <= r && r <= 0xBF23) ||
		(0xBF25 <= r && r <= 0xBF3F) ||
		(0xBF41 <= r && r <= 0xBF5B) ||
		(0xBF5D <= r && r <= 0xBF77) ||
		(0xBF79 <= r && r <= 0xBF93) ||
		(0xBF95 <= r && r <= 0xBFAF) ||
		(0xBFB1 <= r && r <= 0xBFCB) ||
		(0xBFCD <= r && r <= 0xBFE7) ||
		(0xBFE9 <= r && r <= 0xC003) ||
		(0xC005 <= r && r <= 0xC01F) ||
		(0xC021 <= r && r <= 0xC03B) ||
		(0xC03D <= r && r <= 0xC057) ||
		(0xC059 <= r && r <= 0xC073) ||
		(0xC075 <= r && r <= 0xC08F) ||
		(0xC091 <= r && r <= 0xC0AB) ||
		(0xC0AD <= r && r <= 0xC0C7) ||
		(0xC0C9 <= r && r <= 0xC0E3) ||
		(0xC0E5 <= r && r <= 0xC0FF) ||
		(0xC101 <= r && r <= 0xC11B) ||
		(0xC11D <= r && r <= 0xC137) ||
		(0xC139 <= r && r <= 0xC153) ||
		(0xC155 <= r && r <= 0xC16F) ||
		(0xC171 <= r && r <= 0xC18B) ||
		(0xC18D <= r && r <= 0xC1A7) ||
		(0xC1A9 <= r && r <= 0xC1C3) ||
		(0xC1C5 <= r && r <= 0xC1DF) ||
		(0xC1E1 <= r && r <= 0xC1FB) ||
		(0xC1FD <= r && r <= 0xC217) ||
		(0xC219 <= r && r <= 0xC233) ||
		(0xC235 <= r && r <= 0xC24F) ||
		(0xC251 <= r && r <= 0xC26B) ||
		(0xC26D <= r && r <= 0xC287) ||
		(0xC289 <= r && r <= 0xC2A3) ||
		(0xC2A5 <= r && r <= 0xC2BF) ||
		(0xC2C1 <= r && r <= 0xC2DB) ||
		(0xC2DD <= r && r <= 0xC2F7) ||
		(0xC2F9 <= r && r <= 0xC313) ||
		(0xC315 <= r && r <= 0xC32F) ||
		(0xC331 <= r && r <= 0xC34B) ||
		(0xC34D <= r && r <= 0xC367) ||
		(0xC369 <= r && r <= 0xC383) ||
		(0xC385 <= r && r <= 0xC39F) ||
		(0xC3A1 <= r && r <= 0xC3BB) ||
		(0xC3BD <= r && r <= 0xC3D7) ||
		(0xC3D9 <= r && r <= 0xC3F3) ||
		(0xC3F5 <= r && r <= 0xC40F) ||
		(0xC411 <= r && r <= 0xC42B) ||
		(0xC42D <= r && r <= 0xC447) ||
		(0xC449 <= r && r <= 0xC463) ||
		(0xC465 <= r && r <= 0xC47F) ||
		(0xC481 <= r && r <= 0xC49B) ||
		(0xC49D <= r && r <= 0xC4B7) ||
		(0xC4B9 <= r && r <= 0xC4D3) ||
		(0xC4D5 <= r && r <= 0xC4EF) ||
		(0xC4F1 <= r && r <= 0xC50B) ||
		(0xC50D <= r && r <= 0xC527) ||
		(0xC529 <= r && r <= 0xC543) ||
		(0xC545 <= r && r <= 0xC55F) ||
		(0xC561 <= r && r <= 0xC57B) ||
		(0xC57D <= r && r <= 0xC597) ||
		(0xC599 <= r && r <= 0xC5B3) ||
		(0xC5B5 <= r && r <= 0xC5CF) ||
		(0xC5D1 <= r && r <= 0xC5EB) ||
		(0xC5ED <= r && r <= 0xC607) ||
		(0xC609 <= r && r <= 0xC623) ||
		(0xC625 <= r && r <= 0xC63F) ||
		(0xC641 <= r && r <= 0xC65B) ||
		(0xC65D <= r && r <= 0xC677) ||
		(0xC679 <= r && r <= 0xC693) ||
		(0xC695 <= r && r <= 0xC6AF) ||
		(0xC6B1 <= r && r <= 0xC6CB) ||
		(0xC6CD <= r && r <= 0xC6E7) ||
		(0xC6E9 <= r && r <= 0xC703) ||
		(0xC705 <= r && r <= 0xC71F) ||
		(0xC721 <= r && r <= 0xC73B) ||
		(0xC73D <= r && r <= 0xC757) ||
		(0xC759 <= r && r <= 0xC773) ||
		(0xC775 <= r && r <= 0xC78F) ||
		(0xC791 <= r && r <= 0xC7AB) ||
		(0xC7AD <= r && r <= 0xC7C7) ||
		(0xC7C9 <= r && r <= 0xC7E3) ||
		(0xC7E5 <= r && r <= 0xC7FF) ||
		(0xC801 <= r && r <= 0xC81B) ||
		(0xC81D <= r && r <= 0xC837) ||
		(0xC839 <= r && r <= 0xC853) ||
		(0xC855 <= r && r <= 0xC86F) ||
		(0xC871 <= r && r <= 0xC88B) ||
		(0xC88D <= r && r <= 0xC8A7) ||
		(0xC8A9 <= r && r <= 0xC8C3) ||
		(0xC8C5 <= r && r <= 0xC8DF) ||
		(0xC8E1 <= r && r <= 0xC8FB) ||
		(0xC8FD <= r && r <= 0xC917) ||
		(0xC919 <= r && r <= 0xC933) ||
		(0xC935 <= r && r <= 0xC94F) ||
		(0xC951 <= r && r <= 0xC96B) ||
		(0xC96D <= r && r <= 0xC987) ||
		(0xC989 <= r && r <= 0xC9A3) ||
		(0xC9A5 <= r && r <= 0xC9BF) ||
		(0xC9C1 <= r && r <= 0xC9DB) ||
		(0xC9DD <= r && r <= 0xC9F7) ||
		(0xC9F9 <= r && r <= 0xCA13) ||
		(0xCA15 <= r && r <= 0xCA2F) ||
		(0xCA31 <= r && r <= 0xCA4B) ||
		(0xCA4D <= r && r <= 0xCA67) ||
		(0xCA69 <= r && r <= 0xCA83) ||
		(0xCA85 <= r && r <= 0xCA9F) ||
		(0xCAA1 <= r && r <= 0xCABB) ||
		(0xCABD <= r && r <= 0xCAD7) ||
		(0xCAD9 <= r && r <= 0xCAF3) ||
		(0xCAF5 <= r && r <= 0xCB0F) ||
		(0xCB11 <= r && r <= 0xCB2B) ||
		(0xCB2D <= r && r <= 0xCB47) ||
		(0xCB49 <= r && r <= 0xCB63) ||
		(0xCB65 <= r && r <= 0xCB7F) ||
		(0xCB81 <= r && r <= 0xCB9B) ||
		(0xCB9D <= r && r <= 0xCBB7) ||
		(0xCBB9 <= r && r <= 0xCBD3) ||
		(0xCBD5 <= r && r <= 0xCBEF) ||
		(0xCBF1 <= r && r <= 0xCC0B) ||
		(0xCC0D <= r && r <= 0xCC27) ||
		(0xCC29 <= r && r <= 0xCC43) ||
		(0xCC45 <= r && r <= 0xCC5F) ||
		(0xCC61 <= r && r <= 0xCC7B) ||
		(0xCC7D <= r && r <= 0xCC97) ||
		(0xCC99 <= r && r <= 0xCCB3) ||
		(0xCCB5 <= r && r <= 0xCCCF) ||
		(0xCCD1 <= r && r <= 0xCCEB) ||
		(0xCCED <= r && r <= 0xCD07) ||
		(0xCD09 <= r && r <= 0xCD23) ||
		(0xCD25 <= r && r <= 0xCD3F) ||
		(0xCD41 <= r && r <= 0xCD5B) ||
		(0xCD5D <= r && r <= 0xCD77) ||
		(0xCD79 <= r && r <= 0xCD93) ||
		(0xCD95 <= r && r <= 0xCDAF) ||
		(0xCDB1 <= r && r <= 0xCDCB) ||
		(0xCDCD <= r && r <= 0xCDE7) ||
		(0xCDE9 <= r && r <= 0xCE03) ||
		(0xCE05 <= r && r <= 0xCE1F) ||
		(0xCE21 <= r && r <= 0xCE3B) ||
		(0xCE3D <= r && r <= 0xCE57) ||
		(0xCE59 <= r && r <= 0xCE73) ||
		(0xCE75 <= r && r <= 0xCE8F) ||
		(0xCE91 <= r && r <= 0xCEAB) ||
		(0xCEAD <= r && r <= 0xCEC7) ||
		(0xCEC9 <= r && r <= 0xCEE3) ||
		(0xCEE5 <= r && r <= 0xCEFF) ||
		(0xCF01 <= r && r <= 0xCF1B) ||
		(0xCF1D <= r && r <= 0xCF37) ||
		(0xCF39 <= r && r <= 0xCF53) ||
		(0xCF55 <= r && r <= 0xCF6F) ||
		(0xCF71 <= r && r <= 0xCF8B) ||
		(0xCF8D <= r && r <= 0xCFA7) ||
		(0xCFA9 <= r && r <= 0xCFC3) ||
		(0xCFC5 <= r && r <= 0xCFDF) ||
		(0xCFE1 <= r && r <= 0xCFFB) ||
		(0xCFFD <= r && r <= 0xD017) ||
		(0xD019 <= r && r <= 0xD033) ||
		(0xD035 <= r && r <= 0xD04F) ||
		(0xD051 <= r && r <= 0xD06B) ||
		(0xD06D <= r && r <= 0xD087) ||
		(0xD089 <= r && r <= 0xD0A3) ||
		(0xD0A5 <= r && r <= 0xD0BF) ||
		(0xD0C1 <= r && r <= 0xD0DB) ||
		(0xD0DD <= r && r <= 0xD0F7) ||
		(0xD0F9 <= r && r <= 0xD113) ||
		(0xD115 <= r && r <= 0xD12F) ||
		(0xD131 <= r && r <= 0xD14B) ||
		(0xD14D <= r && r <= 0xD167) ||
		(0xD169 <= r && r <= 0xD183) ||
		(0xD185 <= r && r <= 0xD19F) ||
		(0xD1A1 <= r && r <= 0xD1BB) ||
		(0xD1BD <= r && r <= 0xD1D7) ||
		(0xD1D9 <= r && r <= 0xD1F3) ||
		(0xD1F5 <= r && r <= 0xD20F) ||
		(0xD211 <= r && r <= 0xD22B) ||
		(0xD22D <= r && r <= 0xD247) ||
		(0xD249 <= r && r <= 0xD263) ||
		(0xD265 <= r && r <= 0xD27F) ||
		(0xD281 <= r && r <= 0xD29B) ||
		(0xD29D <= r && r <= 0xD2B7) ||
		(0xD2B9 <= r && r <= 0xD2D3) ||
		(0xD2D5 <= r && r <= 0xD2EF) ||
		(0xD2F1 <= r && r <= 0xD30B) ||
		(0xD30D <= r && r <= 0xD327) ||
		(0xD329 <= r && r <= 0xD343) ||
		(0xD345 <= r && r <= 0xD35F) ||
		(0xD361 <= r && r <= 0xD37B) ||
		(0xD37D <= r && r <= 0xD397) ||
		(0xD399 <= r && r <= 0xD3B3) ||
		(0xD3B5 <= r && r <= 0xD3CF) ||
		(0xD3D1 <= r && r <= 0xD3EB) ||
		(0xD3ED <= r && r <= 0xD407) ||
		(0xD409 <= r && r <= 0xD423) ||
		(0xD425 <= r && r <= 0xD43F) ||
		(0xD441 <= r && r <= 0xD45B) ||
		(0xD45D <= r && r <= 0xD477) ||
		(0xD479 <= r && r <= 0xD493) ||
		(0xD495 <= r && r <= 0xD4AF) ||
		(0xD4B1 <= r && r <= 0xD4CB) ||
		(0xD4CD <= r && r <= 0xD4E7) ||
		(0xD4E9 <= r && r <= 0xD503) ||
		(0xD505 <= r && r <= 0xD51F) ||
		(0xD521 <= r && r <= 0xD53B) ||
		(0xD53D <= r && r <= 0xD557) ||
		(0xD559 <= r && r <= 0xD573) ||
		(0xD575 <= r && r <= 0xD58F) ||
		(0xD591 <= r && r <= 0xD5AB) ||
		(0xD5AD <= r && r <= 0xD5C7) ||
		(0xD5C9 <= r && r <= 0xD5E3) ||
		(0xD5E5 <= r && r <= 0xD5FF) ||
		(0xD601 <= r && r <= 0xD61B) ||
		(0xD61D <= r && r <= 0xD637) ||
		(0xD639 <= r && r <= 0xD653) ||
		(0xD655 <= r && r <= 0xD66F) ||
		(0xD671 <= r && r <= 0xD68B) ||
		(0xD68D <= r && r <= 0xD6A7) ||
		(0xD6A9 <= r && r <= 0xD6C3) ||
		(0xD6C5 <= r && r <= 0xD6DF) ||
		(0xD6E1 <= r && r <= 0xD6FB) ||
		(0xD6FD <= r && r <= 0xD717) ||
		(0xD719 <= r && r <= 0xD733) ||
		(0xD735 <= r && r <= 0xD74F) ||
		(0xD751 <= r && r <= 0xD76B) ||
		(0xD76D <= r && r <= 0xD787) ||
		(0xD789 <= r && r <= 0xD7A3)
}

// isCbZWJ returns whether the test of Grapheme_Cluster_Break = ZWJ
// evaluates to true.
//
// Data for this function was obtained from
// http://www.unicode.org/Public/UCD/latest/ucd/auxiliary/GraphemeBreakProperty.txt
// on Feb 20, 2021, and is used under the terms of the Unicode License,
// available at the following address: https://www.unicode.org/license.html
func isCbZWJ(r rune) bool {
	return (0x200D == r)
}

// isExtPicot Checks if r has property Extended_Pictographic = Yes.
//
// Data for this function was obtained from
// https://unicode.org/Public/UNIDATA/emoji/emoji-data.txt on Feb 20, 2021, and
// is used under the terms of the Unicode License, available at the
// following address: https://www.unicode.org/license.html
func isExtPicto(r rune) bool {
	return (0x00A9 == r) ||
		(0x00AE == r) ||
		(0x203C == r) ||
		(0x2049 == r) ||
		(0x2122 == r) ||
		(0x2139 == r) ||
		(0x2194 <= r && r <= 0x2199) ||
		(0x21A9 <= r && r <= 0x21AA) ||
		(0x231A <= r && r <= 0x231B) ||
		(0x2328 == r) ||
		(0x2388 == r) ||
		(0x23CF == r) ||
		(0x23E9 <= r && r <= 0x23EC) ||
		(0x23ED <= r && r <= 0x23EE) ||
		(0x23EF == r) ||
		(0x23F0 == r) ||
		(0x23F1 <= r && r <= 0x23F2) ||
		(0x23F3 == r) ||
		(0x23F8 <= r && r <= 0x23FA) ||
		(0x24C2 == r) ||
		(0x25AA <= r && r <= 0x25AB) ||
		(0x25B6 == r) ||
		(0x25C0 == r) ||
		(0x25FB <= r && r <= 0x25FE) ||
		(0x2600 <= r && r <= 0x2601) ||
		(0x2602 <= r && r <= 0x2603) ||
		(0x2604 == r) ||
		(0x2605 == r) ||
		(0x2607 <= r && r <= 0x260D) ||
		(0x260E == r) ||
		(0x260F <= r && r <= 0x2610) ||
		(0x2611 == r) ||
		(0x2612 == r) ||
		(0x2614 <= r && r <= 0x2615) ||
		(0x2616 <= r && r <= 0x2617) ||
		(0x2618 == r) ||
		(0x2619 <= r && r <= 0x261C) ||
		(0x261D == r) ||
		(0x261E <= r && r <= 0x261F) ||
		(0x2620 == r) ||
		(0x2621 == r) ||
		(0x2622 <= r && r <= 0x2623) ||
		(0x2624 <= r && r <= 0x2625) ||
		(0x2626 == r) ||
		(0x2627 <= r && r <= 0x2629) ||
		(0x262A == r) ||
		(0x262B <= r && r <= 0x262D) ||
		(0x262E == r) ||
		(0x262F == r) ||
		(0x2630 <= r && r <= 0x2637) ||
		(0x2638 <= r && r <= 0x2639) ||
		(0x263A == r) ||
		(0x263B <= r && r <= 0x263F) ||
		(0x2640 == r) ||
		(0x2641 == r) ||
		(0x2642 == r) ||
		(0x2643 <= r && r <= 0x2647) ||
		(0x2648 <= r && r <= 0x2653) ||
		(0x2654 <= r && r <= 0x265E) ||
		(0x265F == r) ||
		(0x2660 == r) ||
		(0x2661 <= r && r <= 0x2662) ||
		(0x2663 == r) ||
		(0x2664 == r) ||
		(0x2665 <= r && r <= 0x2666) ||
		(0x2667 == r) ||
		(0x2668 == r) ||
		(0x2669 <= r && r <= 0x267A) ||
		(0x267B == r) ||
		(0x267C <= r && r <= 0x267D) ||
		(0x267E == r) ||
		(0x267F == r) ||
		(0x2680 <= r && r <= 0x2685) ||
		(0x2690 <= r && r <= 0x2691) ||
		(0x2692 == r) ||
		(0x2693 == r) ||
		(0x2694 == r) ||
		(0x2695 == r) ||
		(0x2696 <= r && r <= 0x2697) ||
		(0x2698 == r) ||
		(0x2699 == r) ||
		(0x269A == r) ||
		(0x269B <= r && r <= 0x269C) ||
		(0x269D <= r && r <= 0x269F) ||
		(0x26A0 <= r && r <= 0x26A1) ||
		(0x26A2 <= r && r <= 0x26A6) ||
		(0x26A7 == r) ||
		(0x26A8 <= r && r <= 0x26A9) ||
		(0x26AA <= r && r <= 0x26AB) ||
		(0x26AC <= r && r <= 0x26AF) ||
		(0x26B0 <= r && r <= 0x26B1) ||
		(0x26B2 <= r && r <= 0x26BC) ||
		(0x26BD <= r && r <= 0x26BE) ||
		(0x26BF <= r && r <= 0x26C3) ||
		(0x26C4 <= r && r <= 0x26C5) ||
		(0x26C6 <= r && r <= 0x26C7) ||
		(0x26C8 == r) ||
		(0x26C9 <= r && r <= 0x26CD) ||
		(0x26CE == r) ||
		(0x26CF == r) ||
		(0x26D0 == r) ||
		(0x26D1 == r) ||
		(0x26D2 == r) ||
		(0x26D3 == r) ||
		(0x26D4 == r) ||
		(0x26D5 <= r && r <= 0x26E8) ||
		(0x26E9 == r) ||
		(0x26EA == r) ||
		(0x26EB <= r && r <= 0x26EF) ||
		(0x26F0 <= r && r <= 0x26F1) ||
		(0x26F2 <= r && r <= 0x26F3) ||
		(0x26F4 == r) ||
		(0x26F5 == r) ||
		(0x26F6 == r) ||
		(0x26F7 <= r && r <= 0x26F9) ||
		(0x26FA == r) ||
		(0x26FB <= r && r <= 0x26FC) ||
		(0x26FD == r) ||
		(0x26FE <= r && r <= 0x2701) ||
		(0x2702 == r) ||
		(0x2703 <= r && r <= 0x2704) ||
		(0x2705 == r) ||
		(0x2708 <= r && r <= 0x270C) ||
		(0x270D == r) ||
		(0x270E == r) ||
		(0x270F == r) ||
		(0x2710 <= r && r <= 0x2711) ||
		(0x2712 == r) ||
		(0x2714 == r) ||
		(0x2716 == r) ||
		(0x271D == r) ||
		(0x2721 == r) ||
		(0x2728 == r) ||
		(0x2733 <= r && r <= 0x2734) ||
		(0x2744 == r) ||
		(0x2747 == r) ||
		(0x274C == r) ||
		(0x274E == r) ||
		(0x2753 <= r && r <= 0x2755) ||
		(0x2757 == r) ||
		(0x2763 == r) ||
		(0x2764 == r) ||
		(0x2765 <= r && r <= 0x2767) ||
		(0x2795 <= r && r <= 0x2797) ||
		(0x27A1 == r) ||
		(0x27B0 == r) ||
		(0x27BF == r) ||
		(0x2934 <= r && r <= 0x2935) ||
		(0x2B05 <= r && r <= 0x2B07) ||
		(0x2B1B <= r && r <= 0x2B1C) ||
		(0x2B50 == r) ||
		(0x2B55 == r) ||
		(0x3030 == r) ||
		(0x303D == r) ||
		(0x3297 == r) ||
		(0x3299 == r) ||
		(0x1F000 <= r && r <= 0x1F003) ||
		(0x1F004 == r) ||
		(0x1F005 <= r && r <= 0x1F0CE) ||
		(0x1F0CF == r) ||
		(0x1F0D0 <= r && r <= 0x1F0FF) ||
		(0x1F10D <= r && r <= 0x1F10F) ||
		(0x1F12F == r) ||
		(0x1F16C <= r && r <= 0x1F16F) ||
		(0x1F170 <= r && r <= 0x1F171) ||
		(0x1F17E <= r && r <= 0x1F17F) ||
		(0x1F18E == r) ||
		(0x1F191 <= r && r <= 0x1F19A) ||
		(0x1F1AD <= r && r <= 0x1F1E5) ||
		(0x1F201 <= r && r <= 0x1F202) ||
		(0x1F203 <= r && r <= 0x1F20F) ||
		(0x1F21A == r) ||
		(0x1F22F == r) ||
		(0x1F232 <= r && r <= 0x1F23A) ||
		(0x1F23C <= r && r <= 0x1F23F) ||
		(0x1F249 <= r && r <= 0x1F24F) ||
		(0x1F250 <= r && r <= 0x1F251) ||
		(0x1F252 <= r && r <= 0x1F2FF) ||
		(0x1F300 <= r && r <= 0x1F30C) ||
		(0x1F30D <= r && r <= 0x1F30E) ||
		(0x1F30F == r) ||
		(0x1F310 == r) ||
		(0x1F311 == r) ||
		(0x1F312 == r) ||
		(0x1F313 <= r && r <= 0x1F315) ||
		(0x1F316 <= r && r <= 0x1F318) ||
		(0x1F319 == r) ||
		(0x1F31A == r) ||
		(0x1F31B == r) ||
		(0x1F31C == r) ||
		(0x1F31D <= r && r <= 0x1F31E) ||
		(0x1F31F <= r && r <= 0x1F320) ||
		(0x1F321 == r) ||
		(0x1F322 <= r && r <= 0x1F323) ||
		(0x1F324 <= r && r <= 0x1F32C) ||
		(0x1F32D <= r && r <= 0x1F32F) ||
		(0x1F330 <= r && r <= 0x1F331) ||
		(0x1F332 <= r && r <= 0x1F333) ||
		(0x1F334 <= r && r <= 0x1F335) ||
		(0x1F336 == r) ||
		(0x1F337 <= r && r <= 0x1F34A) ||
		(0x1F34B == r) ||
		(0x1F34C <= r && r <= 0x1F34F) ||
		(0x1F350 == r) ||
		(0x1F351 <= r && r <= 0x1F37B) ||
		(0x1F37C == r) ||
		(0x1F37D == r) ||
		(0x1F37E <= r && r <= 0x1F37F) ||
		(0x1F380 <= r && r <= 0x1F393) ||
		(0x1F394 <= r && r <= 0x1F395) ||
		(0x1F396 <= r && r <= 0x1F397) ||
		(0x1F398 == r) ||
		(0x1F399 <= r && r <= 0x1F39B) ||
		(0x1F39C <= r && r <= 0x1F39D) ||
		(0x1F39E <= r && r <= 0x1F39F) ||
		(0x1F3A0 <= r && r <= 0x1F3C4) ||
		(0x1F3C5 == r) ||
		(0x1F3C6 == r) ||
		(0x1F3C7 == r) ||
		(0x1F3C8 == r) ||
		(0x1F3C9 == r) ||
		(0x1F3CA == r) ||
		(0x1F3CB <= r && r <= 0x1F3CE) ||
		(0x1F3CF <= r && r <= 0x1F3D3) ||
		(0x1F3D4 <= r && r <= 0x1F3DF) ||
		(0x1F3E0 <= r && r <= 0x1F3E3) ||
		(0x1F3E4 == r) ||
		(0x1F3E5 <= r && r <= 0x1F3F0) ||
		(0x1F3F1 <= r && r <= 0x1F3F2) ||
		(0x1F3F3 == r) ||
		(0x1F3F4 == r) ||
		(0x1F3F5 == r) ||
		(0x1F3F6 == r) ||
		(0x1F3F7 == r) ||
		(0x1F3F8 <= r && r <= 0x1F3FA) ||
		(0x1F400 <= r && r <= 0x1F407) ||
		(0x1F408 == r) ||
		(0x1F409 <= r && r <= 0x1F40B) ||
		(0x1F40C <= r && r <= 0x1F40E) ||
		(0x1F40F <= r && r <= 0x1F410) ||
		(0x1F411 <= r && r <= 0x1F412) ||
		(0x1F413 == r) ||
		(0x1F414 == r) ||
		(0x1F415 == r) ||
		(0x1F416 == r) ||
		(0x1F417 <= r && r <= 0x1F429) ||
		(0x1F42A == r) ||
		(0x1F42B <= r && r <= 0x1F43E) ||
		(0x1F43F == r) ||
		(0x1F440 == r) ||
		(0x1F441 == r) ||
		(0x1F442 <= r && r <= 0x1F464) ||
		(0x1F465 == r) ||
		(0x1F466 <= r && r <= 0x1F46B) ||
		(0x1F46C <= r && r <= 0x1F46D) ||
		(0x1F46E <= r && r <= 0x1F4AC) ||
		(0x1F4AD == r) ||
		(0x1F4AE <= r && r <= 0x1F4B5) ||
		(0x1F4B6 <= r && r <= 0x1F4B7) ||
		(0x1F4B8 <= r && r <= 0x1F4EB) ||
		(0x1F4EC <= r && r <= 0x1F4ED) ||
		(0x1F4EE == r) ||
		(0x1F4EF == r) ||
		(0x1F4F0 <= r && r <= 0x1F4F4) ||
		(0x1F4F5 == r) ||
		(0x1F4F6 <= r && r <= 0x1F4F7) ||
		(0x1F4F8 == r) ||
		(0x1F4F9 <= r && r <= 0x1F4FC) ||
		(0x1F4FD == r) ||
		(0x1F4FE == r) ||
		(0x1F4FF <= r && r <= 0x1F502) ||
		(0x1F503 == r) ||
		(0x1F504 <= r && r <= 0x1F507) ||
		(0x1F508 == r) ||
		(0x1F509 == r) ||
		(0x1F50A <= r && r <= 0x1F514) ||
		(0x1F515 == r) ||
		(0x1F516 <= r && r <= 0x1F52B) ||
		(0x1F52C <= r && r <= 0x1F52D) ||
		(0x1F52E <= r && r <= 0x1F53D) ||
		(0x1F546 <= r && r <= 0x1F548) ||
		(0x1F549 <= r && r <= 0x1F54A) ||
		(0x1F54B <= r && r <= 0x1F54E) ||
		(0x1F54F == r) ||
		(0x1F550 <= r && r <= 0x1F55B) ||
		(0x1F55C <= r && r <= 0x1F567) ||
		(0x1F568 <= r && r <= 0x1F56E) ||
		(0x1F56F <= r && r <= 0x1F570) ||
		(0x1F571 <= r && r <= 0x1F572) ||
		(0x1F573 <= r && r <= 0x1F579) ||
		(0x1F57A == r) ||
		(0x1F57B <= r && r <= 0x1F586) ||
		(0x1F587 == r) ||
		(0x1F588 <= r && r <= 0x1F589) ||
		(0x1F58A <= r && r <= 0x1F58D) ||
		(0x1F58E <= r && r <= 0x1F58F) ||
		(0x1F590 == r) ||
		(0x1F591 <= r && r <= 0x1F594) ||
		(0x1F595 <= r && r <= 0x1F596) ||
		(0x1F597 <= r && r <= 0x1F5A3) ||
		(0x1F5A4 == r) ||
		(0x1F5A5 == r) ||
		(0x1F5A6 <= r && r <= 0x1F5A7) ||
		(0x1F5A8 == r) ||
		(0x1F5A9 <= r && r <= 0x1F5B0) ||
		(0x1F5B1 <= r && r <= 0x1F5B2) ||
		(0x1F5B3 <= r && r <= 0x1F5BB) ||
		(0x1F5BC == r) ||
		(0x1F5BD <= r && r <= 0x1F5C1) ||
		(0x1F5C2 <= r && r <= 0x1F5C4) ||
		(0x1F5C5 <= r && r <= 0x1F5D0) ||
		(0x1F5D1 <= r && r <= 0x1F5D3) ||
		(0x1F5D4 <= r && r <= 0x1F5DB) ||
		(0x1F5DC <= r && r <= 0x1F5DE) ||
		(0x1F5DF <= r && r <= 0x1F5E0) ||
		(0x1F5E1 == r) ||
		(0x1F5E2 == r) ||
		(0x1F5E3 == r) ||
		(0x1F5E4 <= r && r <= 0x1F5E7) ||
		(0x1F5E8 == r) ||
		(0x1F5E9 <= r && r <= 0x1F5EE) ||
		(0x1F5EF == r) ||
		(0x1F5F0 <= r && r <= 0x1F5F2) ||
		(0x1F5F3 == r) ||
		(0x1F5F4 <= r && r <= 0x1F5F9) ||
		(0x1F5FA == r) ||
		(0x1F5FB <= r && r <= 0x1F5FF) ||
		(0x1F600 == r) ||
		(0x1F601 <= r && r <= 0x1F606) ||
		(0x1F607 <= r && r <= 0x1F608) ||
		(0x1F609 <= r && r <= 0x1F60D) ||
		(0x1F60E == r) ||
		(0x1F60F == r) ||
		(0x1F610 == r) ||
		(0x1F611 == r) ||
		(0x1F612 <= r && r <= 0x1F614) ||
		(0x1F615 == r) ||
		(0x1F616 == r) ||
		(0x1F617 == r) ||
		(0x1F618 == r) ||
		(0x1F619 == r) ||
		(0x1F61A == r) ||
		(0x1F61B == r) ||
		(0x1F61C <= r && r <= 0x1F61E) ||
		(0x1F61F == r) ||
		(0x1F620 <= r && r <= 0x1F625) ||
		(0x1F626 <= r && r <= 0x1F627) ||
		(0x1F628 <= r && r <= 0x1F62B) ||
		(0x1F62C == r) ||
		(0x1F62D == r) ||
		(0x1F62E <= r && r <= 0x1F62F) ||
		(0x1F630 <= r && r <= 0x1F633) ||
		(0x1F634 == r) ||
		(0x1F635 == r) ||
		(0x1F636 == r) ||
		(0x1F637 <= r && r <= 0x1F640) ||
		(0x1F641 <= r && r <= 0x1F644) ||
		(0x1F645 <= r && r <= 0x1F64F) ||
		(0x1F680 == r) ||
		(0x1F681 <= r && r <= 0x1F682) ||
		(0x1F683 <= r && r <= 0x1F685) ||
		(0x1F686 == r) ||
		(0x1F687 == r) ||
		(0x1F688 == r) ||
		(0x1F689 == r) ||
		(0x1F68A <= r && r <= 0x1F68B) ||
		(0x1F68C == r) ||
		(0x1F68D == r) ||
		(0x1F68E == r) ||
		(0x1F68F == r) ||
		(0x1F690 == r) ||
		(0x1F691 <= r && r <= 0x1F693) ||
		(0x1F694 == r) ||
		(0x1F695 == r) ||
		(0x1F696 == r) ||
		(0x1F697 == r) ||
		(0x1F698 == r) ||
		(0x1F699 <= r && r <= 0x1F69A) ||
		(0x1F69B <= r && r <= 0x1F6A1) ||
		(0x1F6A2 == r) ||
		(0x1F6A3 == r) ||
		(0x1F6A4 <= r && r <= 0x1F6A5) ||
		(0x1F6A6 == r) ||
		(0x1F6A7 <= r && r <= 0x1F6AD) ||
		(0x1F6AE <= r && r <= 0x1F6B1) ||
		(0x1F6B2 == r) ||
		(0x1F6B3 <= r && r <= 0x1F6B5) ||
		(0x1F6B6 == r) ||
		(0x1F6B7 <= r && r <= 0x1F6B8) ||
		(0x1F6B9 <= r && r <= 0x1F6BE) ||
		(0x1F6BF == r) ||
		(0x1F6C0 == r) ||
		(0x1F6C1 <= r && r <= 0x1F6C5) ||
		(0x1F6C6 <= r && r <= 0x1F6CA) ||
		(0x1F6CB == r) ||
		(0x1F6CC == r) ||
		(0x1F6CD <= r && r <= 0x1F6CF) ||
		(0x1F6D0 == r) ||
		(0x1F6D1 <= r && r <= 0x1F6D2) ||
		(0x1F6D3 <= r && r <= 0x1F6D4) ||
		(0x1F6D5 == r) ||
		(0x1F6D6 <= r && r <= 0x1F6D7) ||
		(0x1F6D8 <= r && r <= 0x1F6DF) ||
		(0x1F6E0 <= r && r <= 0x1F6E5) ||
		(0x1F6E6 <= r && r <= 0x1F6E8) ||
		(0x1F6E9 == r) ||
		(0x1F6EA == r) ||
		(0x1F6EB <= r && r <= 0x1F6EC) ||
		(0x1F6ED <= r && r <= 0x1F6EF) ||
		(0x1F6F0 == r) ||
		(0x1F6F1 <= r && r <= 0x1F6F2) ||
		(0x1F6F3 == r) ||
		(0x1F6F4 <= r && r <= 0x1F6F6) ||
		(0x1F6F7 <= r && r <= 0x1F6F8) ||
		(0x1F6F9 == r) ||
		(0x1F6FA == r) ||
		(0x1F6FB <= r && r <= 0x1F6FC) ||
		(0x1F6FD <= r && r <= 0x1F6FF) ||
		(0x1F774 <= r && r <= 0x1F77F) ||
		(0x1F7D5 <= r && r <= 0x1F7DF) ||
		(0x1F7E0 <= r && r <= 0x1F7EB) ||
		(0x1F7EC <= r && r <= 0x1F7FF) ||
		(0x1F80C <= r && r <= 0x1F80F) ||
		(0x1F848 <= r && r <= 0x1F84F) ||
		(0x1F85A <= r && r <= 0x1F85F) ||
		(0x1F888 <= r && r <= 0x1F88F) ||
		(0x1F8AE <= r && r <= 0x1F8FF) ||
		(0x1F90C == r) ||
		(0x1F90D <= r && r <= 0x1F90F) ||
		(0x1F910 <= r && r <= 0x1F918) ||
		(0x1F919 <= r && r <= 0x1F91E) ||
		(0x1F91F == r) ||
		(0x1F920 <= r && r <= 0x1F927) ||
		(0x1F928 <= r && r <= 0x1F92F) ||
		(0x1F930 == r) ||
		(0x1F931 <= r && r <= 0x1F932) ||
		(0x1F933 <= r && r <= 0x1F93A) ||
		(0x1F93C <= r && r <= 0x1F93E) ||
		(0x1F93F == r) ||
		(0x1F940 <= r && r <= 0x1F945) ||
		(0x1F947 <= r && r <= 0x1F94B) ||
		(0x1F94C == r) ||
		(0x1F94D <= r && r <= 0x1F94F) ||
		(0x1F950 <= r && r <= 0x1F95E) ||
		(0x1F95F <= r && r <= 0x1F96B) ||
		(0x1F96C <= r && r <= 0x1F970) ||
		(0x1F971 == r) ||
		(0x1F972 == r) ||
		(0x1F973 <= r && r <= 0x1F976) ||
		(0x1F977 <= r && r <= 0x1F978) ||
		(0x1F979 == r) ||
		(0x1F97A == r) ||
		(0x1F97B == r) ||
		(0x1F97C <= r && r <= 0x1F97F) ||
		(0x1F980 <= r && r <= 0x1F984) ||
		(0x1F985 <= r && r <= 0x1F991) ||
		(0x1F992 <= r && r <= 0x1F997) ||
		(0x1F998 <= r && r <= 0x1F9A2) ||
		(0x1F9A3 <= r && r <= 0x1F9A4) ||
		(0x1F9A5 <= r && r <= 0x1F9AA) ||
		(0x1F9AB <= r && r <= 0x1F9AD) ||
		(0x1F9AE <= r && r <= 0x1F9AF) ||
		(0x1F9B0 <= r && r <= 0x1F9B9) ||
		(0x1F9BA <= r && r <= 0x1F9BF) ||
		(0x1F9C0 == r) ||
		(0x1F9C1 <= r && r <= 0x1F9C2) ||
		(0x1F9C3 <= r && r <= 0x1F9CA) ||
		(0x1F9CB == r) ||
		(0x1F9CC == r) ||
		(0x1F9CD <= r && r <= 0x1F9CF) ||
		(0x1F9D0 <= r && r <= 0x1F9E6) ||
		(0x1F9E7 <= r && r <= 0x1F9FF) ||
		(0x1FA00 <= r && r <= 0x1FA6F) ||
		(0x1FA70 <= r && r <= 0x1FA73) ||
		(0x1FA74 == r) ||
		(0x1FA75 <= r && r <= 0x1FA77) ||
		(0x1FA78 <= r && r <= 0x1FA7A) ||
		(0x1FA7B <= r && r <= 0x1FA7F) ||
		(0x1FA80 <= r && r <= 0x1FA82) ||
		(0x1FA83 <= r && r <= 0x1FA86) ||
		(0x1FA87 <= r && r <= 0x1FA8F) ||
		(0x1FA90 <= r && r <= 0x1FA95) ||
		(0x1FA96 <= r && r <= 0x1FAA8) ||
		(0x1FAA9 <= r && r <= 0x1FAAF) ||
		(0x1FAB0 <= r && r <= 0x1FAB6) ||
		(0x1FAB7 <= r && r <= 0x1FABF) ||
		(0x1FAC0 <= r && r <= 0x1FAC2) ||
		(0x1FAC3 <= r && r <= 0x1FACF) ||
		(0x1FAD0 <= r && r <= 0x1FAD6) ||
		(0x1FAD7 <= r && r <= 0x1FAFF) ||
		(0x1FC00 <= r && r <= 0x1FFFD)
}
