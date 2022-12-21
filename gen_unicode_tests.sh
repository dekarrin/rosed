#!/bin/bash

repo_path="$(dirname "$0")"

cat << EOF > "$repo_path/internal/gem/unicode_test.go"
package gem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

$(python "$repo_path/get_unicode_grapheme_tests.py")
EOF
