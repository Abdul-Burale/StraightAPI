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

### 2. Install dependencies:
Make sure Go is installed, and then run the following to install necessary packages:

```bash
go mod tidy
3. Set up environment variables:
Create a .env file in the root of the project and add the following variables:

ini
Copy
MONGO_URI=your_mongodb_connection_string
STRIPE_SECRET_KEY=your_stripe_secret_key
FIREBASE_AUTH_KEY=your_firebase_auth_key
4. Run the project locally:
bash
Copy
go run main.go
The API will be available at http://localhost:8080.

API Endpoints
POST /orders
Create a new order.
Request Body:

json
Copy
{
  "product_id": "string",
  "quantity": "integer",
  "customer_name": "string",
  "customer_email": "string"
}
POST /payment
Generate a payment link using Stripe.
Request Body:

json
Copy
{
  "order_id": "string",
  "amount": "float"
}
Deployment
StraightAPI is hosted on Heroku for easy deployment. You can deploy it to your own Heroku account by following these steps:

Create a Heroku app:

bash
Copy
heroku create
Add MongoDB and Stripe configuration to your Heroku app:

bash
Copy
heroku config:set MONGO_URI=your_mongodb_connection_string
heroku config:set STRIPE_SECRET_KEY=your_stripe_secret_key
Push the project to Heroku:

bash
Copy
git push heroku main
Visit the Heroku app URL to access the API.

Still in Development
StraightAPI is a work in progress. New features, improvements, and bug fixes are being actively worked on. Please feel free to contribute or open an issue if you encounter any problems.

Contributing
If you'd like to contribute to the development of StraightAPI, feel free to fork the repository, make changes, and submit a pull request. Please ensure your code adheres to the project's coding style and includes tests where applicable.

License
This project is licensed under the MIT License.

pgsql
Copy

Now you can simply copy this entire block and paste it into your README.md file on GitHub. It sho
