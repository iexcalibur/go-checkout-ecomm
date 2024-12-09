# API Documentation

This document provides detailed information about the API endpoints related to the cart functionality. The examples include requests and responses for adding, updating, retrieving, and clearing cart items.

---

## 1. Add Product to Cart

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

## 2. Retrieve Cart Details

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

## 3. Update Cart Item Quantity

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

## 4. Clear a Specific Cart Item

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

## 5. Checkout Cart

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

## 6. Retrieve Order History

### Endpoint

`GET /api/orders`

### Query Parameters

- `userId`: The ID of the user whose order history is to be retrieved.

### Example Request

`GET http://localhost:8080/api/orders?userId=user123`

### Description

Retrieves the order history for the specified user.

---

## 7. Promo Code System Documentation

### Generate Promo Code

#### Endpoint

`POST /api/admin/discount`

#### Request Body

```json
{
    "code": "SUMMER25",
    "discount_rate": 25.0
}
```

#### Example Response

```json
{
    "id": "uuid",
    "code": "SUMMER25",
    "discount_rate": 25.0,
    "used": false,
    "generated_at": "2024-03-21T12:34:56Z"
}
```

### Description

Creates a new promo code with a unique code, custom discount rate, and tracks its status.

---

### Automatic Promo Code Generation

- Promo codes are automatically generated after every 3rd order.
- Format: PROMOwith random characters.
- Discount rate: Fixed at 10%.
- Marked as an automatic promo.

---

### Get Active Promo Code

#### Endpoint

`GET /api/admin/discount/active`

#### Example Response

```json
{
    "id": "uuid",
    "code": "PROMO0JJHA",
    "discount_rate": 10.0,
    "used": false,
    "generated_at": "2024-03-21T12:34:56Z",
    "is_automatic": true,
    "message": "Congratulations! You've earned this promo code after 5 orders!"
}
```

#### Description

- Returns the most recent unused promo code.
- Prioritizes automatic promos over manual ones.
- Returns a 404 if no active promo codes are found.

---

### Use Promo Code in Checkout

#### Endpoint

`POST /api/cart/checkout`

#### Request Body

```json
{
    "userId": "user123",
    "discountCode": "SUMMER25"
}
```

#### Example Response

```json
{
    "id": "order-uuid",
    "user_id": "user123",
    "items": [...],
    "total_amount": 156.0,
    "discount_code": "SUMMER25",
    "discount_amount": 39.0,
    "created_at": "2024-03-21T12:34:56Z"
}
```

#### Description

- Validates promo code existence and usage status.
- Applies the discount to the cart total.
- Marks the promo code as used after successful checkout.

---

### Notes

- **Manual Promo Codes**: Custom code and discount rate.
- **Automatic Promos**: Generated after every 3rd order with a 10% discount.
- Promo codes can only be used once.
- The system maintains the order count for automatic promo generation.

