# retrieves the tests provided by unicode for testing various implemenations of
# rules and converts them into something that can actually be used as go test
# code

import urllib.request

_ucd_grapheme_test_url = "https://www.unicode.org/Public/UCD/latest/ucd/auxiliary/GraphemeBreakTest.txt"

def main():
    print("Downloading GraphemeBreakTest.txt from current unicode data sources...")
    
    grapheme_test_data = None
    with urllib.request.urlopen(_ucd_grapheme_test_url) as fp:
        grapheme_test_data = fp.read().decode('utf-8')
        
    print("Analyzing GraphemeBreakTest.txt...")
    
    grapheme_lines = grapheme_test_data.splitlines()
    for line in grapheme_lines:
        comment_start = line.find('#')
        line = line[:comment_start]
        line = line.strip()
        if line == "":
            continue
        print(line)

if __name__ == "__main__":
    main()