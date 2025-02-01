# Receipt Processor Challenge

This project is a receipt processor API that allows you to process receipts and calculate reward points based on various criteria such as retailer name, purchase date, total amount, and items in the receipt.

## Requirements

Before running the project, make sure you have the following installed:

- **Go** (version 1.20 or higher)
- **Docker** (optional, if you prefer to run it inside a container)

## Setup and Run

### 1. Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/Jayanthsomuri/receipt-processor-challenge-Jayanth.git
cd receipt-processor-challenge-Jayanth
2. Install Dependencies
If you're not using Docker, you can install the dependencies using the following Go command:
```
bash
Copy
Edit
go mod tidy
This will download and install the necessary dependencies (github.com/google/uuid, github.com/gorilla/mux).

3. Run the Application
You can run the application in two ways:

Option 1: Using Go Directly
To run the application directly using Go:

bash
Copy
Edit
go run main.go
This will start the server on http://localhost:8080.

Option 2: Using Docker (Optional)
Alternatively, you can build and run the application using Docker.

Build the Docker image:
bash
Copy
Edit
docker build -t receipt-processor .
Run the Docker container:
bash
Copy
Edit
docker run -p 8080:8080 receipt-processor
This will also start the server on http://localhost:8080.

Endpoints
Process Receipt
POST /receipts/process
Description: Process a receipt by submitting a JSON payload with retailer, purchase date, purchase time, items, and total amount.

Request Body Example:

json
Copy
Edit
{
  "retailer": "Target",
  "purchaseDate": "2024-02-01",
  "purchaseTime": "14:35",
  "items": [
    {
      "shortDescription": "Pepsi",
      "price": "1.99"
    }
  ],
  "total": "1.99"
}
Response Example:

json
Copy
Edit
{
  "id": "042b4191-8a33-4c26-bfcb-703a58b36cd8"
}
This response will include a unique ID for the processed receipt.

Get Points
GET /receipts/{id}/points
Description: Get the reward points for a specific receipt by providing the receipt ID.

Request Example:

bash
Copy
Edit
GET http://localhost:8080/receipts/042b4191-8a33-4c26-bfcb-703a58b36cd8/points
Response Example:

json
Copy
Edit
{
  "points": 100
}
The response will return the calculated reward points based on the receipt's details.

Testing the API
You can test the API using curl or Postman.

Testing Process Receipt
bash
Copy
Edit
curl -X POST "http://localhost:8080/receipts/process" -H "Content-Type: application/json" -d '{ "retailer": "Target", "purchaseDate": "2024-02-01", "purchaseTime": "14:35", "items": [ { "shortDescription": "Pepsi", "price": "1.99" } ], "total": "1.99" }'
This will return the receipt ID.

Testing Get Points
Once you have the receipt ID (e.g., 042b4191-8a33-4c26-bfcb-703a58b36cd8), use it to get the points:

bash
Copy
Edit
curl "http://localhost:8080/receipts/042b4191-8a33-4c26-bfcb-703a58b36cd8/points"
This will return the points for the receipt.

Docker Setup (Optional)
If you're using Docker, the following commands will build and run the application inside a Docker container:

bash
Copy
Edit
docker build -t receipt-processor .
docker run -p 8080:8080 receipt-processor
The application will be available at http://localhost:8080.

License
This project is open-source and available under the MIT License.

markdown
Copy
Edit

### Key Improvements:
- **Code formatting**: The instructions are clearer with appropriate code blocks for commands and examples.
- **Clarity**: Clean and structured formatting for commands and examples.
- **Simplified text**: Improved readability for Docker setup and API request/response examples.

Feel free to copy and paste it into your `README.md` file in your GitHub repo! Let me know 
