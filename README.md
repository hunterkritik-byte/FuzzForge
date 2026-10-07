# FuzzForge ⚡

**FuzzForge** is our own modern JavaScript engine fuzzer, built in Go and designed around a coverage-guided fuzzing pipeline.

It takes inspiration from research fuzzers such as Fuzzilli, but the implementation, architecture, corpus, mutation system, and feedback loop are being built independently.

## Status

### Phase 0 — Bootstrap
- [x] Go CLI
- [x] JavaScript seed corpus
- [x] Mutation primitives
- [x] d8 execution harness
- [x] Crash artifact preservation
- [ ] Timeout enforcement
- [ ] Coverage feedback
- [ ] Corpus scheduling
- [ ] JavaScript-aware generation
- [ ] Minimization

## Quick start

Build:

```bash
go build -o fuzzforge ./cmd/fuzzforge
```

Run against a local V8 `d8`:

```bash
./fuzzforge -engine ./d8 -corpus ./corpus -runs 100
```

Generated programs and crash candidates are written to `artifacts/`.

## Architecture

```
             seed corpus
                  |
                  v
       mutation / generation
                  |
                  v
          JavaScript program
                  |
                  v
             d8 / engine
             /    |    \
            /     |     \
         clean  timeout  crash
            \     |     /
             \    |    /
              artifact
                  |
          coverage feedback
                  |
             corpus queue
                  |
                  +------> next generation
```

## Roadmap

1. Coverage-guided corpus scheduling
2. Persistent corpus + deduplication
3. JavaScript AST/program generation
4. V8/JIT-focused mutation operators
5. Automatic testcase minimization
6. Crash bucketing + reproducibility
7. Differential engine fuzzing
8. CI regression mode

## Safety

Run FuzzForge only against JavaScript engines and software you are authorized to test. Crash artifacts can consume disk space quickly, so use bounded campaigns while developing.

## License

MIT
