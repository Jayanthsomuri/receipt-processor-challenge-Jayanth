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
```
2. Install Dependencies
If you're not using Docker, you can install the dependencies using the following Go command:
```bash
go mod tidy
```
3. Run the Application
You can run the application in two ways:

Option 1: Using Go Directly
To run the application directly using Go:
```bash
go run main.go
```
This will start the server on http://localhost:8080.


Endpoints
Process Receipt
POST /receipts/process

Description: Process a receipt by submitting a JSON payload with retailer, purchase date, purchase time, items, and total amount.
Request Body Example:
json
Copy
Edit

```bash
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
```
Response Example:
```bash
{
  "id": "042b4191-8a33-4c26-bfcb-703a58b36cd8"
}

```
Docker Setup (Optional)
If you're using Docker, the following command will build and run the application inside a Docker container:
```bash
docker build -t receipt-processor .
docker run -p 8080:8080 receipt-processor
```
This project is open-source and available under the MIT License.
