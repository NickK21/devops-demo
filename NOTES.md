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

## 2025-10-10 — Containerization Completed
- Added `.dockerignore` to keep image lean (exclude binaries, editor files, etc.).
- Created a multi-stage `Dockerfile`:
  - Uses `golang:1.22-alpine` for build stage.
  - Compiles a static binary with `CGO_ENABLED=0` for portability.
  - Uses `gchr.io/distroless/static` for minimal runtime.
  - Exposes port 80.
- Successfully built and tested container locally:
  - `docker build -t devops-demo:local .`
  - `docker run --rm -p 8000:80 devops-demo:local`
  - Verified JSON response at `http://localhost:8000/`.
- Committed Dockerfile and .dockerignore to GitHub.

---

## 2025-10-12 — CI diagnostics + ECS deploy (public IP)

- Added/updated CI workflow (`.github/workflows/ci.yml`):
  - Build image with Docker Buildx.
  - Run container on port 80 and perform diagnostics:
    - Fetch `/` and verify `Content-Type: application/json`.
    - Enforce exact minified body (`diff` vs `jq -cj .` and `JSON.stringify(JSON.parse(body))`).
  - Run `liatrio/github-actions/apprentice-action` (final “minified JSON” test is a stub; diagnostics cover it).
  - Login to Docker Hub and push image with tags `<git-sha>-<run-number>` and `latest`.
- Created ECS task definition on Fargate:
  - Image: `<dockerhub-username>/devops-demo:latest`
  - Port 80, assign public IP, SG allows inbound TCP/80 for the demo.
- Deployed ECS service (no ALB) and verified with `curl` against the public IP:
  - Example: `{"message":"My name is Nick Kaplan","timestamp":1760326289394}`

---

## 2025-10-16 — Final CI/CD Automation + ECS Verification

- Confirmed working CI/CD pipeline end-to-end:
  - Builds, tests, and pushes Docker image to Docker Hub (`nickkap/devops-demo`).
  - ECS Fargate task automatically redeploys via GitHub Actions.
- Pinned `liatrio/github-actions/apprentice-action` to a specific commit SHA for deterministic builds.
- Updated workflow (`ci.yml`) to:
  - Embed `git SHA` via Go `-ldflags` (included as `"version"` field in JSON output).
  - Push image to Docker Hub with tags `:<git-sha>-<run-number>` and `:latest`.
  - Force ECS service update and wait for steady state.
- Verified ECS task revision update using:
  - `aws ecs list-tasks` and `aws ecs describe-tasks`
  - Confirmed new task running on `devops-demo-cluster` with container `devops-demo:latest`.
- Retrieved new public IP (`18.246.227.230`) from ENI details and verified live API response:

```bash
curl -i http://18.246.227.230/

# HTTP/1.1 200 OK
# {"message":"My name is Nick Kaplan","timestamp":<ms>,"version":"<git-sha>"}
```