# cURL commands

##### POST JSON login request

```sh
$ curl -i -XPOST \
> http://localhost:8080/loginJSON \
> -H 'content-type: application/json' \
> -d '{"user": "anton", "password": "123"}'
```

##### POST Form login request

```sh
$ curl -i -XPOST \
http://localhost:8080/loginForm \
-H 'content-type: application/x-www-form-urlencoded' \
-d "user=anton&password=123"
```