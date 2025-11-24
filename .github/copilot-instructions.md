# Copilot Instructions for m346-go-exercises

## Project Overview
This repository contains Go programming exercises for the Swiss vocational education module 346. The exercises are organized by topic and increasing complexity, with each topic in its own set of folders and markdown guides.

## Structure & Conventions
- Each exercise is in its own folder (e.g., `go-1-ex-1/`, `go-2-ex-4/`), containing a `main.go` file.
- Markdown files (`go-1-exercises.md`, etc.) describe the tasks and learning objectives for each topic.
- The `formulas/` directory contains LaTeX files for mathematical formulas referenced in exercises.
- The `tests/` directory contains text files for expected outputs and errors, but no automated test runner is present.

## Developer Workflows
- **Run an exercise:**
  - Navigate to the exercise folder (e.g., `cd go-2-ex-4`)
  - Run with `go run main.go`
- **No central build or test script:** Each exercise is standalone. There is no global build or test command.
- **No external dependencies:** The `go.mod` exists for Go tooling, but exercises use only the Go standard library.

## Project-Specific Patterns
- Each `main.go` is self-contained and does not import code from other exercises.
- Exercises are designed for learning and experimentation; code style may vary between folders.
- Mathematical logic may reference LaTeX in `formulas/`, but is implemented directly in Go.

## Key Files & Directories
- `README.md`: High-level overview and links to exercise guides.
- `go-*-exercises.md`: Exercise instructions per topic.
- `go-*-ex-*/main.go`: Exercise implementations.
- `formulas/`: LaTeX source for formulas (not used programmatically).
- `tests/`: Example outputs for manual comparison.

## Guidance for AI Agents
- When adding new exercises, follow the existing folder and file naming conventions.
- Do not introduce external dependencies or frameworks.
- Keep each exercise independent; avoid cross-folder imports.
- Reference the relevant markdown and LaTeX files for context when implementing or updating exercises.
- If automating tests, use the `tests/` outputs as a reference for expected behavior.
