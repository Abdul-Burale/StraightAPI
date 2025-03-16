# **StraightAPI**

**StraightAPI** is an eCommerce Transaction API designed for businesses to create orders and facilitate secure online payments. The API allows users to create orders, store order data in MongoDB, and generate payment links via Stripe for easy and secure transactions.

## **Features**
- **Order Creation**: Users can create orders with product details, which are stored in MongoDB.
- **Payment Integration**: Stripe is used to generate payment links, allowing users to securely complete transactions.
- **MongoDB Storage**: Order details are saved in MongoDB for easy access and management.
- **Authentication**: Google Firebase authentication for secure user access.
- **CORS Support**: Cross-Origin Resource Sharing (CORS) enabled for secure access across different domains.

## **Technologies Used**
- **Go (Golang)**: Backend framework for building the RESTful API.
- **MongoDB**: NoSQL database for storing order data.
- **Stripe**: Payment gateway for generating payment links and handling transactions.
- **Google Firebase**: User authentication and security.
- **Heroku**: Deployment platform for hosting the API.

## **Setup and Installation**

### 1. Clone the repository:
```bash
git clone https://github.com/Abdul-Burale/straightapi.git
```
### 2. Install dependencies:
Make sure Go is installed, and then run the following to install necessary packages:

```bash
go mod tidy
3. Set up environment variables:
Create a .env file in the root of the project and add the following variables:
```

```ini
MONGO_URI=your_mongodb_connection_string
STRIPE_SECRET_KEY=your_stripe_secret_key
FIREBASE_AUTH_KEY=your_firebase_auth_key
```

4. Run the project locally:
```bash
go run main.go
```
The API will be available at http://localhost:8080.


## **API Endpoints**

- **POST /orders**  
  Create a new order.  
  *Request Body:*

  ```json
  {
    "product_id": "string",
    "quantity": "integer",
    "customer_name": "string",
    "customer_email": "string"
  }
  
## **Still in Development**

:construction: **StraightAPI** is a work in progress! New features, improvements, and bug fixes are being actively worked on. We encourage you to contribute or open an issue if you encounter any problems. Your feedback and contributions are highly appreciated!

## **Contributing**

We welcome contributions to **StraightAPI**! If you'd like to contribute, follow these steps:

1. **Fork the repository** to your own GitHub account.
2. **Create a new branch** for your feature or bug fix:
   ```bash
   git checkout -b feature-name

