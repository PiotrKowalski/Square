# Square

## Environemnt Variables

- `GRPC_PORT` - port on which grpc service is (default: 30000)
- `REST_PORT` - port on which http service is (default: 40000)
- `SERVE_HTTP` - if it should serve HTTPS or HTTP (default: true)
- `ECHO_URL` - url used for echo calling (default: https://postman-echo.com/get)

## Initizalization 

1. Install TaskFile
  https://taskfile.dev/installation/

2. Run `go run .` to launch service

## Tasks

Example `task test`

- test - Runs unit tests
- coverage - Check coverage of code
- proto - Generates protobuf gateway files
- openapi - Generate documentation for API
