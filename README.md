# GPG Terraform Provider

This provider lets you generate a GPG key pair.

Once you've written your provider, you'll want to [publish it on the Terraform Registry](https://www.terraform.io/docs/registry/providers/publishing.html) so that others can use it.

## Using the provider

```hcl
resource "gpg_private_key" "key" {
    name = "John Doe"
    email = "john@doe.com"
}
```

## Contributing

If you wish to work on the provider, you'll need:

- [Terraform](https://www.terraform.io/downloads.html) >= 0.15.x
- [Go](https://golang.org/doc/install) >= 1.16

To compile the provider, run `make`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

Add the local build to your local Terraform plugins so you can test it in your project context.

```sh
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/Olivr/gpg/0.0.1/darwin_amd64
ln -s $GOPATH/bin/terraform-provider-gpg ~/.terraform.d/plugins/registry.terraform.io/Olivr/gpg/0.0.1/darwin_amd64/terraform-provider-gpg
```

In order to run the full suite of Acceptance tests, run `make test`.

To generate or update documentation, run `make doc`.
