# devops-demo — Notes Log

This file tracks progress and decisions made during the DevOps assessment project.

---

## 2025-10-03 — Repo Setup, Go Module + API Stub
- Created repo `devops-demo`.
- Added README with project scope:
  - Small Go API returning my name + timestamp
  - Dockerized for consistent builds
  - GitHub Actions CI/CD with apprentice-action
  - Deployment to cloud (AWS)
- Pushed initial README commit to GitHub.
- Initialized Go module with `go mod init`.
- Installed Fiber framework.
- Added `main.go` with a simple API stub that returns "API stub working".
- Verified locally with `go run main.go` at http://localhost:80.

---

## Next Steps
- Add a JSON endpoint that returns `{ "message": "Nick Kaplan", "timestamp": <unix> }`.
- Support `PORT` via environment variable (default to 80 for CI; use 8080 locally).
- Add a `.gitignore` to exclude binaries/build artifacts.
- Update README with “how to run” (local) and a short note about the `PORT` env var.