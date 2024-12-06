# E-Commerce Backend Documentation

## Project Structure
bash
backend/
├── cmd/
│ └── api/ # Application entry point
│ └── main.go
├── internal/
│ ├── config/ # Server configuration
│ │ └── server.go
│ ├── core/
│ │ └── ports/ # Interface definitions
│ │ └── repositories.go
│ ├── data/ # Sample data
│ │ └── sample_products.go
│ ├── handlers/ # HTTP handlers
│ │ ├── cart_handler.go
│ │ ├── order_handler.go
│ │ ├── product_handler.go
│ │ └── promo_handler.go
│ ├── models/ # Data models
│ │ ├── cart.go
│ │ ├── order.go
│ │ ├── product.go
│ │ └── promo.go
│ ├── repositories/ # Data storage implementations
│ │ └── memory/ # In-memory storage
│ │ └── store.go
│ └── routes/ # Route definitions
│ └── routes.go
├── go.mod
├── go.sum
└── README.md


## Running the Application

1. Start the server:
bash
From the backend directory
go run cmd/api/main.go


The server will start on port 8080.

## API Documentation

### Products API

#### Get All Products
Products
bash
GET /api/products
Response:
json
[
{
"id": "1",
"name": "Spy x Family Tshirt",
"price": 26,
"rating": 4.6,
"sold": 1015,
"location": "South Jakarta",
"image_url": "https://via.placeholder.com/250",
"description": "Highly rated for durability and style.",
"category": "Clothing"
}
]
Cart

### Cart API

#### Get Cart
bash
GET /api/cart?userId=user123
Response:
json
{
"id": "cart123",
"user_id": "user123",
"items": [
{
"product_id": "1",
"name": "Spy x Family Tshirt",
"price": 26,
"quantity": 2,
"image_url": "https://via.placeholder.com/250"
}
],
"total": 52,
"discount": 0
}

#### Add to Cart
bash
POST /api/cart/add
Content-Type: application/json
{
"userId": "user123",
"productId": "1",
"quantity": 2,
"price": 26,
"name": "Spy x Family Tshirt",
"imageUrl": "https://via.placeholder.com/250"
}


#### Update Cart Item
bash
PUT /api/cart/{productId}
Content-Type: application/json
{
"userId": "user123",
"quantity": 3
}


#### Remove from Cart
bash
DELETE /api/cart/{productId}?userId=user123


#### Checkout
POST /api/cart/checkout
Content-Type: application/json
{
"userId": "user123",
"discountCode": "PROMO10"
}


### Order API

#### Get User Orders
bash
GET /api/orders?userId=user123
Response:
json
[
{
"id": "order123",
"user_id": "user123",
"items": [
{
"product_id": "1",
"name": "Spy x Family Tshirt",
"price": 26,
"quantity": 2
}
],
"total_amount": 52,
"discount_code": "PROMO10",
"discount_amount": 5.2,
"created_at": "2024-03-15T14:30:00Z"
}
]


### Promo Code API

#### Generate Promo Code
bash
POST /api/admin/discount
Response:
json
{
"id": "promo123",
"code": "PROMO10",
"discount_rate": 10.0,
"used": false,
"generated_at": "2024-03-15T14:30:00Z"
}


#### Get Active Promo Code
bash
GET /api/admin/discount/active
Response:
json
{
"id": "promo123",
"code": "PROMO10",
"discount_rate": 10.0,
"used": false,
"generated_at": "2024-03-15T14:30:00Z"
}


## Error Responses
All endpoints may return the following error responses:

json
{
"error": "Error message description"
}


Common HTTP Status Codes:
- 200: Success
- 400: Bad Request
- 404: Not Found
- 500: Internal Server Error

## Development Tips

1. Use Postman or similar tool for API testing
2. Monitor server logs for errors
3. Run tests before committing changes
4. Keep test coverage high
5. Follow Go best practices and conventions

## Testing

Run all tests:
bash
go test ./...


Generate coverage report:
bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out

```

This documentation now includes:
1. Updated project structure
2. Detailed API endpoints with request/response examples
3. Error handling information
4. Development and testing guidelines