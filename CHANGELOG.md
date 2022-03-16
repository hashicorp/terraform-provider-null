## 3.1.1 (March 16, 2022)

NOTES:

* Updated [terraform-plugin-docs](https://github.com/hashicorp/terraform-plugin-docs) to `v0.7.0`:
  this improves generated documentation, with attributes now correctly formatted as `code`
  and provided with anchors.
* Functionally identical to the previous 3.1.0 release.

## 3.1.0 (February 19, 2021)

Binary releases of this provider now include the darwin-arm64 platform. This version contains no further changes.

## 3.0.0 (October 08, 2020)

Binary releases of this provider now include the linux-arm64 platform.

BREAKING CHANGES:

* Upgrade to version 2 of the Terraform Plugin SDK, which drops support for Terraform 0.11. This provider will continue to work as expected for users of Terraform 0.11, which will not download the new version. ([#47](https://github.com/terraform-providers/terraform-provider-null/issues/47))

## 2.1.2 (April 30, 2019)

* This releases includes another Terraform SDK upgrade intended to align with that being used for other providers as we prepare for the Core v0.12.0 release. It should have no significant changes in behavior for this provider.

## 2.1.1 (April 11, 2019)

* This releases includes only a Terraform SDK upgrade intended to align with that being used for other providers as we prepare for the Core v0.12.0 release. It should have no significant changes in behavior for this provider.

## 2.1.0 (February 27, 2019)

IMPROVEMENTS:

* The previous release contains an SDK incompatible with TF 0.12. Fortunately 0.12 was not released yet so upgrading the vendored sdk makes this release compatible with 0.12.

## 2.0.0 (January 18, 2019)

IMPROVEMENTS:

* The provider is now compatible with Terraform v0.12, while retaining compatibility with prior versions.

## 1.0.0 (September 26, 2017)

* No changes from 0.1.0; just adjusting to [the new version numbering scheme](https://www.hashicorp.com/blog/hashicorp-terraform-provider-versioning/).

## 0.1.0 (June 21, 2017)

NOTES:

* Same functionality as that of Terraform 0.9.8. Repacked as part of [Provider Splitout](https://www.hashicorp.com/blog/upcoming-provider-changes-in-terraform-0-10/)
