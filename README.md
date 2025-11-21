# Container Registry

A lightweight Docker registry setup with a sample containerized application.

## Dependencies

- Docker
- Go 1.x (for local endpoint testing)
- [Task](https://taskfile.dev/) (optional, for running commands)

## What's Included

- **Docker Registry**: Local Docker registry server (port 5000)
- **Sample Image**: Alpine-based container that makes HTTP POST requests
- **Test Endpoint**: Go HTTP server to receive container requests (port 3000)

## Quick Start

### Run the Local Registry

```bash
task start-registry
# or
docker run --restart=always \
  -p 5000:5000 \
  -v ./registry/data:/var/lib/registry \
  -v ./registry/config/config.yml:/etc/docker/registry/config.yml:ro \
  registry:2
```

### Build and Run the Sample Image

```bash
# Build the image
task build

# Run the container
task run
```

### Test with Local Endpoint

```bash
# Start the test endpoint
task start-endpoint

# In another terminal, run the container
task run
```

## Caddy Configuration

```
registry.curiola.com:80 {
    encode zstd gzip

    log {
        output stdout
    }

    @protected {
        path /v2/*
    }

    basic_auth @protected {
        ci-user $2a$14$VTH/mtlPTELmJfq46WU2QebrMnVdXKi/s4Q/780yAlFkS3LFMfIqK
    }

    reverse_proxy 127.0.0.1:5000
})
```

## Configuration

### Registry Config

The registry config is located at `registry/config/config.yml`. Data is persisted in `registry/data/`.

**Important for Caddy/Cloudflare Tunnels:**

The `relativeurls: true` setting in the HTTP config is critical for working with reverse proxies like Caddy and Cloudflare Tunnels. Without this, Docker clients will receive absolute URLs pointing to the internal registry address instead of your public domain.

```yaml
http:
  addr: :5000
  relativeurls: true  # Required for reverse proxy setups
```

### Sample Image Build Arguments

- `API_TOKEN`: Bearer token for authentication
- `TARGET_URL`: Endpoint URL (default: https://amaze-api.curiola.com/hello)
