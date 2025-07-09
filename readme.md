POST /api/auth/sinup
{
"email": "user@example.com",
"password": "securepassword123",
"name": "John Doe"
}

POST /api/auth/singin
use jwt for middleware
{
"email": "user@example.com",
"password": "securepassword123"
}

Get All Transactions
GET /api/transactions
[{
"id": 1,
"amount": 150.50,
"type": "expense",
"category": "groceries",
"description": "Weekly shopping",
"created_at": "2023-07-20T10:00:00Z"
},
{
"id": 2,
"amount": 2000.00,
"type": "income",
"category": "salary",
"description": "Monthly salary",
"created_at": "2023-07-01T09:00:00Z"
}
Create Transaction
POST /api/transactions
{
"amount": 75.25,
"type": "expense",
"category": "dining",
"description": "Dinner with friends"
}

Get Single Transaction
GET /api/transactions/:id
{
"id": 1,
"amount": 150.50,
"type": "expense",
"category": "groceries",
"description": "Weekly shopping",
"created_at": "2023-07-20T10:00:00Z"
}
