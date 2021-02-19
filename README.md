Terraform `null` Provider
=========================

See the documentation for this provider on the [Terraform Registry](https://registry.terraform.io/providers/hashicorp/null/latest/docs) for usage information.

Maintainers
-----------

This provider plugin is maintained by the Terraform team at [HashiCorp](https://www.hashicorp.com/).

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.16 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-null`

```sh
$ git clone git@github.com:terraform-providers/terraform-provider-null $GOPATH/src/github.com/terraform-providers/terraform-provider-null
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-null
$ make build
```


Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.16+ is *required*).

To generate the documentation, run `go generate` in the root of the codebase.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-null
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
