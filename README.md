# devops-demo

![ci](https://github.com/NickK21/devops-demo/actions/workflows/ci.yml/badge.svg)

Simple Go API that returns my name and a timestamp.  
Built with Fiber for routing and designed to be containerized and deployed on AWS.

---

## Features

- Lightweight Go API (Fiber framework)
- `JSON`: `{"message":"My name is Nick Kaplan","timestamp":<ms>,"version":"<git-sha>"}`
- Minified response (no spaces / no trailing newline)
- Port configurable via `PORT` env var (defaults to 80)
- `Dockerfile` for multi-stage build → distroless final image
- GitHub Actions: build, verify, push to Docker Hub (+ diagnostic minified `JSON` check)
- Deployed on AWS ECS Fargate with a public IP

---

## Links

- **Repository:** https://github.com/NickK21/devops-demo
- **GitHub Actions (CI):** https://github.com/NickK21/devops-demo/actions/workflows/ci.yml
- **Docker Hub image:** https://hub.docker.com/r/nickkap/devops-demo
- **Live service (ECS):** http://18.246.227.230/

---

## How to Run

**Default run (port 80, may require admin privileges):**

```bash
go run main.go
# Visit http://localhost/
```

**Running with custom port:**

```bash
PORT=8000 go run main.go
# Visit http://localhost:8000/
```

**Expected output:**

```json
{"message":"My name is Nick Kaplan","timestamp":1760150974123,"version":"<git-sha>"}
```

**To test with curl:**

```bash
# Pretty Print
curl -s http://localhost/ | jq

# Has required fields
curl -s http://localhost/ | jq -e 'has("message") and has("timestamp")'

# Body is exactly minified
BODY=$(curl -s http://localhost/)
diff <(echo -n "$BODY") <(echo -n "$BODY" | jq -cj .)
```

---

## Run in Docker

**Build the image:**

```bash
docker build -t devops-demo:local .
```
**Run the container (host port 8000 → container port 80):**

```bash
docker run --rm -p 8000:80 --name devops-demo devops-demo:local
```
**Test:**

```bash
curl -s http://localhost:8000/ | jq
```

---

## CI/CD (GitHub Actions)

**Pipeline steps:**

- Build the image from the `Dockerfile`.
- Spin up the container and run diagnostics:
  - `GET /` must return `Content-Type: application/json`
  - Verify the body is exactly minified:
    - `diff` against `jq -cj .`
    - `JSON.stringify(JSON.parse(body)) === body` in Node
- Run Liatrio’s apprentice-action suite.
- Push the image to Docker Hub with two tags:
  - :`<git-sha>-<run-number>`
  - :`latest`
- The workflow pins liatrio/github-actions/apprentice-action to commit # `7208146...`

---

## AWS Deployment (ECS Fargate)

- **Cluster/Service Group:** `devops-demo-cluster` / `devops-demo-svc`
- **Task family/Container:** `devops-demo` / `devops-demo` (port 80)
- **Network:** Public subnets; task assigned a public IP
- **Security Group:** Inbound TCP/80 (open for demo)

### Verify Live Service:

```bash
curl http://18.246.227.230/
```

**Expected Response**

```json
{"message":"My name is Nick Kaplan","timestamp":<ms>,"version":"<git-sha>"}
```