#!/bin/bash
# Build HTML and PDF from AsciiDoc sources.
# Usage: ./go [file.adoc ...]   (defaults to example.adoc)
set -euo pipefail

export PATH="$(ruby -e 'print Gem.user_dir')/bin:$PATH"

files=("${@:-example.adoc}")

for f in "${files[@]}"; do
  asciidoctor -a stylesheet=my.css "$f"
  asciidoctor-pdf "$f"
done

xdg-open "${files[0]%.adoc}.html" >/dev/null 2>&1 &
