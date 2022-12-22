#!/bin/bash

repo_path="$(dirname "$0")"

cat << EOF | gofmt > "$repo_path/internal/gem/unicode_test.go"
package gem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

$(python "$repo_path/scripts/get_unicode_grapheme_tests.py")
EOF
