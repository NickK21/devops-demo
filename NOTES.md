# devops-demo — Notes Log

This file tracks progress and decisions made during the DevOps assessment project.

---

## 2025-10-03 — Repo Setup, Go Module + API Stub
- Created repo `devops-demo`.
- Added README with initial project scope:
  - Small Go API returning my name and a timestamp
  - Plan to containerize with Docker
  - Plan to use GitHub Actions for CI/CD (apprentice-action)
  - Deploy to AWS
- Pushed initial README commit to GitHub.
- Initialized Go module with `go mod init`.
- Installed the Fiber framework.
- Added `main.go` with a simple API stub that returned "API stub working".
- Verified locally with `go run main.go` at http://localhost:80.

---

## 2025-10-04 — Base API Completed + Run Docs
- Implemented the root endpoint returning `"My name is Nick Kaplan"` and a current timestamp in JSON.
- Added support for the `PORT` environment variable (defaults to 80 for CI; configurable locally).
- Verified locally on a custom port using `PORT=8000 go run main.go`.
- Added `.gitignore` to exclude binaries and editor/OS files.
- Updated README with a **How to Run** section (default port, `PORT` override, and `curl` example).
- Pushed updates to GitHub.

---

## Next Steps
- Containerize the app with a clear, well-documented Dockerfile.
- Add a GitHub Actions workflow for automated build and verification (apprentice-action).
- Deploy the containerized app to AWS (evaluate Fargate, Elastic Beanstalk, or EC2).