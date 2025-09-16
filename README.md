# 🚀 AWS CDK Go Project (Lambda-Based)

This project is built with the [AWS Cloud Development Kit (CDK)](https://docs.aws.amazon.com/cdk/latest/guide/home.html) using **Go**.  
It demonstrates how to define and deploy AWS infrastructure as code, focusing on a **serverless Lambda-based architecture**.

The `cdk.json` file tells the CDK toolkit how to run your application.

---

## 📂 Project Structure

- **/cmd/** – Main entry points for your Lambda functions written in Go  
- **/pkg/** – Reusable Go packages (business logic, helpers, etc.)  
- **/infrastructure/** – CDK stack definitions (infrastructure as code)  
- **cdk.json** – CDK execution configuration  
- **go.mod / go.sum** – Go dependencies  

---

## 🔧 Prerequisites

Before you begin, make sure you have the following installed:

- [Go](https://go.dev/dl/) (≥1.19 recommended)  
- [AWS CDK](https://docs.aws.amazon.com/cdk/latest/guide/getting_started.html) (≥2.x)  
- [AWS CLI](https://docs.aws.amazon.com/cli/) configured with credentials  

---

## 🛠️ Useful Commands

| Command            | Description |
|--------------------|-------------|
| `cdk bootstrap`    | Prepare your AWS environment for CDK deployments |
| `cdk synth`        | Generate the CloudFormation template |
| `cdk diff`         | Compare deployed stack with local state |
| `cdk deploy`       | Deploy the stack to your AWS account/region |
| `cdk destroy`      | Remove the deployed stack |
| `go test ./...`    | Run unit tests across all packages |

---

## 📝 Development Workflow

1. **Write Infrastructure**  
   Define resources in the `/infrastructure` directory using CDK and Go.

2. **Implement Lambda Functions**  
   Add your Go code for Lambda functions under `/cmd`.

3. **Build Lambda Binaries**  
   ```bash
   GOOS=linux GOARCH=amd64 go build -o main cmd/your_lambda/main.go
