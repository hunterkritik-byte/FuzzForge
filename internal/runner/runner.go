package runner

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/hunterkritik-byte/FuzzForge/internal/mutator"
)

type Config struct {
	Engine      string
	CorpusDir   string
	ArtifactDir string
	Runs        int
}

func Run(cfg Config) error {
	seeds, err := os.ReadDir(cfg.CorpusDir)
	if err != nil {
		return fmt.Errorf("read corpus: %w", err)
	}
	if len(seeds) == 0 {
		return fmt.Errorf("corpus is empty")
	}
	if err := os.MkdirAll(cfg.ArtifactDir, 0755); err != nil {
		return err
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < cfg.Runs; i++ {
		entry := seeds[r.Intn(len(seeds))]
		data, err := os.ReadFile(filepath.Join(cfg.CorpusDir, entry.Name()))
		if err != nil {
			return err
		}
		program := mutator.Mutate(string(data), r)
		path := filepath.Join(cfg.ArtifactDir, fmt.Sprintf("case-%06d.js", i))
		if err := os.WriteFile(path, []byte(program), 0644); err != nil {
			return err
		}

		cmd := exec.Command(cfg.Engine, path)
		output, err := cmd.CombinedOutput()
		if err != nil {
			kind := "crash"
			if cmd.ProcessState != nil && cmd.ProcessState.ExitCode() == 124 {
				kind = "timeout"
			}
			dst := filepath.Join(cfg.ArtifactDir, fmt.Sprintf("%s-%06d.js", kind, i))
			_ = os.WriteFile(dst, []byte(program), 0644)
			fmt.Printf("[%s] %s\n", kind, dst)
			if len(output) > 0 {
				fmt.Printf("%s\n", output)
			}
		}
	}
	return nil
}
