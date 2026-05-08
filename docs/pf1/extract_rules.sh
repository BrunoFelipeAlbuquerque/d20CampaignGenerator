#!/bin/sh
set -eu

script_dir=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
text_dir="$script_dir/text"

if ! command -v pdftotext >/dev/null 2>&1; then
	echo "pdftotext is required to extract PF1 rules text" >&2
	exit 1
fi

mkdir -p "$text_dir"

found=0
for pdf in "$script_dir"/*.pdf "$script_dir/raw"/*.pdf; do
	if [ ! -f "$pdf" ]; then
		continue
	fi

	found=1
	name=$(basename -- "$pdf")
	base=${name%.*}
	pdftotext -layout "$pdf" "$text_dir/$base.txt"
done

if [ "$found" -eq 0 ]; then
	echo "no PF1 PDF files found under $script_dir or $script_dir/raw" >&2
	exit 1
fi
