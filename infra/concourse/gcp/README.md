# Infrasturcture Setup

`terraform init`

`terraform plan`

`terraform apply`

SSH into the virtual machine

Download `docker-compose` binary

```sh
wget https://raw.githubusercontent.com/concourse/concourse-docker/master/docker-compose.yml
```  

and execute

```sh
docker-compose up -d
```

to bring up Concourse