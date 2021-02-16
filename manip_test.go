package rosed

import "testing"

func Test_Manip_Wrap(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		width    int
		sep      string
		expected []string
	}{
		{"empty input", "", 80, "\n", []string{""}},
		{"not enough to wrap", "a test string", 80, "\n", []string{"a test string"}},
		{"2 line wrap", "a string long enough to be wrapped", 20, "\n", []string{"a string long enough", "to be wrapped"}},
		// PLACE IS HERE
		{"multi line wrap", args{20, nil, "a string long enough to be wrapped more than once"}, []string{"a string long enough", "to be wrapped more", "than once"}},
		{"invalid width of -1 is interpreted as 2", args{-1, nil, "test"}, []string{"t-", "e-", "st"}},
		{"invalid width of 0 is interpreted as 2", args{0, nil, "test"}, []string{"t-", "e-", "st"}},
		{"invalid width of 1 is interpreted as 2", args{1, nil, "test"}, []string{"t-", "e-", "st"}},
		{"valid width of 2", args{2, nil, "test"}, []string{"t-", "e-", "st"}},
		{"valid width of 3", args{3, nil, "test"}, []string{"te-", "st"}},
		{"2 paragraphs, not preserved", args{20, nil, "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph"}, []string{"this is a line that", "is split by", "paragraph in the", "input. This is the", "second paragraph"}},
		{"1 paragraph, preserved", args{20, &WrapOptions{PreserveParagraphs: true}, "this is a line that is split by paragraph in the input."}, []string{"this is a line that", "is split by", "paragraph in the", "input."}},
		{"2 paragraphs, preserved", args{20, &WrapOptions{PreserveParagraphs: true}, "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph"}, []string{"this is a line that", "is split by", "paragraph in the", "input.", "", "This is the second", "paragraph"}},
		{"3 paragraphs, preserved", args{20, &WrapOptions{PreserveParagraphs: true}, "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph.\n\nAnd this is the third"}, []string{"this is a line that", "is split by", "paragraph in the", "input.", "", "This is the second", "paragraph.", "", "And this is the", "third"}},
		{"3 paragraphs, preserved with suffix & prefix", args{20, &WrapOptions{PreserveParagraphs: true, Suffix: "]", Prefix: "["}, "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph.\n\nAnd this is the third"}, []string{"[this is a line]", "[that is split by]", "[paragraph in the]", "[input.]", "", "[This is the second]", "[paragraph.]", "", "[And this is the]", "[third]"}},
		{"3 paragraphs, preserved with para prefix & suffix & prefix", args{20, &WrapOptions{PreserveParagraphs: true, ParagraphPrefix: "- ", Suffix: "]", Prefix: "["}, "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph.\n\nAnd this is the third"}, []string{"- [this is a line]", "[that is split by]", "[paragraph in the]", "[input.]", "", "- [This is the]", "[second paragraph.]", "", "- [And this is the]", "[third]"}},
		{"2 paragraphs, unpreserved with para prefix", args{20, &WrapOptions{ParagraphPrefix: "- "}, "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph"}, []string{"- this is a line", "that is split by", "paragraph in the", "input. This is the", "second paragraph"}},
	}
}
