# Flagly
Flagly is a dead-simple feature flag server.

## Getting Started
### 1. Create `flagly.yml`
Example:
```yml
flags:
  - key: new_dashboard
    description: Enables the new dashboard
    environments:
      production:
        enabled: false
        rollout:
          percentage: 25
          stickiness: user_id
      staging:
        enabled: true

environments:
  - production
  - staging
```
### 2. Run the container

```sh
docker run -p 8080:8080 -v ./flagly.yml:/root/flagly.yml ghcr.io/cheetahbyte/flagly:latest
```
## Configurations
### Changing the path of `flagly.yml`
You can change the path of the default `flagly.yml` by providing the -config flag
```sh
docker run -p 8080:8080 -v ./flagly.yml:/root/config.yml ghcr.io/cheetahbyte/flagly:latest -config config.yml
```
Note: You have to make sure the new file gets mounted correctly.

## API Reference
| Method | Path                 | Description                                    | Body                                                               |
| :----- | :------------------- | :--------------------------------------------- | :----------------------------------------------------------------- |
| `GET`  | `/flags/`            | Retrieves a list of all flags.                 | –                                                                  |
| `GET`  | `/flags/:flag`       | Fetches details for a specific flag by key.    | –                                                                  |
| `POST` | `/flags/evaluate`    | Evaluates the status of a flag for a user/env. | `{ "flag": "key", "environment": "production", "user": { "id": "..." } }` |
| `GET`  | `/environments`      | Retrieves a list of all environments.          | –                                                                  |
| `GET`  | `/environments/:env` | Fetches details for a specific environment.    | –                                                                  |


## Future Plans
Future plans for flagly include:
- GitOps (serve config directly from Git repositories)
- Webinterface (view flags and settings directly from the interface)
- More powerful contexts
- Rollout strategies

## Client Integrations
Currently, you would have to write any clients that get flags from the server yourself.
I dont have any intention to build clients.