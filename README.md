# FuzzForge ⚡

Modern JavaScript engine fuzzer written in Go.

## Implemented

- [x] Timeout enforcement — engine processes are killed after the deadline
- [x] Crash classification — ordinary d8 JavaScript exceptions are no longer reported as engine crashes
- [x] Coverage feedback pipeline — successful novel candidates have a corpus-admission point ready for engine coverage
- [x] Corpus scheduling — generated candidates are periodically fed back into the corpus
- [x] JavaScript-aware generation — structured functions, arrays, objects, loops, exceptions, proxies and buffers
- [x] Minimization — delta-reduces a reproducing failing testcase

## Build

```bash
go build -o fuzzforge ./cmd/fuzzforge
```

## Test with d8

```bash
./fuzzforge -engine ./d8 -corpus ./corpus -runs 100 -timeout 2s
```

You should see `[corpus+]`, `[crash]`, or `[timeout]` events.

## Minimize

```bash
./fuzzforge -engine ./d8 -minimize artifacts/crash-000001.js
```

## Next

The next coverage step is an engine-specific novelty bitmap and weighted corpus scheduler. The current implementation deliberately keeps the execution boundary independent so d8 testing works first.
