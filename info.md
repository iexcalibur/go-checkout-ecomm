# E-Commerce Backend Documentation

## Project Structure
bash
backend/
├── cmd/
│ └── api/ # Application entry point
├── internal/
│ ├── config/ # Server configuration
│ ├── core/ # Core business logic
│ ├── data/ # Sample data
│ ├── handlers/ # HTTP handlers
│ ├── models/ # Data models
│ ├── repositories/ # Data storage interfaces
│ ├── routes/ # Route definitions
│ └── storage/ # Storage implementation


## Running the Application

1. Start the server:
bash
From the backend directory
go run cmd/api/main.go


The server will start on port 8080.

## Testing API Endpoints

### Products API

1. Get All Products
bash
curl http://localhost:8080/api/products

Example Response:
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
},
...
]


### Cart API

1. Add to Cart
bash
curl -X POST http://localhost:8080/api/cart/add \
-H "Content-Type: application/json" \
-d '{
"userId": "user123",
"itemId": "1",
"quantity": 2
}'

2. View Cart
bash
curl http://localhost:8080/api/cart?userId=user123

3. Update Cart Item Quantity
bash
curl -X PUT http://localhost:8080/api/cart/1/quantity \
-H "Content-Type: application/json" \
-d '{
"userId": "user123",
"quantity": 3
}'


4. Remove from Cart
bash
curl -X DELETE http://localhost:8080/api/cart/1?userId=user123


5. Checkout
bash
curl -X POST http://localhost:8080/api/cart/checkout \
-H "Content-Type: application/json" \
-d '{
"userId": "user123",
"discountCode": "DISC10"
}'


### Admin API

1. Get Active Discount

Discount
bash
curl http://localhost:8080/api/admin/discount/active


2. Generate Discount Code
bash
curl -X POST http://localhost:8080/api/admin/discount


3. Get Statistics
bash
curl http://localhost:8080/api/admin/stats

## Running Tests

1. Run all tests:
bash
go test ./...


2. Run specific test package:
bash
Test handlers
go test ./internal/handlers -v
Test storage
go test ./internal/storage -v
Test services
go test ./internal/core/services -v



3. Run with coverage:
bash
go test ./... -cover


4. Generate coverage report:
bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out


## Common Test Cases

### Products
- Verify all products are returned
- Check product details are correct
- Verify category filtering works

### Cart
- Add items to cart
- Update quantities
- Remove items
- Apply discount codes
- Complete checkout process

### Admin
- Generate discount codes
- View active discounts
- Check order statistics

## Troubleshooting

1. Server won't start:
   - Check port 8080 is available
   - Verify all dependencies are installed

2. API returns 404:
   - Verify the endpoint URL is correct
   - Check route registration in routes.go

3. Tests failing:
   - Run with -v flag for detailed output
   - Check test data setup
   - Verify mock expectations

## Development Tips

1. Use Postman or similar tool for API testing
2. Monitor server logs for errors
3. Run tests before committing changes
4. Keep test coverage high
5. Follow Go best practices and conventions


# 1. First, check if there's any active discount
curl http://localhost:8080/api/admin/discount/active

# 2. Generate a new discount code
curl -X POST http://localhost:8080/api/admin/discount

# 3. Verify the new code is now active
curl http://localhost:8080/api/admin/discount/active

# 4. Try to generate another code (should fail because one is active)
curl -X POST http://localhost:8080/api/admin/discount

# 5. Use the code in a checkout
curl -X POST http://localhost:8080/api/cart/checkout \
-H "Content-Type: application/json" \
-d '{
    "userId": "user123",
    "discountCode": "DISC123"
}'

# 6. Check that the code is now marked as used
curl http://localhost:8080/api/admin/discount/active