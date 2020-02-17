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