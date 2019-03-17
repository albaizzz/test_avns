# microservice_distributor_deposit

## Running Application on Docker Container

This application you can run on docker container and configuration remote to vault.

### Prerequisite :
* Docker
* OS Linux Docker Image, please look on wiki standardize docker images  provide by devOps
* ask devOps where the host url path of vault engine

### Environment Variable available

* VAULT_ADDR (vault address  base url
* VAULT_TOKEN (vault tocken to access your secret config)
* VAULT_PATH ( path of your config on vault engine)


### build docker image
```bash
$ docker build -t {IMAGE_NAME} --build-arg SSH_PRIVATE_KEY="{SSH_PRIVATE_KEY}" --build-arg ORIGIN="{BRANCH_OR_TAG}" -f Dockerfile .
# example
$ docker build -t spring-connector:1.0 --build-arg SSH_PRIVATE_KEY="$(cat ~/.ssh/id_rsa)" --build-arg ORIGIN="development" -f Dockerfile .
```


### run docker container after build the image
```bash
$ docker container run -i --name {uniqe container name} --ulimit nofile=262144:262144 -p 8081:8081 -e VAULT_ADDR="{vault address}" -e VAULT_TOKEN="vault token" -e VAULT_PATH="{path of your secret}"  -t {image_name:image_version_or_tag} sh -c "/opt/microservice_distributor_deposit/run.sh && /opt/microservice_distributor_deposit/microservice_distributor_deposit serve"
# example run htt serve:
$ docker container  run -i --name spring-connector --ulimit nofile=262144:262144 -p 8081:8081 -e VAULT_ADDR="http://vault.kudoplay.net" -e VAULT_TOKEN="secret" -e VAULT_PATH="microservice_distributor_deposit"  -t spring-connector:1.0  sh -c "/opt/microservice_distributor_deposit/run.sh && /opt/microservice_distributor_deposit/microservice_distributor_deposit serve"
```

### check log the docker container:
```bash
$ docker logs -f {container_name|container_id}
# example to see the log of spring-connector:-->
$ docker logs -f spring-connector
```

### healthy check

```
# liveness
{base_url}/spring/v1/health/liveness

# readiness
{base_url}/spring/v1/health/readiness
```

