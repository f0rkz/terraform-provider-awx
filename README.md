# Terraform Provider AWX

_Fork from [mrcrilly/terraform-provider-awx](https://github.com/mrcrilly/terraform-provider-awx) for develop additional functions_

Coming soon.

## Local Development

```sh
# Start a Local AWX deployed to a Kind
cd ./tools && go run mage.go -v reCreate && cd ..
```

```sh
# Build Provider
goreleaser build --snapshot --rm-dist

# Copy Provider
mkdir -p ~/.terraform.d/plugins/github.com/mrcrilly/awx/0.1/linux_amd64/terraform-provider-awx
find ./dist/terraform-provider-awx_linux_amd64/* -name 'terraform-provider-awx*' -print0 | xargs -0 -I {} mv {} ~/.terraform.d/plugins/github.com/mrcrilly/awx/0.1/linux_amd64/terraform-provider-awx

# start the Tests
go test ./test -count=1

# or as one command
goreleaser build --snapshot --rm-dist \
    && mkdir -p ~/.terraform.d/plugins/github.com/mrcrilly/awx/0.1/linux_amd64/ \
    && find ./dist/terraform-provider-awx_linux_amd64/* -name 'terraform-provider-awx*' -print0 | xargs -0 -I {} mv {} ~/.terraform.d/plugins/github.com/mrcrilly/awx/0.1/linux_amd64/terraform-provider-awx \
    && go test ./test -count=1


```

## Documentation

The files from `./docs` are generated by `cd ./tools && go run mage.go -v genDocumentation && cd ..`
