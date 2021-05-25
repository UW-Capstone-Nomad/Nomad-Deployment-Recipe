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

Backpack is currently tested against Nomad version 0.12.8







# Get Start Guide

## Download & Install


Backpack can be built from source by firstly cloning the repository git clone https://github.com/UW-Capstone-Nomad/Nomad-Deployment-Recipe.git. Once cloned, a binary can be built using the make build command which will be available at ./bin/backpack.
```shell
git clone https://github.com/UW-Capstone-Nomad/Nomad-Deployment-Recipe.git
cd Nomad-Deployment-Recipe/backpack
sudo apt-get install make
make build
cd /build
./backpack create nginx
./backpack pack ./nginx-0.1.0/
./backpack unpack values ./nginx-0.1.0.backpack -f ./values.yaml
./backpack plan ./nginx-0.1.0.backpack -v ./values.yaml

```

## How to use

1. Add the host_volume information to the client stanza in the Nomad configuration
   - create config file: sudo touch /etc/nomad.d/client.hcl
   - edit client.hcl
```
client {
  enabled = true
  host_volume "my-website-db" {
    path = "/opt/volumes/my-website-db"
    read_only = false
  }
}

```

2. Create a folder on one of the Nomad clients to host the registry files
```
sudo mkdir /opt/volumes/my-db
```

3. Start nomad agent with -config to load config file (start agent)
```
sudo nomad agent -config=/etc/nomad.d/client.hcl -dev -bind 0.0.0.0 -log-level INFO
```

4. Deploy the application

In this part, users will fellow the guide of [backpack](https://gitlab.com/koalalorenzo/backpack/-/blob/master/README.md). It mainly include create, plan and run.
Here, take minio as an example.

- Create your first pack, by using the boilerplate directory structure:

```shell
./backpack create minio
```

- Plan and validate (dry-run) the jobs of a package before running:
```shell
./backpack plan ./nginx-0.1.0.backpack
```

- Run your Nomad Jobs with my custom values:
```shell
./backpack run ./nginx-0.1.0.backpack
```








## Read More

* [How to build Backpack binaries from source code](docs/build.md)
* [How to install Backpack from source code](docs/build.md#installing)
* [How to use backpack with Helm](docs/usage.md)

# Copyright and License

Copyright © 2020 Lorenzo Setale https://setale.me
The full license is available in the file [LICENSE](LICENSE)
