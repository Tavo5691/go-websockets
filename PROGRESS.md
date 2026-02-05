# Project 2 Progress: Real-time Chat Application

## Completed Phases

### Phase 1: Foundation & Docker Setup
- [x] docker-compose.yml with PostgreSQL and Redis (using profiles)
- [x] Project structure (cmd/api, internal/config, handlers, hub, models, middleware)
- [x] Environment config with validation
- [x] Health endpoint
- [x] Dockerfile with multi-stage build
- [x] Local + Docker workflow working

### Phase 2: Goroutines & Channels Fundamentals
- [x] Learned goroutine lifecycle with sync.WaitGroup (including new Go() method)
- [x] Understood unbuffered vs buffered channels
- [x] Built basic echo WebSocket server
- [x] Practiced channel blocking behavior

### Phase 3: Connection Hub & Broadcasting
- [x] Hub pattern with channels (register, unregister, broadcast)
- [x] Client struct with read/write loops (separate goroutines)
- [x] Hub.Run() with select statement
- [x] Broadcast messages to all connected clients

### Phase 4: User Authentication & Private Messaging
- [x] JWT middleware for WebSocket auth (query param token)
- [x] User identity tied to connections (userId in Client)
- [x] Message struct with JSON tags (From, To, Content, Timestamp)
- [x] Direct messaging routing (Hub looks up by userId)
- [x] Dev token endpoint for testing (/dev/token)

### Phase 5: Rooms/Groups
- [x] Room struct with members map
- [x] MessageType "enum" (dm, room, join, leave)
- [x] Hub tracks rooms (map[uuid.UUID]*Room)
- [x] Join/Leave room functionality (auto-create on first join)
- [x] Route messages to room members only
- [x] Clean up room memberships on disconnect
- [x] Updated test HTML with room support

---

## Current Phase: Phase 6 - Presence & Typing Indicators

### Status: Just Started

### What's Left to Implement

**Step 1: Add new message types**
```go
const (
    // ... existing types
    Typing  MessageType = "typing"   // user is typing
    Online  MessageType = "online"   // user came online
    Offline MessageType = "offline"  // user went offline
)
```

**Step 2: Broadcast presence on connect/disconnect**
- When client registers → notify room members they're online
- When client unregisters → notify room members they're offline
- Consider: client has no rooms on initial register (handle global vs room presence)

**Step 3: Handle typing indicators**
- DM typing: `{"type": "typing", "to": "user-id"}`
- Room typing: `{"type": "typing", "room": "room-id"}`
- Route to recipient (DM) or room members (room)

**Step 4: Update HTML test page**
- Show online users
- Show typing indicators
- Handle debounce on client side

### Design Decisions Made
- Presence updates: Room members only (more scalable than global)
- Typing scope: DM → recipient only, Room → room members
- Typing debounce: Client stops sending after ~2s, recipient auto-clears after ~3s

---

## Remaining Phases

### Phase 7: Message Persistence & History
- Store messages to PostgreSQL
- Retrieve message history with cursor-based pagination
- Async message persistence (don't block sender)

### Phase 8: Redis Pub/Sub for Horizontal Scaling
- Publish messages to Redis channel
- Subscribe and fan-out to local connections
- Handle Redis connection failures

### Phase 9: Connection Recovery & Read Receipts
- Track "last seen" message per user
- Resume from disconnect point
- Implement read receipts

### Phase 10: Testing & Race Detection
- Unit test Hub logic
- Test concurrent operations with `go test -race`
- Mock WebSocket connections
- Integration tests with real Redis

---

## Key Files

```
cmd/api/main.go              # Wiring: hub, handlers, routes
internal/
├── config/config.go         # Config struct + Load()
├── handlers/
│   ├── handlers.go          # Handler struct with hub + jwtKey
│   ├── health.go            # Health endpoint
│   ├── ws.go                # WebSocket upgrade + auth
│   └── dev.go               # Dev token generator
├── hub/
│   ├── hub.go               # Hub, Client, ReadLoop, WriteLoop
│   └── room.go              # Room struct
├── middleware/
│   └── auth.go              # JWT validation middleware
└── models/
    ├── message.go           # Message struct + MessageType enum
    └── responses.go         # API response structs
web/
└── index.html               # Test UI for WebSocket chat
```

---

## Concepts Learned So Far

- Goroutines and channels for concurrency
- sync.WaitGroup (traditional and new Go() method)
- Hub pattern for managing WebSocket connections
- Channel-based communication (no locks needed)
- select statement for multiplexing channels
- WebSocket upgrade from HTTP
- JWT auth on WebSocket connections
- Gin context for passing data between middleware and handlers
- Message routing (broadcast, DM, rooms)
- Go "enum" pattern with typed string constants
