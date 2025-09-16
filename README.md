# AWS Account Handler

This project is a serverless, Lambda-based AWS account handler built using the AWS Cloud Development Kit (CDK) in Go. It allows you to define and manage your AWS infrastructure as code.

## Features

* **Infrastructure as Code**: Define your AWS resources using familiar programming languages (Go).
* **Serverless Architecture**: Leverages AWS Lambda for efficient, event-driven execution.
* **AWS CDK**: Utilizes the AWS Cloud Development Kit for robust infrastructure provisioning.

## Project Structure

The repository is organized as follows:

* `/cmd/`: Contains the entry points for the Go Lambda functions.
* `/pkg/`: Houses reusable Go packages.
* `/infrastructure/`: Defines the AWS CDK stacks and constructs for your infrastructure.
* `cdk.json`: Configuration file for the AWS CDK.
* `go.mod` and `go.sum`: Go module and checksum files for dependency management.

## Prerequisites

Before you can work with this project, ensure you have the following installed:

* **Go**: Version 1.19 or higher.
* **AWS CDK**: Version 2.x or higher.
* **AWS CLI**: Installed and configured with appropriate AWS credentials.

## Getting Started

### Installation

1.  **Clone the repository**:
    ```bash
    git clone [https://github.com/Sebjen0209/aws-account-handler.git](https://github.com/Sebjen0209/aws-account-handler.git)
    cd aws-account-handler
    ```

2.  **Install Go dependencies**:
    ```bash
    go mod tidy
    ```

### Useful Commands

* **Bootstrap AWS CDK**: Prepares your AWS account and bootstrap environment.
    ```bash
    cdk bootstrap
    ```

* **Synthesize CloudFormation templates**: Generates the CloudFormation templates from your CDK code.
    ```bash
    cdk synth
    ```

* **Deploy to AWS**: Deploys your infrastructure to your AWS account.
    ```bash
    cdk deploy
    ```

* **Build Lambda binaries**: Compiles the Go Lambda functions into executable binaries.
    ```bash
    # Example for a specific Lambda function (adjust as needed)
    GOOS=linux GOARCH=amd64 go build -o ./bin/my-lambda-function ./cmd/my-lambda-function/main.go
    ```
    *(Note: Specific build commands for each Lambda function might vary based on your project's needs. You might integrate this into your CDK deployment process).*

* **Run Unit Tests**: Executes all unit tests in the project.
    ```bash
    go test ./...
    ```

## Contributing

Contributions are welcome! Please refer to the contribution guidelines (if available) or open an issue/pull request on GitHub.

## License

This project is licensed under the [MIT License] (LICENSE).
