#!/bin/bash
# Build HTML and PDF from AsciiDoc sources.
# Usage: ./go [file.adoc ...]   (defaults to example.adoc)
set -euo pipefail

export PATH="$(ruby -e 'print Gem.user_dir')/bin:$PATH"
# asciidoctor-pdf-mathjax renders STEM with MathJax 4 from global npm packages
export NODE_PATH="$(npm root -g)"

files=("${@:-example.adoc}")

for f in "${files[@]}"; do
  asciidoctor -a stylesheet=my.css "$f"
  asciidoctor-pdf -r asciidoctor-pdf-mathjax "$f"
done

[[ -n "${NO_OPEN:-}" ]] || xdg-open "${files[0]%.adoc}.html" >/dev/null 2>&1 &
