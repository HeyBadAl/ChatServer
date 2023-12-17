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

## TODO 

1. **Improve error handling and return appropriate HTTP status codes in handlers:**
   - [ ] Enhance error handling in each handler function to provide meaningful error messages.
   - [ ] Implement consistent use of HTTP status codes for different scenarios (e.g., 200 for success, 400 for bad request, 404 for not found, 500 for server errors).
   - [ ] Include additional context in error responses, such as error codes or descriptions, to assist clients in understanding and responding to errors.

2. **Implement unit tests for critical parts of the application:**
   - [ ] Create unit tests for each handler function to validate their behavior under various conditions.
   - [ ] Cover critical utility functions with unit tests to ensure their correctness.
   - [ ] Use testing frameworks like Go's built-in `testing` package or external packages like `testify` for comprehensive test coverage.

3. **Add configuration options for the server:**
   - [ ] Implement a configuration mechanism, allowing users to customize server settings.
   - [ ] Consider using environment variables, a configuration file, or command-line flags for configuring parameters like the server port, WebSocket settings, and other relevant options.

4. **Assess and implement security best practices:**
   - [ ] Conduct a security review of the application, identifying potential vulnerabilities.
   - [ ] Implement input validation to prevent common security issues, such as injection attacks or malformed requests.
   - [ ] Ensure secure connections, especially when dealing with sensitive information.

5. **Document the code:**
   - [ ] Provide clear and concise comments for complex or critical sections of code.
   - [ ] Update the documentation to reflect any changes made to the codebase.
   - [ ] Consider generating documentation using tools like GoDoc for an automated and consistent documentation process.

6. **Set up a continuous integration (CI) pipeline:**
   - [ ] Configure a CI pipeline to automate the testing and building processes.
   - [ ] Integrate the CI pipeline with version control (e.g., GitHub Actions, GitLab CI) for seamless integration.
   - [ ] Ensure that the CI pipeline includes running tests, building the application, and possibly deploying to staging environments for thorough validation.

7. **Consider container orchestration tools for deployment:**
   - [ ] Explore container orchestration tools like Kubernetes or Docker Compose for managing and scaling the application in production.
   - [ ] Investigate deployment strategies, load balancing, and recovery mechanisms provided by container orchestration platforms.

8. **Conduct code reviews:**
   - [ ] Establish a code review process within the team to improve code quality.
   - [ ] Encourage team members to review each other's code, providing constructive feedback.
   - [ ] Use code review tools integrated with version control systems for a streamlined process.

# Contributing

Feel free to contribute to this project by opening an issue or submitting a pull request.

Let's make this project better!

