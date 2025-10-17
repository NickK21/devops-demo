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

- CI/CD now runs end-to-end on every push to `main`:
  - Builds the Docker image and embeds the Git commit via Go `-ldflags` → JSON includes `"version":"<git-sha>"`.
  - Pushes the image to Docker Hub with two tags: `:<git-sha>-<run-number>` and `:latest`.
  - Deploys the SHA-tagged image to ECS Fargate (not `latest`) using the rendered task definition, then waits for steady state.
- Pinned `liatrio/github-actions/apprentice-action` to a specific commit SHA for deterministic tests; other GitHub Actions use stable version tags.
- Verified deploy:
  - Confirmed the task definition’s image equals the SHA tag:
    - `aws ecs describe-task-definition --task-definition <TD_ARN> --query 'taskDefinition.containerDefinitions[0].image'`
  - Confirmed the running task uses that same image:
    - `aws ecs describe-tasks --cluster devops-demo-cluster --tasks <TASK_ARN> --query 'tasks[0].containers[0].image'`
  - Retrieved the public IP from the task ENI and validated the live response:
    - `ENI_ID=$(aws ecs describe-tasks ... | jq -r '.tasks[0].attachments[] | select(.type=="ElasticNetworkInterface") | .details[] | select(.name=="networkInterfaceId") | .value')`
    - `PUB_IP=$(aws ec2 describe-network-interfaces --network-interface-ids "$ENI_ID" --query 'NetworkInterfaces[0].Association.PublicIp' --output text)`
    - `curl -s "http://$PUB_IP/" | jq .  # shows {"message","timestamp","version":"<git-sha>"}`
