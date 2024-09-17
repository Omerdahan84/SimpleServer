Note: This project is a work in progress.

Overview
This Go-based server handles requests, readiness checks, and JSON parsing, with integrated SQL database support.

Key Features
Error Handling: Graceful error handling using Go best practices.
Readiness Probe: A readiness probe to check the application's health, ensuring that it's ready to serve traffic.
JSON Parsing: Robust parsing of JSON requests using encoding/json.
SQL Integration: Interaction with an SQL database is configured using sqlc for structured queries.
File Structure
main.go: Entry point of the application, where the server is initialized and routes are defined.
handler_err.go: Contains custom error-handling logic to gracefully manage various failure scenarios.
handler_readiness.go: Implements a readiness probe to check the health of the server.
jason.go: Handles JSON parsing, including decoding and encoding of request and response objects.
rssagg: Contains RSS aggregation logic (if applicable to the project).
sqlc.yaml: Configuration for SQL code generation with sqlc.
go.sum: Contains the Go module dependencies for the project.
