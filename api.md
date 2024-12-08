# API Documentation

This document provides detailed information about the API endpoints related to product and cart functionality. The examples include requests and responses for fetching products, adding items to the cart, retrieving cart details, updating quantities, and other operations.

---

## 1. Retrieve List of Products

### Endpoint

`GET /api/product`

### Description

Retrieves a complete list of available products, including details such as name, price, description, and availability status.

### Example Request

`GET http://localhost:8080/api/product`

### Example Response

```json
[
    {
        "product_id": "1",
        "name": "Spy x Family Tshirt",
        "price": 26,
        "description": "High-quality T-shirt inspired by Spy x Family",
        "availability": true
    },
    {
        "product_id": "2",
        "name": "One Piece Hoodie",
        "price": 40,
        "description": "Comfortable hoodie with One Piece designs",
        "availability": true
    }
]
```

---

## 2. Add Product to Cart

### Endpoint

`POST /api/cart/add`

### Request Body

```json
{
    "userId": "user123",
    "productId": "1",
    "quantity": 2,
    "price": 26,
    "name": "Spy x Family Tshirt",
    "imageUrl": "https://fansarmy.in/cdn/shop/products/anyafront_1800x1800.jpg?v=1659786556"
}
```

### Description

Adds a product to the user's cart with the specified quantity and product details.

---

## 3. Retrieve Cart Details

### Endpoint

`GET /api/cart`

### Query Parameters

- `userId`: The ID of the user whose cart details are to be retrieved.

### Example Request

`GET http://localhost:8080/api/cart?userId=user123`

### Example Response

```json
{
   "id": "f3e08b3d-af44-4c3d-a25d-3f433475361b",
   "user_id": "user123",
   "items": [
       {
           "product_id": "1",
           "name": "Spy x Family Tshirt",
           "price": 26,
           "quantity": 20,
           "image_url": "https://fansarmy.in/cdn/shop/products/anyafront_1800x1800.jpg?v=1659786556"
       }
   ],
   "total": 520,
   "discount": 0
}
```

### Description

Retrieves the current state of the user's cart, including all items, total price, and any applicable discounts.

---

## 4. Update Cart Item Quantity

### Endpoint

`PUT /api/cart/{productId}`

### Example Request

`PUT http://localhost:8080/api/cart/1`

### Request Body

```json
{
    "userId": "user123",
    "quantity": 5
}
```

### Example Response

```json
{
    "id": "f3e08b3d-af44-4c3d-a25d-3f433475361b",
    "user_id": "user123",
    "items": [
        {
            "product_id": "1",
            "name": "Spy x Family Tshirt",
            "price": 26,
            "quantity": 50,
            "image_url": "https://fansarmy.in/cdn/shop/products/anyafront_1800x1800.jpg?v=1659786556"
        }
    ],
    "total": 1300,
    "discount": 0
}
```

### Description

Updates the quantity of a specific product in the user's cart and recalculates the total.

---

## 5. Clear a Specific Cart Item

### Endpoint

`DELETE /api/cart/{productId}`

### Query Parameters

- `userId`: The ID of the user whose cart item is to be cleared.

### Example Request

`DELETE http://localhost:8080/api/cart/1?userId=user123`

### Example Response

```json
{
    "id": "f3e08b3d-af44-4c3d-a25d-3f433475361b",
    "user_id": "user123",
    "items": null,
    "total": 0,
    "discount": 0
}
```

### Description

Removes the specified product from the user's cart and updates the cart details accordingly.

---

## 6. Checkout Cart

### Endpoint

`POST /api/cart/checkout`

### Request Body

```json
{
    "userId": "user123"
}
```

### Example Response

```json
{
    "id": "",
    "user_id": "user123",
    "items": [
        {
            "product_id": "1",
            "name": "Spy x Family Tshirt",
            "price": 26,
            "quantity": 6
        }
    ],
    "total_amount": 156,
    "discount_code": "",
    "discount_amount": 0,
    "created_at": "0001-01-01T00:00:00Z"
}
```

### Description

Finalizes the user's cart and prepares it for processing. Returns the details of the finalized cart.

---

## 7. Retrieve Order History

### Endpoint

`GET /api/orders`

### Query Parameters

- `userId`: The ID of the user whose order history is to be retrieved.

### Example Request

`GET http://localhost:8080/api/orders?userId=user123`

### Description

Retrieves the order history for the specified user.

