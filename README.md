# Go E-commerce Checkout System

A backend service for e-commerce checkout system built with Go, featuring cart management, order processing, and automatic promo code generation.

## Project Structure
This project follows Clean Architecture principles. For detailed documentation, see [doc.md](doc.md).

## Prerequisites

- Go 1.21 or higher
- Git

## Installation

1. **Clone the repository**
```bash
git clone https://github.com/iexcalibur/go-checkout-ecomm.git
cd go-checkout-ecomm
```

2. **Initialize Go modules**
```bash
go mod init github.com/iexcalibur/go-checkout-ecomm
go mod tidy
```

## Running the Application

1. **Start the server**
```bash
go run cmd/api/main.go
```
The server will start on `http://localhost:8080`

## API Endpoints

### Cart Operations
- Add to Cart: `POST /api/cart/add`
- Get Cart: `GET /api/cart?userId={userId}`
- Update Cart Item: `PUT /api/cart/{productId}`
- Remove from Cart: `DELETE /api/cart/{productId}?userId={userId}`
- Checkout: `POST /api/cart/checkout`

### Promo Code Operations
- Generate Promo: `POST /api/admin/discount`
- Get Active Promo: `GET /api/admin/discount/active`


## Features

- Cart Management
- Order Processing
- Automatic Promo Code Generation (every 3rd order)
- Discount Application
- In-memory Data Storage

## Testing

Run all tests:
```bash
go test ./...
```



