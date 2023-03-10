
## tf-cdk-go project


This project is a demo project meant to test out CDKTF implemented in Go.

### Compile:
```bash
go build
```
### Synthesize Terraform resources to cdktf.out/:
```bash
cdktf synth [stack]
```
### Diff Terraform plan for the given stack:
```bash
cdktf diff [stack]
```
### Deploy given stack:
``` bash
cdktf deploy [stack]
```
### Destroy given stack:
```bash
cdktf destroy [stack]
```

### Providers
```bash
cdktf provider add "aws@~>3.0" null kreuzwerker/docker
```

### Local Setup
I ran into issues with Docker Compose running on arm hardware so we just run the containers ourselves.
``` bash
docker run -d -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 -h mysql --network tf-cdk_default mysql:5.7
```
``` bash
docker run -d -p 6379:6379 -h redis --network tf-cdk_default redis/redis-stack:latest
```
``` bash
docker run -p 8080:8080 -h app --network tf-cdk_default app:latest
```