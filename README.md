# Simple Chat Server with Webhooks in Go

This project demonstrates a basic chat server implementation in Go featuring webhooks for real-time communication. Users can connect to the server via WebSocket to receive updates about new message, and messages can be sent to the chat through a RESTful API endpoint. 

## Getting Started

Follow the steps below to run the chat server locally:

1. Clone the repository
```bash
git clone https://github.com/HeyBadAl/ChatServer
```

2. Navigate to the project directory
```bash
cd ChatServer
```

3. Build and run the docker 
```bash
docker build -t chatserver .
docker run -p 8080:8080 chatserver
```
The server will start running at `http://localhost:8080`.

# API Endpoints

## WebSockets Endpoint

Connect to this endpoint using WebSockets to receive real-time updates about new messages.

- Endpoint: `/webhook`

## Send Message Endpoint

Send a message to the chat through this RESTful endpoint.

- Endpoint: `/send`
- Method: POST 
- Request Body: 

```json
{
    "user": "user",
    "content": "Hello, World!"
}
```

# Example Usage 

1. Connect to the WebSockets endpoint using a WebSocket client or a compatible tool.
2. Send a message using the `/send` endpoint.

```bash
curl -X POST -d '{"user": "user", "content": "Hello, World!"}' http://localhost:8080/send
```
This message will be broadcasted to all connected clients.


# Running without docker 

If you prefer to run the server without docker, you can use the following command to start the server:
```bash
go run .
```

# Contributing

Feel free to contribute to this project by opening an issue or submitting a pull request.

# License

MIT License

Let's make this project better!

