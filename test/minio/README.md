# Get started guide for deploying minIO using backpack
## MinIO deployment with backpack using vagrant 

step 1 under the file directory with the vagrant file `vagrant up`

step 2 `vagrant ssh` in two different terminals

step 3 in one of the terminal: `nomad agent -config=./client.hcl -dev -bind 0.0.0.0 -log-level INFO`
Or if you want to use your client configuatin file, add the following scripts to your nomad client file first: 
client {
  enabled = true
  host_volume "minio-data" {
    path = "/opt/volumes/minio-data"
    read_only = false
  }
}

step 4 `nomad run minio.nomad` (replace with backpack operations later)

step 5 check the running minio instance on your browser "localhost:9000", or:
"http://localhost:4646/ui" to use nomad ui, and then find the minIO job and its port to see your deployment


## Steps to use the tool: 
under test/ directory:
backpack pack ../minio
backpack plan ../minio-0.1.0.backpack
backpack run ../minio-0.1.0.backpack

reference:
https://github.com/angrycub/nomad_example_jobs/tree/main/applications