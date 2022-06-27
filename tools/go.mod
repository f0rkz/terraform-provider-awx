module github.com/mrcrilly/terraform-provider-awx/tools

go 1.14

require (
	github.com/golang/snappy v0.0.1 // indirect
	github.com/hashicorp/terraform v0.13.2
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.8.0
	github.com/hashicorp/terraform-plugin-test/v2 v2.1.2 // indirect
	github.com/magefile/mage v1.11.0
	github.com/mrcrilly/terraform-provider-awx v0.1.2
	github.com/nolte/plumbing v0.0.1
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
)

replace github.com/mrcrilly/goawx => github.com/f0rkz/goawx v0.6.0

replace github.com/mrcrilly/terraform-provider-awx => ../.
