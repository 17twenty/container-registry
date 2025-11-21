# Dockerfile
FROM alpine:3.20

RUN apk add --no-cache curl

# Build arguments to accept values at build time
ARG API_TOKEN
ARG TARGET_URL=https://amaze-api.curiola.com/hello

# Bake the values into the image as environment variables
ENV API_TOKEN=${API_TOKEN}
ENV TARGET_URL=${TARGET_URL}

COPY entrypoint.sh /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
