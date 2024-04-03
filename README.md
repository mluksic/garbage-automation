# Garbage pickup notification system

Send garbage pickup alerts to users via Email/SMS

## Dependencies

-   [Go v1.18+](https://go.dev/doc/install)
-   [AWS Lambda](https://aws.amazon.com/lambda/)
-   [Terraform](https://www.terraform.io/)

## Prerequisites

Download and install:

-   [Go](https://go.dev/doc/install)
-   [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html) - optional (for testing Lambda function) 

## Usage

1. Create `garbage.csv` file with garbage pickup schedule
2. Change env variables to fit your
3. Run the command below to start the project

```bash
$ go run main.go
```

## Build

1. Build binary:
```bash
$ GOOS=linux GOARCH=amd64 go build -o garbage
```

2. Create ZIP file (binary + CSV) for AWS Lambda
```bash
$ zip garbage.zip garbage garbage.csv
```

## Deploy

Project uses [Terraform](https://www.terraform.io/) to deploy and provising AWS Lambda function, triggers ect.

Create `secret.tfvars` and fill it with your variables
```
$ cd tf
$ cp example.tfvars secret.tfvars
```

Basic TF commands:
- `terraform plan --var-file="secret.tfvars"` - compares current state and config file, and displays required provision steps
- `terraform apply --var-file="secret.tfvars"` - triggers execution plan (spins up AWS Lambda, creates rules, ect.)

Lambda is triggered every workday (MON - FRI) at 6pm UTC.

## Test

When Lambda function has been successfully deployed to AWS, run this command:

- `aws lambda invoke --function-name garbageAutomation response.json`

## Authors

👤 **Miha Luksic**
