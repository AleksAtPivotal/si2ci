#!/usr/bin/env bash

set -e -x

# For ssh tunnel test
/etc/init.d/ssh start

export PATH=/usr/local/ruby/bin:/usr/local/go/bin:$PATH
export GOPATH=$(pwd)/gopath

cd gopath/src/github.com/cloudfoundry/bosh-cli
bin/clean
bin/install-ginkgo
bin/test-integration
