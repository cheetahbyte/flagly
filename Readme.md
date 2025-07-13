# Flagly
Flagly is a dead-simple feature flag server.

## Getting Started
### 1. Create `flagly.yml`
Example:
```yml
flags:
  - key: new_login
    description: Aktiviert den neuen Login
    enabled: true
    conditions:
      - environments: ["production", "staging"]
  - key: use_beta_api
    description: Verwendet die neue Beta API
    enabled: true
    conditions:
      - environments: ["staging"]

environments:
  - production
  - staging
```
### 2. Run the container

```sh
docker run -p 8080:8080 -v ./flagly.yml:/root/flagly.yml flagly:latest
```


## API Reference
| Method   | Path                  | Description                                       | 
| :------- | :-------------------- | :------------------------------------------------ | 
| `GET`    | `/flags/`       | Retrieves a list of all flags.                            | 
| `GET`    | `/flags/:flag`  | Fetches details for a specific flag by their key.           |
| `GET`    | `/flags/:flag/enabled`       | Fetches the status for a certain flag.        | 
| `GET`    | `/environments`  | Retrieves a list of all environments                   |
| `GET`    | `/environments/:env`  | Fetches details for a specific environment         | 

## Future Plans
Future plans for flagly include:
- GitOps (serve config directly from Git repositories)
- Webinterface (view flags and settings directly from the interface)
- More powerful contexts
- Rollout strategies

## Client Integrations
Currently, you would have to write any clients that get flags from the server yourself.
I dont have any intention to build clients.