# Libreraium – Microservices-based E-commerce System

**Libreraium** is a distributed e-commerce system for buying digital and physical books, built with a **microservices architecture** and accessed via a **console-based client**. It focuses on modularity, fault tolerance, and message-oriented communication.

---

## 1. Requirements

### Functional Requirements

- User authentication (register, login)
- Book product catalog (CRUD)
- Shopping cart & checkout
- Order messaging via RabbitMQ (MOM)
- Console-based interface (no web UI)

### Technical Requirements

- Java-based API Gateway
- Node.js Auth service (with JWT + SQLite)
- Go-based Cart & Product services (with SQLite)
- Communication via HTTP
- Asynchronous messaging via RabbitMQ (MOM)
- Token-based authorization using JWT

---

## 2. Architecture & Analysis

### High-Level Architecture (C4 model)

- **API Gateway (Java):** Central point for client communication, token validation, and routing
- **Auth Service (Node.js):** Handles user login and registration, emits events to RabbitMQ
- **Cart Service (Go):** Manages cart items, checkout, and sends events to MOM
- **Product Service (Go):** Manages catalog of books
- **MOM Middleware (RabbitMQ):** Guarantees delivery of checkout and login events
- **SQLite Databases:** Lightweight storage for each service (decoupled persistence)

---

## 3. Design

### Microservices Breakdown

| Service         | Language | Responsibilities                            |
|------------------|----------|---------------------------------------------|
| API Gateway      | Java     | Routing, token validation, logging          |
| Auth Service     | Node.js  | Auth, token generation, MOM publishing      |
| Cart Service     | Go       | Cart logic, JWT check, checkout via MOM     |
| Product Service  | Go       | Catalog logic (add/list books)              |
| MOM (RabbitMQ)   | N/A      | Message bus for `checkout`, `login` events  |

###  Token Handling

- JWTs are created in the Auth Service
- API Gateway and Cart Service verify token using shared secret
- Token includes `sub` (username) and `role`

---

##  4. Implementation Details

###  Languages & Frameworks

- Java (`HttpServer`) – API Gateway
- Node.js (`express`, `jsonwebtoken`, `better-sqlite3`, `amqplib`) – Auth
- Go (`mux`, `sqlite`, `jwt/v5`, `streadway/amqp`) – Cart/Product
- RabbitMQ – via direct exchange

###  Testing Tools

- PowerShell / cURL for API requests
- Logs and terminal outputs for tracking requests and failures

###  Example Request (Login)

POST /login
{
  "username": "admin",
  "password": "1234"
}

Returns:

{ "token": "..." }

 5. Running the System
 Requirements

    Node.js (v18+)

    Go (v1.20+)

    Java (v11+)

    RabbitMQ server

    SQLite (bundled in code)

 Steps

# 1. Run Auth Service
cd auth-service
npm install
node index.js

# 2. Run Cart Service
cd cart-service
go mod tidy
go run main.go

# 3. Run Product Service
cd product-service
go mod tidy
go run main.go

# 4. Run API Gateway (Java)
cd api-gateway
javac -d out src/**/*.java
java -cp out gateway.ApiGateway

    RabbitMQ should be running and reachable by all services.

 6. Usage Example

    Login to receive a token

    Add books via Product Service

    Add books to cart using token (JWT in Authorization)

    Checkout triggers event to RabbitMQ

    MOM receives checkout and logs or forwards to consumers

 Authors

    Esteban Giraldo [@egiraldol]
    Mariana Gutierrez [@gutim1011]
    Sofia Zapata [@alkran93]


    2025-1
