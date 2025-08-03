# hsi-sandbox-rest
HSI sanbox REST API template

# Getting Started
## How to run
1. Clone the repository
2. Install dependencies
   ```bash
   go mod tidy
   ```
3. Run the application
   ```bash
    go run main.go
    ```
4. Access the API at `http://localhost:55667`
5. Check if service is running
   ```bash
   curl http://localhost:55667/ping
   ```
   
## pre-requisites
- Go 1.20 or later
- Insomnia or Postman for testing the API
  - [Insomnia](https://insomnia.rest/download)
  - [Postman](https://www.postman.com/downloads/)