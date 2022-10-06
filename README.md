# strategy-rest-seeder 

> WIP

Executes a series of instructions against an endpoint using a specified authentication on a given strategy.


## Problem space

Often we are faced with external products/services either self-hosted or SaaS where they require some level of configuration this often happens at CI (deploy) time, or on continous basis.

A complete definition contains an [`auth`](#auth) map and a [`seed`](#strategy) map, you will 

## Auth

Is a map of authentication objects used by each action respectively.

```yaml
ouath1:
  type: OAuthClientCredentials
  username: randClientIdOrUsernameForBasicAuth
  password: randClientSecret
  oauth:
    serverUrl: http://localhost:8080/token
    scopes:
      - https://www.some-api-provider.com/scopes-example1
    endpointParams:
      foo:
        - bar
        - baz
  httpHeaders:
    X-Foo: bar
basic1:
  type: BasicAuth
  username: randClientIdOrUsernameForBasicAuth
  password: randClientSecret
  httpHeaders:
    X-Foo: bar
```

## Strategy

Strategy is a setting against which to perform one or more rest calls to ensure an idempotent update.

The top level `seed` can contain multiple blo

```yaml
# the name of the Action is the property itself 
# should conform to YAML language specifications
find-put-post-not-found-id:
  endpoint: https://postman-echo.com
  strategy: FIND/PUT/POST
  getEndpointSuffix: /get?json=emtpy&notfound=bar
  postEndpointSuffix: /post
  findByJsonPathExpr: "$.array[?(@.name=='fubar')].id"
  authMapRef: ouath1
  payloadTemplate: |
    { "value": "$foo" }
  variables:
    foo: bar
  # RunTime Vars are captured from a PUT or POST and can be used further down the strategy tree
  runtimeVars:
    someId: "$.array[?(@.name=='fubar')].id"
```

      runtimeVariables:
        trustedcertificategroups_id: "$.id"
        highavailabilityprofiles_id: "$.id"

Each Strategy option is defined below, and can be extended.

### `GET/POST`

FindPostStrategyFunc strategy calls a GET endpoint and if item ***FOUND it does NOT do a POST*** this strategy should be used sparingly and only in cases where the service REST implementation does not support an update of existing item.

### `FIND/POST`

FindPostStrategyFunc strategy calls a GET endpoint and if item ***FOUND it does NOT do a POST*** this strategy should be used sparingly and only in cases where the service REST implementation does not support an update of existing item.

### `PUT/POST`

PutPostStrategyFunc is useful when the resource is created a user specified Id
the PUT endpoint DOES NOT support a creation of the resource. PUT should throw a 4XX
for the POST fallback to take effect

### `GET/PUT/POST`

GetPutPostStrategyFunc known ID and only know a name or other indicator the pathExpression must not evaluate to an empty string in order to for the PUT to be called else POST will be called as item was not present.


### `FIND/PUT/POST`

FindPutPostStrategyFunc is useful when the resource Id is unknown i.e. handled by the system.
providing a pathExpression will evaluate the response.
the pathExpression must not evaluate to an empty string in order to for the PUT to be called
else POST will be called as item was not present

### `FIND/PATCH/POST`
 
FindPutPatchStrategyFunc is the same as FindPutPostStrategyFunc but uses PATCH instead of PUT

### `FIND/DELETE`

### `FIND/DELETE/POST`

FindDeletePostStrategyFunc is useful for when you cannot update a resource but it can be safely destroyed an recreated

### `PUT`


## Rest Action

```yaml
endpoint: https://postman-echo.com
strategy: FIND/PUT/POST
getEndpointSuffix: /get?json=provided&valid=true
postEndpointSuffix: /post
putEndpointSuffix: /put
findByJsonPathExpr: "$.array[?(@.name=='fubar')].id"
authMapRef: ouath1
payloadTemplate: |
    {
    "value": "$foo"
    }
variables:
    foo: bar
runtimeVars:
    someId: "$.array[?(@.name=='fubar')].id"
```

### Contribution

...
Todo fill me out
...
standard go styling and contribution guide lines

#### Local testing

The test app is using [pocketbase](https://pocketbase.io/)

```bash
docker build --build-arg "VERSION=0.7.9" -t pocketbase-sample -f sample-app.Dockerfile .
docker run --name=pb-app --detach -p 8090:8090 pocketbase-sample 
```

Then navigate to this [page](http://127.0.0.1:8090/_/?installer#)

enter any username/password combo this is only for local testing
test@example.com/P4s$w0rd123!

See test/test.yaml for an integration style test locally

For OAuth2 Server mocking - you can use/run this [OAuth2 mock server](https://github.com/axa-group/oauth2-mock-server#readme) locally.
