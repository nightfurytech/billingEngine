# billingEngine
A simple billing engine application written in Go using websockets.

## Overview
This application provides a basic billing system where users can track their loan payments. It exposes an HTTP API to retrieve the total amount paid by a user.

## Input
The application uses a predefined set of loans as input:

```go
loans = []Loan{
    {UserID: 1, LoanStartDate: "2024-01-01", LoanEndDate: "2024-12-31", AmountTaken: 5000, AmountPaid: 2000},
    {UserID: 2, LoanStartDate: "2024-02-01", LoanEndDate: "2024-11-30", AmountTaken: 7000, AmountPaid: 3500},
    {UserID: 1, LoanStartDate: "2023-05-01", LoanEndDate: "2024-04-30", AmountTaken: 3000, AmountPaid: 3000},
}
```

## API Endpoint
### Get User Outstanding Amount
**Endpoint:** `/getOutstandingAmount/:userId`

**Method:** `GET`

**URL Parameter:**
- `userId` (integer) - The ID of the user whose outstanding amount paid needs to be retrieved.

**Response:**
```json
{
   "outStandingAmount": 4620000
}
```

### Get if User is Delinquent
**Endpoint:** `/delinquent/:userId`

**Method:** `GET`

**URL Parameter:**
- `userId` (integer) - The ID of the user to find if he/she is delinquent.

**Response:**
```json
{
   "IsDelinquent": false
}
```

**Response:**
```json
{
   "IsDelinquent": true
}
```

### Make payment for loan
**Endpoint:** `/payment/:userId?amount=110000`

**Method:** `POST`


**URL Parameter:**
- `userId` (integer) - The ID of the user to add payment.


**Query Parameter:**
- `amount` (integer) - Amount of the user wants to pay (should be more than or equal to weekly payment amount).

**Response:**
```json
{
   "payment": "success"
}
```

**Response:**
```json
{
   "error": "amount should be equal to or more than 110000.00"
}
```

## Running the Application
1. Clone the repository.
2. Install Go (if not already installed).
3. Run the application using:
   ```sh
   go run main.go
   ```
4. The server will start at `http://localhost:8080`.

## Logging
The application logs when the server is running:
```sh
Server is serving at port 8080
```

```mermaid
graph TB
%% === STYLES ===
classDef core fill:#1E90FF,stroke:#000,color:#000,stroke-width:2px,rx:10px,ry:10px;
classDef data fill:#9ACD32,stroke:#000,color:#000,stroke-width:2px,rx:10px,ry:10px;
classDef external fill:#FFD700,stroke:#000,color:#000,stroke-width:2px,rx:10px,ry:10px;

%% === USERS ===
Client(("Client<br/>HTTP Requests"))

%% === WEB API SERVER ===
subgraph "Web API Server"
  API["Billing Engine<br/>Go + Gin Framework"]:::core
  DataStore["In-Memory Data Store<br/>sync.Map"]:::data
end

Client -->|"sends HTTP requests"| API

%% === API HANDLERS ===
subgraph "API Handlers"
  GetOutstanding["Get Outstanding Amount<br/>GET /getOutstandingAmount/:userId"]:::core
  IsDelinquent["Check Delinquency<br/>GET /delinquent/:userId"]:::core
  MakePayment["Process Payment<br/>POST /payment/:userId"]:::core
end

API -->|"routes requests"| GetOutstanding
API -->|"routes requests"| IsDelinquent
API -->|"routes requests"| MakePayment

%% === DATA FLOW ===
GetOutstanding -->|"retrieves loan data"| DataStore
IsDelinquent -->|"retrieves loan data"| DataStore
MakePayment -->|"retrieves loan data"| DataStore
MakePayment -->|"updates AmountPaid"| DataStore

%% === DATA MODELS ===
subgraph "Data Models"
  LoanModel["Loan<br/>UserID, LoanStartDate, LoanEndDate, AmountTaken, AmountPaid"]:::data
end

DataStore -->|"stores loan data"| LoanModel

%% === CONSTANTS ===
subgraph "Constants"
  InterestRate["Interest Rate<br/>10%"]:::core
  HoursInWeek["Hours in a Week<br/>168 hours"]:::core
end

LoanModel -->|"uses constants for calculations"| InterestRate
LoanModel -->|"uses constants for calculations"| HoursInWeek

%% === EXTERNAL DEPENDENCIES ===
subgraph "External Dependencies"
  WebSocket["WebSocket Communication<br/>Potential Extension"]:::external
  PersistentStorage["Persistent Storage<br/>Potential Extension"]:::external
end

API -->|"potential integration"| WebSocket
API -->|"potential integration"| PersistentStorage

%% === REQUEST HANDLING ===
Client -->|"HTTP request"| API
API -->|"matches URL to handler"| GetOutstanding
API -->|"matches URL to handler"| IsDelinquent
API -->|"matches URL to handler"| MakePayment

GetOutstanding -->|"calculates outstanding amount"| DataStore
IsDelinquent -->|"calculates delinquency status"| DataStore
MakePayment -->|"validates and updates payment"| DataStore
```

