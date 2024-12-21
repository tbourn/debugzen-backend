# DebugZen Backend

DebugZen Backend is a Go-based API server for reviewing and analyzing code snippets using OpenAI's API.

## Features

- Analyzes code for errors, best practices, and improvements.
- RESTful API using Gin framework.
- Structured feedback with titles and detailed descriptions.
- Integration with OpenAI's GPT-3.5-turbo.

---

## Architecture

- **Handlers:** Defines the API endpoints.
- **Services:** Handles business logic and OpenAI interaction.
- **Utils:** Contains reusable utilities (e.g., error handling).
- **Config:** Handles environment variables.

---

## Prerequisites

1. **Go**: Install [Go](https://golang.org/doc/install).
2. **OpenAI API Key**: Obtain an API key from OpenAI.
3. **Environment Variables**: Create a `.env` file in the root:
   - OPENAI_API_KEY=your-api-key
   - PORT=app-port, defaults to 8080
   - BASE_URL=frontend-base-URL, defaults to http://localhost:5173

---

## Setup Instructions

### Clone the Repository

```bash
git clone https://github.com/tbourn/debugzen.git
cd debugzen
go mod tidy
go run .
```

## API Endpoints

### POST /review

Analyzes a given code snippet and provides structured feedback.

### Request

```json
{
  "code": "def hello_world(): print('Hello, World!')"
}
```

### Response

```json
{
  "feedback": [
    {
      "title": "Correct Use of Functions",
      "description": "This function is well-defined and prints a message as expected."
    },
    {
      "title": "Code Readability",
      "description": "Consider adding comments to explain the purpose of this function."
    }
  ]
}
```

## Testing

```bash
go test ./... -v
```
