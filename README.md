# go-websockets

## Project: Real-time Chat Application
Build a scalable WebSocket-based chat system with presence detection and message history.

Incremental learning using Claude Code in the terminal as a Guide, all code is researched and written by hand to ensure maximum practice, then the agent gives corrections and next steps.

## Project Overview


## Initialize dependencies in containers
```sh 
docker compose --profile dependencies up -d

docker compose --profile dependencies down
```

## Run app in local machine
```sh 
go run cmd/api/main.go
```

## Run both the app and dependencies in containers
```sh 
docker compose --profile full up --build -d

docker compose --profile full down
```