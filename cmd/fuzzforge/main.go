package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hunterkritik-byte/FuzzForge/internal/runner"
)

func main() {
	engine := flag.String("engine", "d8", "JavaScript engine executable")
	corpus := flag.String("corpus", "corpus", "seed corpus directory")
	out := flag.String("out", "artifacts", "artifact directory")
	runs := flag.Int("runs", 100, "number of executions")
	flag.Parse()

	if err := runner.Run(runner.Config{Engine:*engine, CorpusDir:*corpus, ArtifactDir:*out, Runs:*runs}); err != nil {
		fmt.Fprintln(os.Stderr, "fuzzforge:", err)
		os.Exit(1)
	}
}
