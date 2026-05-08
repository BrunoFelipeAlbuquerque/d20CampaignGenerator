package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const (
	chunkSize    = 3200
	chunkOverlap = 500
)

type ruleChunk struct {
	Source   string `json:"source"`
	TextFile string `json:"text_file"`
	Page     int    `json:"page"`
	Chunk    int    `json:"chunk"`
	Text     string `json:"text"`
}

func main() {
	dryRun := flag.Bool("dry-run", false, "validate chunk generation without writing files")
	flag.Parse()

	root, err := pf1Root()
	if err != nil {
		fatal(err)
	}

	chunks, err := buildChunks(root)
	if err != nil {
		fatal(err)
	}

	if *dryRun {
		fmt.Printf("generated %d chunk files\n", len(chunks))
		for file, fileChunks := range chunks {
			fmt.Printf("%s: %d chunks\n", file, len(fileChunks))
		}
		return
	}

	if err := writeChunks(root, chunks); err != nil {
		fatal(err)
	}
}

func pf1Root() (string, error) {
	candidates := []string{
		"docs/pf1",
		".",
	}

	for _, candidate := range candidates {
		textDir := filepath.Join(candidate, "text")
		if stat, err := os.Stat(textDir); err == nil && stat.IsDir() {
			return candidate, nil
		}
	}

	return "", errors.New("could not find docs/pf1 text directory")
}

func buildChunks(root string) (map[string][]ruleChunk, error) {
	textFiles, err := filepath.Glob(filepath.Join(root, "text", "*.txt"))
	if err != nil {
		return nil, err
	}
	if len(textFiles) == 0 {
		return nil, errors.New("no PF1 text files found")
	}
	sort.Strings(textFiles)

	chunks := make(map[string][]ruleChunk, len(textFiles))
	for _, textPath := range textFiles {
		fileChunks, err := chunkTextFile(root, textPath)
		if err != nil {
			return nil, err
		}

		outName := strings.TrimSuffix(filepath.Base(textPath), filepath.Ext(textPath)) + ".jsonl"
		chunks[outName] = fileChunks
	}

	return chunks, nil
}

func chunkTextFile(root string, textPath string) ([]ruleChunk, error) {
	data, err := os.ReadFile(textPath)
	if err != nil {
		return nil, err
	}

	textFile := filepath.Base(textPath)
	source := sourceName(root, textFile)
	pages := splitPages(string(data), firstPageNumber(string(data)))

	var chunks []ruleChunk
	for _, page := range pages {
		pageChunks := chunkPageText(page.text)
		for i, text := range pageChunks {
			chunks = append(chunks, ruleChunk{
				Source:   source,
				TextFile: textFile,
				Page:     page.number,
				Chunk:    i + 1,
				Text:     text,
			})
		}
	}

	if len(chunks) == 0 {
		return nil, fmt.Errorf("no chunks generated for %s", textFile)
	}

	return chunks, nil
}

func sourceName(root string, textFile string) string {
	base := strings.TrimSuffix(textFile, filepath.Ext(textFile))
	pdfName := base + ".pdf"

	for _, candidate := range []string{
		filepath.Join(root, pdfName),
		filepath.Join(root, "raw", pdfName),
	} {
		if _, err := os.Stat(candidate); err == nil {
			return pdfName
		}
	}

	return textFile
}

type pageText struct {
	number int
	text   string
}

func splitPages(text string, firstPage int) []pageText {
	parts := strings.Split(text, "\f")
	pages := make([]pageText, 0, len(parts))

	for i, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}

		pages = append(pages, pageText{
			number: firstPage + i,
			text:   trimmed,
		})
	}

	return pages
}

var pageRangePattern = regexp.MustCompile(`(?i)\bpages?\s+([0-9]+)`)

func firstPageNumber(text string) int {
	if strings.Contains(text, "\f") {
		return 1
	}

	matches := pageRangePattern.FindStringSubmatch(text)
	if len(matches) != 2 {
		return 1
	}

	var page int
	if _, err := fmt.Sscanf(matches[1], "%d", &page); err != nil || page < 1 {
		return 1
	}

	return page
}

func chunkPageText(text string) []string {
	if len(text) <= chunkSize {
		return []string{text}
	}

	var chunks []string
	for start := 0; start < len(text); {
		end := start + chunkSize
		if end > len(text) {
			end = len(text)
		}

		chunks = append(chunks, strings.TrimSpace(text[start:end]))
		if end == len(text) {
			break
		}

		start = end - chunkOverlap
	}

	return chunks
}

func writeChunks(root string, chunks map[string][]ruleChunk) error {
	chunkDir := filepath.Join(root, "chunks")
	if err := os.MkdirAll(chunkDir, 0o755); err != nil {
		return err
	}

	names := make([]string, 0, len(chunks))
	for name := range chunks {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		if err := writeChunkFile(filepath.Join(chunkDir, name), chunks[name]); err != nil {
			return err
		}
	}

	return nil
}

func writeChunkFile(path string, chunks []ruleChunk) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	encoder := json.NewEncoder(writer)
	for _, chunk := range chunks {
		if err := encoder.Encode(chunk); err != nil {
			return err
		}
	}

	return writer.Flush()
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
