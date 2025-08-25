# Kubernetes resource example plugin

`kubectl` plugin for retrieving resource example YAMLs.

All examples originate from kubernetes.io.

### Usage

#### Build and test using Go

`make build`

`./kubectl-example deployment`

Output:
```
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: nginx:1.14.2
          ports:
            - containerPort: 80
```

#### Or download precompiled binaries

Github releases [link](https://github.com/talos-labs/kubectl-example/releases).

#### Installing as a `kubectl` plugin

Simply move the compiled binary into a directory in `PATH`

`mv ./kubectl-example /usr/local/bin`

and then can be used with `kubectl` as

`kubectl example pod`


#### Build and test using Docker

`make build-docker`

`docker run --rm -i -t kubectl-example deploy`

#### Acknowledgements

This project is a fork of the work of [xuguohao123456/kubectl-example](https://github.com/xuguohao123456/kubectl-example) which itself is a fork of [seredot/kubectl-example](https://github.com/seredot/kubectl-example)


