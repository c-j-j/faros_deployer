Deployment tool for [Faros](https://github.com/felipesere/faros).

# Contribution

1. `go get github.com/c-j-j/faros_deployer`

This will import the project into $GOPATH/src

This project uses [glide](https://github.com/Masterminds/glide) for dependency management, and uses the Vendor Experiment feature added in Go 1.5.

2. `brew install glide`

3. `export GO15VENDOREXPERIMENT=1`

4. To run all tests within faros_deployer, excluding vendor tests: `go test $(glide novendor)`
