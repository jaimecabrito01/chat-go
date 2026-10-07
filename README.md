# Go Chat

Chat bidirecional desenvolvido em Go utilizando WebSockets.

O projeto tem como objetivo construir, inicialmente, um **MVP de chat privado entre usuários**, com persistência de usuários e mensagens. A aplicação será evoluída posteriormente para suportar comunicação P2P e criptografia de ponta a ponta.

## Objetivo

O projeto começou como um protótipo simples de comunicação via WebSocket, no qual todas as mensagens eram transmitidas para todos os clientes conectados.

A próxima evolução é transformar esse broadcast global em um sistema de **chats privados**, onde uma mensagem enviada por um usuário seja entregue somente aos participantes daquele chat.

### Fluxo inicial

```text
User A
   │
   │ WebSocket
   ▼
Go Server
   │
   │ roteamento da mensagem
   ▼
Chat
   │
   └── User B
```

O servidor será responsável inicialmente por intermediar as mensagens e armazenar o histórico.

---

## Tecnologias

- Go
- WebSocket
- Gorilla WebSocket
- GORM
- SQLite

Tecnologias que poderão ser adicionadas futuramente:

- PostgreSQL
- WebRTC
- STUN/TURN
- Criptografia de ponta a ponta

---


## Estrutura inicial

```text
chat-go/
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── database/
│   │   └── database.go
│   │
│   ├── user/
│   │   ├── model.go
│   │   ├── repository.go
│   │   └── service.go
│   │
│   ├── chat/
│   │   ├── model.go
│   │   └── manager.go
│   │
│   └── websocket/
│       ├── client.go
│       ├── handler.go
│       └── hub.go
│
│
├── data/
│   └── data.db
│
├── go.mod
└── go.sum
```

