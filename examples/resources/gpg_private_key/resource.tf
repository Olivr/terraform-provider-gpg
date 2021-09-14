resource "gpg_private_key" "secure_key" {
  name       = "John Doe"
  email      = "john@doe.com"
  passphrase = "this is not a secure passphrase"
  rsa_bits   = 3072
}
