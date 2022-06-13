# Garbage notifier

Send SMS notification the evening before garbage collection

## Dependencies

-   [Go](https://go.dev/doc/install)
-   [AWS Lambda](https://aws.amazon.com/lambda/)
-   [Twilio](https://www.twilio.com/sms)
-   [Terraform](https://www.terraform.io/)

### Prerequisites

Download and install:

-   [Go](https://go.dev/doc/install)
-   [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html) - optional (TerraForm)

### Running the app

1. Create `garbage.csv` file with garbage collection schedule
2. Change environment variables
3. Run command below to start the project

```bash
$ go run main.go
```

### Build

1. Build binary:
```bash
$ GOOS=linux GOARCH=amd64 go build -o garbage
```

2. Create ZIP file (binary + CSV):
```bash
$ zip garbage.zip garbage garbage.csv
```

### Deploy

Project uses [Terraform](https://www.terraform.io/) to deploy and provising AWS Lambda function, triggers ect.

Basic TF commands:
- `terraform plan` - compares current state and config file, and displays required provision steps
- `terraform apply` - triggers execution plan (spins up AWS Lambda, creates rules, ect.)

Lambda is currently being triggered every workday (MON - FRI) at 6pm UTC.

### Test

When Lambda function has been successfuly deployed to AWS, run this command:

- `aws lambda invoke --function-name garbageAutomation response.json`


## Authors

👤 **Miha Luksic**
