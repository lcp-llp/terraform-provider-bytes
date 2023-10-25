# 1. Specify the version of the Bytes Provider to use
terraform {
  required_providers {
    bytes = {
      version = "~> 0.0.2"
      source  = "terraform.test.com/alpha/bytes"
    }
  }
}
# 2. Configure the Bytes Provider
provider "bytes" {
  username = "xxx"
  password = "xxx"
  identity_api_url = "xxx"
  commerce_api_url = "xxx"
  contract_id = "12345"
}

# 3. Query an existing order
data "bytes_order" "example" {
  order_id = "12345"
}

# 4. Create a new subscription
resource "bytes_subscription" "example" {
  friendly_name = "examplesub"
  po_number = "13102023-example"
  default_admin = "username@domain.uk.com"
}