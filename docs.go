package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const outputpath = "docs/_recipes"

func Walker(pathmap map[string]string) filepath.WalkFunc {
	return func(pathstr string, info fs.FileInfo, err error) error {
		if info.IsDir() || filepath.Ext(pathstr) != ".cook" {
			return nil
		}
		dir, cookfilename := filepath.Split(strings.Replace(pathstr, "recipes", outputpath, 1))
		if dir != "" {
			if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
				return fmt.Errorf("failed to create dir %s: %w", dir, err)
			}
		}
		cookcli := exec.Command("cook", "recipe", "-f", "markdown", pathstr)
		var stdout, stderr strings.Builder
		cookcli.Stdout = &stdout
		cookcli.Stderr = &stderr
		if err := cookcli.Run(); err != nil {
			return fmt.Errorf("failed to create markdown using cookcli for file %s: %w: %s", cookfilename, err, stderr.String())
		}
		outbytes, err := io.ReadAll(strings.NewReader(stdout.String()))
		if err != nil {
			return fmt.Errorf("failed to read stdout of cookcli process for file %s: %w", cookfilename, err)
		}
		lines := make([]string, 0)
		for _, line := range strings.Split(string(outbytes), "\n") {
			if line != "# " {
				lines = append(lines, line)
			}
		}
		if err := os.WriteFile(strings.Replace(dir+cookfilename, ".cook", ".md", 1), []byte(strings.Join(lines, "\n")), fs.ModePerm); err != nil {
			return fmt.Errorf("failed to write stdout of cookcli process for file %s: %w", cookfilename, err)
		}
		return nil
	}
}

func main() {
	pathmap := make(map[string]string)
	if err := filepath.Walk("recipes", Walker(pathmap)); err != nil {
		log.Fatalf("failed to traverse cook files: %v", err)
	}
	log.Printf("%+v", pathmap)
}
