variable "APP_ENV" {
  description = "app env"
  type        = string
  sensitive   = false
}
variable "APP_PASSWORD" {
  description = "APP password"
  type        = string
  sensitive   = true
}
variable "FROM_EMAIL" {
  description = "Email from where the mail will be sent to"
  type        = string
  sensitive   = true
}
variable "EMAIL_RECEIVERS" {
  description = "List of emails that will receive the email"
  type        = string
  sensitive   = true
}
