# devops-demo

![ci](https://github.com/NickK21/devops-demo/actions/workflows/ci.yml/badge.svg)

Simple Go API that returns my name and a timestamp.  
Built with Fiber for routing and designed to be containerized and deployed on AWS.

---

## Features

- Lightweight Go API (Fiber framework)
- `JSON`: `{"message":"My name is Nick Kaplan","timestamp":<ms>}`
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
- **Live service (ECS):** http://18.246.75.48/

---

## How to Run

**Default run (port 80, may require admin privileges):**

```bash
go run main.go
```
Visit http://localhost/

**Running with custom port:**

```bash
PORT=8000 go run main.go
```
Visit http://localhost:8000/

**Expected output:**

```json
{
  "message": "My name is Nick Kaplan",
  "timestamp": 1760150974123 
}
```

**To test with curl:**

```bash
curl -s http://localhost/ | jq
curl -s http://localhost/ | jq -e 'has("message") and has("timestamp")'
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
- Run diagnostics against a local container:
  - Request `http://localhost/`
  - Verify Content-Type is `application/json`
  - Verify the body is exactly minified:
    - `diff` against `jq -cj .`
    - `JSON.stringify(JSON.parse(body)) === body` in Node
- Run Liatrio’s apprentice-action (first six tests pass; “minified JSON” test is stubbed).
- Push the image to Docker Hub with two tags:
  - :`<git-sha>-<run-number>`
  - :`latest`

---

## AWS Deployment (ECS Fargate)

- **Platform:** ECS on Fargate (1 task, no load balancer for the demo)
- **Network:** Public subnets; task assigned a public IP
- **Ingress:** Allow HTTP (TCP/80) for demo purposes

### Test

```bash
curl http://18.246.75.48/
```

**Expected Response**

```json
{"message":"My name is Nick Kaplan","timestamp":1760326289394}
```