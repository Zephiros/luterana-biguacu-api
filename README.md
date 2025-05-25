# Luterana Biguacu API

A simple Go API for the Luterana Biguacu website.

## Docker Support

This project includes Docker support for easy deployment and development.

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/) (usually included with Docker Desktop)

### Building and Running with Docker

#### Using Docker Compose (Recommended)

1. Build and start the containers:

```bash
docker-compose up -d
```

2. The API will be available at http://localhost:8080

3. The PostgreSQL database will be available at localhost:5432

4. To stop the containers:

```bash
docker-compose down
```

5. To stop the containers and remove the database volume:

```bash
docker-compose down -v
```

#### Using Docker Directly

1. Build the Docker image:

```bash
docker build -t luterana-biguacu-api .
```

2. Run the container:

```bash
docker run -p 8080:8080 luterana-biguacu-api
```

### API Endpoints

- `GET /health` - Health check endpoint
- `POST /api/contact` - Submit contact form

#### YouTube Lives Endpoints
- `GET /api/youtube/lives` - Get YouTube live streams from the configured channel

#### Testimonials Endpoints
- `GET /api/testimonials` - Get all testimonials
- `GET /api/testimonials/:id` - Get testimonial by ID
- `POST /api/testimonials` - Create a new testimonial
- `PUT /api/testimonials/:id` - Update an existing testimonial
- `DELETE /api/testimonials/:id` - Delete a testimonial

## Configuration

The application uses environment variables for configuration. These can be set in a `.env` file in the root directory.

### Environment Variables

- `YOUTUBE_API_KEY` - Your YouTube API key for accessing the YouTube API
- `YOUTUBE_CHANNEL_ID` - The ID of your YouTube channel to fetch live streams from
- `PORT` - The port the server will listen on (default: 8080)
- `GIN_MODE` - The Gin framework mode (debug or release)

### Database Configuration

The application uses PostgreSQL as its database. The following environment variables are used for database configuration:

- `DB_HOST` - The hostname of the PostgreSQL server (default: localhost)
- `DB_PORT` - The port of the PostgreSQL server (default: 5432)
- `DB_USER` - The username for the PostgreSQL server (default: postgres)
- `DB_PASSWORD` - The password for the PostgreSQL server (default: postgres)
- `DB_NAME` - The name of the PostgreSQL database (default: luterana)

## Development

### Running Locally (Without Docker)

1. Install Go (version 1.16 or later)
2. Install PostgreSQL (version 12 or later)
3. Clone the repository
4. Install dependencies:

```bash
go mod download
```

5. Create a PostgreSQL database:

```bash
createdb luterana
```

6. Configure the database connection in the `.env` file:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_postgres_username
DB_PASSWORD=your_postgres_password
DB_NAME=luterana
```

7. Run the application:

```bash
go run main.go
```

The API will be available at http://localhost:8080
