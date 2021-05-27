# NomadManagementTool
The idea is to optimize the deployment experience on Nomad using a Helm like strategy.
write up document: https://docs.google.com/document/d/1x5ptaMzx3bB-xUZB6xUCGeOvNiah_Cl0OZuhcJt81a0/edit#

TODO:
Pending example scripts: Wordpress, Minio, Prometheus, Grafana, SockShop, Web server (Apache)

# Nomad's Backpack 🎒

[Backpack](https://backpack.qm64.tech) 🎒 is a packaging system for
[Hashicorp Nomad](https://www.nomadproject.io) that allows to:

* Helps you define and install complex jobs configuration
* Helps building reproducible jobs across multiple Nomad clusters
* Simplifies updates to new version of jobs
* Allows you to publish and share packages of applications

Please, keep in mind that this is designed for Nomad, and it might result as
very different than Helm, as Kubernetes is way more than a scheduler.
Read more [here](https://www.nomadproject.io/intro/vs/kubernetes.html) about
the differences between k8s and nomad

To learn more about the motivation behind the development of this project
check [the blog post on Qm64 website](https://qm64.tech/posts/202011-hashicorp-nomad-backpack/).

If you need some help or you want to stay updated with the latest news,
[join Qm64's chatroom on Matrix](https://matrix.to/#/#qm64:matrix.org?via=matrix.org)

Backpack is currently tested against Nomad version 1.1.0

This version of Backpack is baesd on the work of [Lorenzo Setale](https://blog.setale.me/) 
but get modified by a team of the University of Washington to make the Backpack closer to [Helm](https://helm.sh/).
There are new features such as:
* Exporting deployed job port
* Offering well-documented example applications
* Easier to maintain and integrate with other software long-term


## TL;DR: Install
You can manually download release from our version from:
[the release page here] https://github.com/UW-Capstone-Nomad/Nomad-Deployment-Recipe/tree/main/backpack/build

Or original version: [the release page here](https://gitlab.com/Qm64/backpack/-/releases).

Or compile the binaries:
```shell
go get -v gitlab.com/Qm64/backpack
cd $GOPATH/src/gitlab.com/Qm64/backpack/
make install
```

## TL;DR How to Use
Refer to get started guide under /test/minio directory
or under /backpack directory
Happy Backpacking! 🎒😀 

## Read More

* [How to build Backpack binaries from source code](docs/build.md)
* [How to install Backpack from source code](docs/build.md#installing)
* [How to use backpack with Helm](docs/usage.md)

# Copyright and License

Copyright © 2020 Lorenzo Setale https://setale.me
The full license is available in the file [LICENSE](LICENSE)