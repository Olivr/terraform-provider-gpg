---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gpg_private_key Resource - terraform-provider-gpg"
subcategory: ""
description: |-
  The resource gpg_private_key generates a GPG private/public key pair in ASCII-armored format.
---

# gpg_private_key (Resource)

The resource `gpg_private_key` generates a GPG private/public key pair in ASCII-armored format.

## Example Usage

```terraform
resource "gpg_private_key" "secure_key" {
  name       = "John Doe"
  email      = "john@doe.com"
  passphrase = "this is not a secure passphrase"
  rsa_bits   = 3072
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **email** (String) Email attached to the key.
- **name** (String) Name attached to the key.

### Optional

- **id** (String) The ID of this resource.
- **passphrase** (String, Sensitive) Passphrase protecting the private key.
- **rsa_bits** (Number) Number of bits to use when generating RSA key.

### Read-Only

- **fingerprint** (String) Public key fingerprint.
- **private_key** (String, Sensitive) Generated private key in ASCII-armored format.
- **public_key** (String) Generated public key in ASCII-armored format.


