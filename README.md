# devops-demo

Simple Go API that returns my name and a timestamp.  
Built with Fiber for routing and designed to be containerized and deployed on AWS.

## Features
- Lightweight Go API (Fiber framework)
- JSON response with `message` and `timestamp`
- Environment-based port configuration (`PORT` variable)
- Plans for Docker containerization and CI/CD with GitHub Actions
- Cloud deployment (AWS)

## How to Run

Default run (port 80, may require admin privileges):

```bash
go run main.go
```
Visit http://localhost/

Running with custom port:

```bash
PORT=8000 go run main.go
```
Visit http://localhost:8000/

To test with curl:
```
curl -i http://localhost:8000/
```
Expected output:
```
{
  "message": "My name is Nick Kaplan",
  "timestamp": 1759726241 
}
```
## Run in Docker

Build the image:
```bash
docker build -t devops-demo:local .
```
Run the container (host port 8000 → container port 80):
```bash
docker run --rm -p 8000:80 --name devops-demo devops-demo:local
```
http://localhost:8000/

Expected output:
```bash
{
  "message": "My name is Nick Kaplan",
  "timestamp": 1759726241
}
```