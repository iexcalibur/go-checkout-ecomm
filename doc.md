# Backend Documentation

## Project Structure

This project follows the principles of **Clean Architecture** to ensure a scalable, testable, and maintainable codebase.

```
backend/
├── cmd/
│   └── api/
│       └── main.go                # Application entry point
├── internal/
│   ├── config/
│   │   └── server.go              # Server configuration
│   ├── core/
│   │   ├── domain/                # Business domain definitions
│   │   │   ├── entities.go        # Core business entities
│   │   │   ├── cart.go            # Cart domain interfaces
│   │   │   ├── order.go           # Order domain interfaces
│   │   │   ├── product.go         # Product domain interfaces
│   │   │   └── promo.go           # Promo domain interfaces
│   │   ├── middleware/
│   │   │   └── middleware.go      # HTTP middleware
│   │   ├── ports/
│   │   │   └── repositories.go    # Repository interfaces
│   │   └── services/              # Business logic implementation
│   │       ├── cart_service.go
│   │       ├── order_service.go
│   │       └── promo_service.go
│   ├── data/
│   │   └── sample_products.go     # Sample data initialization
│   ├── handlers/                  # HTTP request handlers
│   │   ├── cart_handler.go
│   │   ├── order_handler.go
│   │   ├── product_handler.go
│   │   └── promo_handler.go
│   ├── models/                    # Data models
│   │   ├── cart.go
│   │   ├── order.go
│   │   ├── product.go
│   │   └── promo.go
│   ├── routes/
│   │   └── routes.go             # Route definitions
│   └── storage/
│       └── memory_store.go       # In-memory data storage
├── README.md                     # Project overview
└── doc.md                        # This documentation file
└── api.md                        # This API documentation file
```

## Component Descriptions

### 1. Entry Point

- `cmd/api/main.go`
  - Application bootstrap
  - Initializes components
  - Sets up middleware
  - Starts the server

### 2. Core Layer

- `core/domain/`

  - Contains business rules and interfaces
  - Defines domain entities and their behavior
  - Independent of external concerns

- `core/services/`

  - Implements business logic
  - Uses repositories for data access
  - Enforces business rules

- `core/middleware/`

  - HTTP middleware for logging, CORS, etc.
  - Cross-cutting concerns

### 3. Data Layer

- `storage/memory_store.go`

  - In-memory data storage implementation
  - Implements repository interfaces
  - Manages data persistence

- `models/`

  - Data structures for storage
  - Used by storage and handlers

### 4. API Layer

- `handlers/`

  - HTTP request handlers
  - Converts HTTP requests to domain operations
  - Uses services to execute business logic

- `routes/`

  - API route definitions
  - Maps URLs to handlers

### 5. Configuration

- `config/`
  - Server configuration
  - Environment settings

## Component Interactions

1. **Request Flow**

```
HTTP Request → Router → Handler → Service → Repository → Storage
Response    ←        ←         ←         ←            ←
```

2. **Dependency Flow**

```
Handlers → Services → Repositories
                  ↓
               Storage
```

3. **Domain Layer**

- Defines core business rules
- Used by services
- Independent of other layers

## Key Features

1. **Cart Management**

- Add/remove items
- Calculate totals
- Apply discounts

2. **Order Processing**

- Create orders from carts
- Track order history
- Handle discounts

3. **Promo Code System**

- Manual promo code creation
- Automatic promo generation
- Discount application

4. **Product Catalog**

- Product listing
- Product details
- Sample data initialization

## Design Patterns

1. **Repository Pattern**

- Abstracts data storage
- Defined in `ports/repositories.go`
- Implemented in `storage/memory_store.go`

2. **Service Layer**

- Encapsulates business logic
- Coordinates between handlers and repositories
- Implements domain interfaces

3. **Dependency Injection**

- Components receive dependencies
- Facilitates testing
- Loose coupling

## Testing

- Unit tests for components
- Handler tests with HTTP mocks
- Service tests with repository mocks
- Storage tests for data operations

## Security

- CORS middleware
- Input validation
- Error handling

