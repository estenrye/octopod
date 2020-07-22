[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/estenrye/octopod)

# Octopod

![CI (Build, Lint, Test)](https://github.com/estenrye/octopod/workflows/CI%20(Build,%20Lint,%20Test)/badge.svg)

## What this tool does

This tool exposes an un-authenticated JSON endpoint to export out docker swarm service information.

## How do you run it?

From the commandline:
```bash
docker run --name octopod-server -d -p 9042:9042 -v /var/run/docker.sock:/var/run/docker.sock:ro estenrye/octopod
```

In a docker stack file:
```yaml
version: "3.7"
services:
  server:
    image: estenrye/octopod
    deploy:
      mode: global
      placement:
        constraints:
          - node.role==manager
    ports:
      - target: 9042
        published: 9042
        protocol: tcp
        mode: ingress
    volumes:
      - type: bind
        source: /var/run/docker.sock
        target: /var/run/docker.sock
        read_only: true
```
To deploy:
```bash
docker stack deploy -c ./stack.yml octopod
```

## What APIs does it provide?

| API Endpoint | Description |
| :--- | :--- |
| `/summary` | Returns a list of all services in the swarm and all task<br/>instances for each service. |
| `/summary/{name}` | Returns a list of services that prefix match `{name}`<br/>and all task instances for each matching service. |
| `/services` | Returns a list of all services in the swarm with their<br/>environment variables and service labels. |
| `/services/{name}` | Returns a list of servicesz that prefix match `{name}`<br/>and their environment variables and service labels. |