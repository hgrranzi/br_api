# br_api

A simple REST API for managing Brainee resources.

## Endpoints

### Create a Brainee

- **URL:** `/brainees`
- **Method:** `POST`
- **Body:**

  ```json
  {
    "text": "Your text",
    "author": "Author name",
    "brand": "Brand name"
  }
  ```

### Get a Brainee by ID

- **URL:** `/brainees/{braineeId}`
- **Method:** `GET`

### Get All Brainees

- **URL:** `/brainees`
- **Method:** `GET`

## Configuration
If you have a PostgreSQL database, specify the connection details in the [.env file](./.env). <br>
If no database is configured, the application will store data in-memory.

## Build and Run

```
git clone https://github.com/hgrranzi/br_api.git
cd br_api
go mod tidy
go build -o brainee ./cmd/app
./brainee
```
The server will start on port 8080.