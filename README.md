# E-Commerce API Documentation

## Products API
### Create Product
POST /api/products
Request:

# Test Files Documentation
## Overview
Our test files ensure the core functionality of the e-commerce system works correctly. Here's a breakdown of each test file and its purpose:

### 1. Cart Handler Tests (`cart_handler_test.go`)
Tests the HTTP endpoints for cart operations:
- `TestCartHandler_AddToCart`: Tests adding items to cart
  - Verifies successful item addition
  - Checks error handling for invalid items
- `TestCartHandler_Checkout`: Tests checkout process
  - Tests checkout without discount
  - Tests checkout with invalid user

Run with: 

bash
go test ./internal/handlers -v -run "TestCartHandler"


### 2. Discount Service Tests (`discount_service_test.go`)
Tests the discount code generation and validation logic:
- `TestDiscountService_GenerateDiscountCode`: Tests discount code generation
  - Verifies new code generation when none exists
  - Checks prevention of multiple active codes
- `TestDiscountService_ValidateDiscountCode`: Tests discount code validation
  - Verifies valid discount code acceptance
  - Checks rejection of used codes

Run with:

bash
go test ./internal/core/services -v -run "TestDiscountService"


### 3. Memory Store Tests (`memory_store_test.go`)
Tests the in-memory storage operations:
- `TestMemoryStore_AddToCart`: Tests cart storage operations
  - Tests adding new items
  - Tests updating existing items
- `TestMemoryStore_RemoveFromCart`: Tests item removal
  - Tests successful item removal
  - Tests error handling for non-existent items

Run with:

bash
go test ./internal/storage -v -run "TestMemoryStore"

## How to Verify Tests are Working

1. **Run All Tests**
bash
go test ./... -v

2. **Check Test Coverage**
bash
go test ./... -cover

3. **Generate Coverage Report**
bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out


## Test Success Indicators
1. All tests should pass with "PASS" status
2. No error messages in output
3. Coverage report should show good coverage of critical paths

## Example Test Output

bash
=== RUN TestCartHandler_AddToCart
--- PASS: TestCartHandler_AddToCart (0.00s)
=== RUN TestCartHandler_Checkout
--- PASS: TestCartHandler_Checkout (0.00s)
=== RUN TestDiscountService_GenerateDiscountCode
--- PASS: TestDiscountService_GenerateDiscountCode (0.00s)
PASS
coverage: 85.7% of statements


## Common Issues and Solutions
1. **Failed Tests**: Check error messages for specific failures
2. **Low Coverage**: Add more test cases for untested scenarios
3. **Mock Errors**: Ensure all interface methods are implemented in mocks

## Best Practices
1. Run tests before committing changes
2. Add new tests when adding features
3. Keep tests focused and independent
4. Use meaningful test names and descriptions