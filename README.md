<div align="center">
<h1>Candado</h1>

[Candado](https://github.com/goodjobtech/candado), provides a global lock service for distributed systems.



</div>

## Use Cases

### Accessing a shared resource

Imagine a scenario where you have a shared resource that you want to access. You want to ensure that only process can access the resource at a time. If the way you are accessing the does not provides a **strong read-after-write consistency**, then there is a chance that the resource will be modified by another process while you are reading it, so called **dirty reads**. To prevent this, you can use a **lock** to ensure that only one process can access the resource at a time.


```sh
1. process 1: read the resource
2. process 2: read the resource
3. process 1: modify the resource 
4. process 2: modify the resource
5. process 1: write the modified resource
5. process 2: write the modified resource
```

## Architecture

The current implementation relies on Redis since it prevents **dirty reads** with sub-millisecond latency as it is an in-memory database.

## Deployment

### Knative

1. Download candado and apply the manifest to your cluster.

```sh
git clone git@github.com:GoodJobTech/candado.git
cd candado
k apply -f deploy/knative/yaml
```

2. Wait for the deployment to be ready.

```sh
$ kubectl get pods

candado-00006-deployment-596558b4cf-wqsm8        2/2     Running   0          7s
candado-redis-b9c5878dc-s5xmq                    1/1     Running   0          34m
```



### Kubernetes

```sh
```

## API


### Acquire the Lock 


#### Request

```
GET /acquire/{lock-id}
```

#### Response

```json
{
    "error": "",
    "state": 1,
    "success": true
}
```

### Release the Lock 


#### Request

```
GET /release/{lock-id}
```

#### Response

```json
{
    "error": "",
    "state": 0,
    "success": true
}
```

### Heartbeat 


#### Request

```
GET /heartbeat/{lock-id}
```

#### Response

```json
{
    "error": "",
    "state": 0,
    "success": true
}
```


## Contributing

All kinds of pull request and feature requests are welcomed!

## License

Candado's source code is licensed under [MIT License](https://choosealicense.com/licenses/mit/).
