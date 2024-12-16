# Chill: A Social Network for Weirdos

## Features

- **User Management**:
  - User registration, login, and token-based authentication.
  - Profile management.
  - Follow/unfollow functionality.
  - User activation through invitation tokens.

- **Posts**:
  - Create, update, delete, and fetch posts.
  - Support for tags and metadata.

- **Feeds**:
  - Fetch user-specific feeds.
  - Filtering, sorting, and pagination support.

- **Operations**:
  - Healthcheck endpoint for system monitoring.

---

## API Overview

Base Path: `/v1`

### Authentication
- **Create Token**: `/authentication/token` (POST)
  - Generates a token for a registered user.
- **Register User**: `/authentication/user` (POST)
  - Registers a new user.

### Posts
- **Create Post**: `/posts` (POST)
  - Creates a new post.
- **Fetch Post**: `/posts/{id}` (GET)
  - Retrieves a post by its ID.
- **Update Post**: `/posts/{id}` (PATCH)
  - Updates a post by its ID.
- **Delete Post**: `/posts/{id}` (DELETE)
  - Deletes a post by its ID.

### Users
- **Fetch Profile**: `/users/{id}` (GET)
  - Retrieves a user profile by ID.
- **Follow User**: `/users/{userID}/follow` (PUT)
  - Follows a user by ID.
- **Unfollow User**: `/users/{userID}/unfollow` (PUT)
  - Unfollows a user by ID.
- **Activate User**: `/users/activate/{token}` (PUT)
  - Activates a user using an invitation token.

### Feed
- **Fetch User Feed**: `/users/feed` (GET)
  - Retrieves the user’s feed with filters, sorting, and pagination options.

### Operations
- **Healthcheck**: `/health` (GET)
  - Verifies the system’s health status.

---

## Prerequisites

1. **Docker**: Ensure Docker is installed on your machine.
2. **Docker Compose**: Required to orchestrate services.
3. **Go**: If running locally, ensure Go (>=1.22) is installed.
4. **PostgreSQL**: Database for persistent storage.
5. **Redis**: Cache and session management.

---

## Installation and Setup

### Using Docker Compose
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd chill
   ```

2. Start the services:
   ```bash
   docker-compose up --build
   ```

3. Access the following services:
   - API: [http://localhost:8080](http://localhost:8080)
   - Redis Commander: [http://localhost:8081](http://localhost:8081)

4. Run database migrations (if required):
   ```bash
   make migrate-up
   ```

### Running Locally
1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

3. Access the API: [http://localhost:8080](http://localhost:8080)

---

## Environment Variables

The application requires the following environment variables:

| Variable           | Description                              | Default Value         |
|--------------------|------------------------------------------|-----------------------|
| `DB_ADDR`          | Database connection string              | `postgres://...`      |
| `REDIS_ADDR`       | Redis connection string                 | `redis:6379`          |
| `JWT_SECRET`       | Secret key for JWT tokens               | `your-secret-key`     |

---

## Database Schema

The database includes the following key entities:

- **Users**:
  - Fields: `id`, `email`, `username`, `password`, `created_at`, `role_id`, etc.
- **Posts**:
  - Fields: `id`, `title`, `content`, `created_at`, `updated_at`, `user_id`, etc.
- **Comments**:
  - Fields: `id`, `content`, `post_id`, `user_id`, `created_at`, etc.

---

## Testing

Run tests using the Makefile:
```bash
make test
```

