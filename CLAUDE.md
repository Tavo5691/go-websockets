# Go Learning Journey

  ## Context
  I'm a backend engineer learning Go. Guide me with hints rather than writing code for me unless explicitly asked.

  ## Code Review Expectations
  1. Explain WHY something is better, not just WHAT to change
  2. Point me to official Go documentation when relevant
  3. Show examples of idiomatic Go patterns
  4. Help me understand race conditions and concurrency issues

  ## Completed: Project 1 - Task API

  ### Concepts Mastered
  - **Go fundamentals**: packages, imports, exports (capitalization), struct tags
  - **Types**: structs, slices, maps, pointers vs values
  - **Error handling**: explicit error returns, `error` interface, nil checks
  - **HTTP with Gin**: handlers, routing, path params, route groups, middleware
  - **JSON**: encoding/decoding, struct tags (`json`, `binding`), `omitempty`
  - **Database**: `database/sql`, parameterized queries, `Scan()`, `Query()`, `QueryRow()`, `Exec()`
  - **Authentication**: bcrypt password hashing, JWT creation/validation
  - **Project structure**: `cmd/`, `internal/`, handler structs, dependency injection
  - **Environment config**: godotenv, `os.Getenv()`
  - **Docker**: multi-stage builds, docker-compose

  ### Patterns Learned
  - Handler struct with dependencies (idiomatic DI)
  - Middleware as closure capturing config
  - Validation with struct binding tags
  - Structured error responses
  - Config loading at startup with validation

  ## Project Structures Used So Far
  cmd/api/main.go           # Wiring only (~50 lines)
  internal/
  ├── config/config.go      # Config struct + Load()
  ├── handlers/             # HTTP handlers as methods on Handler struct
  ├── middleware/           # Auth middleware
  └── models/               # Data structures

  ## Current Focus: Project 2

  ### Learning Objectives:
  Implement WebSocket connections and message broadcasting
  Manage concurrent connections with goroutines
  Use channels for message routing between goroutines
  Implement Redis Pub/Sub for horizontal scaling
  User presence detection (online/offline status)
  Message persistence and history retrieval
  Handle connection recovery and reconnection logic

  ### Features to Implement:
  One-on-one messaging
  Group chats/rooms
  Typing indicators
  Message read receipts
  Online user list
  Message history with pagination

  ## Remaining Learning Goals
  - Goroutines and channels (not yet covered)
  - Testing (`go test`, mocking, table-driven tests)
  - Interfaces (beyond basic usage)
  - Context package for cancellation
  - Race condition detection

  ## Preferences
  - Guide with hints first, give solutions only when asked
  - Explain the "why" behind Go idioms
  - Point out security issues and performance concerns
  - Challenge me to think through problems
  - If needed, you can propose to generate simple Frontends for manually testing the components/final app
  - Since we already covered Docker, we should start every project by laying the dockerization baseground 
  - Depencies should be able to run in containers (DB, Redis, Queues) and the app to be run both locally and in a container