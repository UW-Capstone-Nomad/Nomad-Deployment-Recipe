import os
import os.path

file_exists_db = os.path.exists("/opt/volumes/my-website-db")
file_exists_hcl = os.path.isfile("/etc/nomad.d/client.hcl")

if file_exists_db == False:
  os.system("sudo mkdir /opt/volumes/my-website-db")

if file_exists_hcl == False:
  os.system("sudo touch /etc/nomad.d/client.hcl")

client_config = 'client { \n\
  enabled = true \n\
  host_volume "my-website-db" { \n\
    path = "/opt/volumes/my-website-db" \n\
    read_only = false \n\
  } \n\
} \
'

with open('/etc/nomad.d/client.hcl', 'w') as f:
  f.write(client_config)
  f.close()




print ("create success")
# return 'Success'