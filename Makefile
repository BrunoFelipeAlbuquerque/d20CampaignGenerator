.PHONY: pf1-extract pf1-chunk pf1-search test

pf1-extract:
	./docs/pf1/extract_rules.sh

pf1-chunk:
	go run docs/pf1/chunk_rules.go

pf1-search:
	@test -n "$(q)" || (echo 'usage: make pf1-search q="ability damage"' && exit 1)
	rg -i "$(q)" docs/pf1/chunks

test:
	go test ./...
