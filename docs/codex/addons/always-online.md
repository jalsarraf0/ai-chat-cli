This document augments the baseline Codex Rules.

0. Internet Access Policy
------------------------
1. Assume continuous Internet connectivity in all local scripts, helper tools, and CI jobs.
2. Whenever a dependency is missing or out-of-date you must run:
   go get -u ./...
   go mod tidy -e -v
3. If a third-party CLI is absent, install it via go install <module>@latest.

1. Coverage Gate ≥ 83%
----------------------
Use go test -race -covermode=atomic -coverprofile=coverage.out ./...
Then ensure coverage ≥ 83%.

2. Conflict-Free Pull Requests
------------------------------
Rebase onto dev before opening or updating a PR and ensure no merge conflicts.

3. CI/CD Compliance
--------------------
You may edit the single workflow file .github/workflows/ci.yml to extend jobs but must keep stage order.

4. Quick Pre-PR Checklist
-------------------------
- go vet ./... passes
- go test -race ./... ≥ 83% coverage
- go mod tidy -e leaves no changes
- Docs regenerated (make docs) with no diff noise
- shellcheck scripts/*.sh passes (ignore SC2086 where documented)
- Branch rebased onto dev – no conflicts
- Commit message ≤ 72 chars, imperative; PR labelled codex-generated
