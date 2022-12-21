# retrieves the tests provided by unicode for testing various implemenations of
# rules and converts them into something that can actually be used as go test
# code

import urllib.request
import sys

_ucd_grapheme_test_url = "https://www.unicode.org/Public/UCD/latest/ucd/auxiliary/GraphemeBreakTest.txt"

def py_listlist_to_golang_int_sliceslice(int_listlist):
    output = "[][]int{"
    for int_list in int_listlist:
        output += "{"
        for i in int_list:
            output += str(i) + ", "
            
        if len(int_list) > 0:
            output = output[:-2]  # chop off trailing comma+space
        output += "}, "
    if len(int_listlist) > 0:
        output = output[:-2]  # chop off trailing comma+space
    output += "}"
    return output


def main():
    print("Downloading GraphemeBreakTest.txt from current unicode data sources...", file=sys.stderr)
    
    grapheme_test_data = None
    with urllib.request.urlopen(_ucd_grapheme_test_url) as fp:
        grapheme_test_data = fp.read().decode('utf-8')
        
    print("Analyzing GraphemeBreakTest.txt...", file=sys.stderr)
    
    grapheme_lines = grapheme_test_data.splitlines()
    first_line = True
    test_line = 0
    test_cases = list()
    for line in grapheme_lines:
        if first_line:
            first_line = False
            name_base = line[line.find('#')+1:].strip().split("-", 1)[0]
            continue
        
        comment_start = line.find('#')
        line = line[:comment_start]
        
        line = line.strip()
            
        if line == "":
            continue
        if line.startswith('÷'):
            test_line += 1
            line = line[1:]
        
        # dont cut off the final ÷ tho bc it makes parsing easier
        line = line.strip()
        
        parts = line.split(' ')
        grapheme_clusters = list()
        cur_cluster = ""
        cur_rune_idx = -1
        cur_index = [0,]
        gc_indexes = list()
        test_case = {'name': name_base + " #" + str(test_line).zfill(3)}
        
        input_line = ""
        
        for p in parts:
            if p == '÷':
                grapheme_clusters.append(cur_cluster)
                input_line += cur_cluster
                cur_cluster = ""
                cur_index.append(cur_rune_idx + 1)
                gc_indexes.append(cur_index)
                cur_index = [cur_rune_idx + 1,]
            elif p == '×':
                pass #nothing to do, its joined with the next one
            else:
                cur_rune_idx += 1
                if len(p) <= 4:
                    cur_cluster += r'\u' + p.zfill(4)
                else:
                    cur_cluster += r'\U' + p.zfill(8)
        
        test_case['input'] = input_line
        test_case['expect'] = py_listlist_to_golang_int_sliceslice(gc_indexes)
        
        test_cases.append(test_case)
        
        
    test_func_start = "func Test_GraphemeClusterBreak(t *testing.T) {\n"
    test_func_start += "\ttestCases := []struct {\n"
    test_func_start += "\t\tname   string\n"
    test_func_start += "\t\tinput  String\n"
    test_func_start += "\t\texpect [][]int\n"
    test_func_start += "\t}{"

    test_func_end = "\t}\n"
    test_func_end += "\n"
    test_func_end += "\tfor _, tc := range testCases {\n"
    test_func_end += "\t\tt.Run(tc.name, func(t *testing.T) {\n"
    test_func_end += "\t\t\tassert := assert.New(t)\n"
    test_func_end += "\n"
    test_func_end += "\t\t\tactual := tc.input.GraphemeIndexes()\n"
    test_func_end += "\t\t})\n"
    test_func_end += "\t}\n"
    test_func_end +="}"
    
    print(test_func_start)
    for test_case in test_cases:
        fmt_str = '\t\t{{"{:s}", New("{:s}"), {:s}}},'
        print(fmt_str.format(test_case['name'], test_case['input'], test_case['expect']))
    print(test_func_end)
        
    
        
    # prog: jello, for each line, first use the breaks to build a preferred
    # GraphemeIndexes from a string composed of the codepoints.
    # then, remove the break/no-break chars and create a sequence made ONLY of
    # those codepoints, as \U or \u escape sequences in a go-string.
    # finally, assemble each into a test case. and at end and start, output a
    # test name, scaffolding, and ending

if __name__ == "__main__":
    main()