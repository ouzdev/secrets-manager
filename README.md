# AWS Secrets Manager Client Example

This repository hosts a Go application that demonstrates how to retrieve secrets stored in AWS Secrets Manager using the AWS SDK for Go v2. The example is intended to help developers understand how to interact with AWS Secrets Manager programmatically.

## Prerequisites

To run and modify this application, you need:

- **Go (version 1.15 or newer):** Ensure Go is installed and properly set up on your development environment. You can download it from [Go's official site](https://golang.org/dl/).
- **AWS Account:** An AWS account with access to AWS Secrets Manager.
- **AWS Credentials:** Ensure you have valid AWS credentials configured with permissions to access Secrets Manager. It's recommended to configure these credentials securely using IAM roles or environment variables for production setups.

## Installation

Follow these steps to get the project running on your local machine:

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/ouzdev/secrets-manager.git
    cd your-repository
    ```

2. **Install Dependencies:**

    The project uses the AWS SDK which should be automatically handled by Go modules. If you have dependencies to manage, ensure to document them here.

## Configuration

Update the `accessKeyID`, `secretAccessKey`, and `region` in the `main.go` file to match your AWS credentials and desired AWS region. These are crucial for the AWS SDK to authenticate and interact with AWS services.

## Usage

To run the application, execute the following command from the root of the project directory:

```bash
go run main.go
