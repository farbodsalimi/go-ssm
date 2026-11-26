# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

`go-ssm` is a single static binary that fetches AWS SSM parameters, injects them into the environment, and then `exec`s a target command with that environment. It is meant to wrap a process (e.g. in an ECS container `Command`) so secrets land in `process.env` of the child without touching disk.

## Runtime contract

The binary reads two environment variables and trailing CLI args:

- `SSM_KEYS` — comma-separated list of SSM parameter paths (e.g. `/PATH/KEY_1,/PATH/KEY_2`). Required; empty/undefined aborts.
- `AWS_REGION` (plus standard AWS credential env/config) — consumed by `config.LoadDefaultConfig`.
- Trailing args — the command to exec, e.g. `go_ssm node main.js`.

For each key the full SSM path is split on `/` and only the **last segment** becomes the exported env var name (`/foo/bar/DB_URL` is exported as `DB_URL`). The process is replaced via `syscall.Exec`, so the child inherits the augmented environment and go-ssm's PID.

## Architecture

- `main.go` — thin entry; calls `ssm.Run()`.
- `pkg/ssm.go` — the whole flow: parse flags, read `SSM_KEYS`, load AWS config, `GetParameter` (with decryption) per key, `os.Setenv`, resolve + `syscall.Exec` the target. Uses AWS SDK for Go **v2** (`aws-sdk-go-v2/config`, `aws-sdk-go-v2/service/ssm`). Errors are fatal via `log.Fatalf` — no partial-failure handling by design.
- `internal/utils/env.go` — `GetEnvStr` env lookup with optional default and an empty-var error.

## Commands

```bash
make build        # build ./bin/go_ssm
make build-linux  # static linux/amd64 cross-compile (CGO disabled)
go build -v ./... # what CI runs
go test -v ./...  # run all tests
go test -v -run TestName ./internal/utils/   # single test
```

## Gotchas

- **Makefile test targets are stale**: `make test` / `make test-cover` point at `./tests/...`, but tests now live under `internal/utils/`. Use `go test ./...` or fix the Makefile path before relying on them.
- **CI branch mismatch**: `.github/workflows/go.yml` triggers on `master`, but the default branch is `main`, so the workflow does not currently run on pushes/PRs. It also pins Go 1.17 while `go.mod` declares 1.27.
