---
layout: "ibm"
page_title: "IBM : ibm_is_private_path_service_gateway_account_policies"
description: |-
  Get information about PrivatePathServiceGatewayAccountPolicyCollection
subcategory: "VPC infrastructure"
---

# ibm_is_private_path_service_gateway_account_policies

Provides a read-only data source for PrivatePathServiceGatewayAccountPolicyCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_subnet" "example" {
  name = "example-subnet"
  vpc = ibm_is_vpc.example.id
  zone = "us-south-2"
  ipv4_cidr_block = "10.240.0.0/24"
}
resource "ibm_is_lb" "example" {
  name = "example-lb"
  subnets = [ibm_is_subnet.example.id]
}
resource "ibm_is_private_path_service_gateway" "example" {
  default_access_policy = "review"
  name = "my-example-ppsg"
  load_balancer = ibm_is_lb.example.id
  zonal_affinity = true
  service_endpoints = ["example-fqdn"]
}
resource "ibm_is_private_path_service_gateway_account_policy" "example" {
  private_path_service_gateway = ibm_is_private_path_service_gateway.example.id
  access_policy = "review"
  account = "fee82deba12e4c0fb69c3b09d1f12345"
}
data "ibm_is_private_path_service_gateway_account_policies" "example" {
	private_path_service_gateway = ibm_is_private_path_service_gateway.example.id
	account = "fee82deba12e4c0fb69c3b09d1f12345"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `private_path_service_gateway` - (Required, String) The private path service gateway identifier.
- `account` - (Optional, String) - ID of the account to retrieve the policies for.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `account_policies` - (List) Collection of account policies.
	Nested scheme for **account_policies**:
	- `access_policy` - (String) The access policy for the account:- permit: access will be permitted- deny:  access will be denied- review: access will be manually reviewedThe enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
	- `account` - (List) The account for this access policy.
		Nested scheme for **account**:
		- `id` - (String)
		- `resource_type` - (String) The resource type.
	- `created_at` - (String) The date and time that the account policy was created.
	- `href` - (String) The URL for this account policy.
	- `id` - (String) The unique identifier for this account policy.
	- `resource_type` - (String) The resource type.
	- `updated_at` - (String) The date and time that the account policy was updated.
