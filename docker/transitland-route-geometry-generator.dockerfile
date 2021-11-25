FROM golang:1.16 as builder

WORKDIR /app

RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin


COPY . /app
RUN task build-transitland-route-geometry-generator-prod
RUN ls -lah build && chmod +x build/transitland-route-geometry-generator && pwd

FROM alpine

LABEL maintainer="Phoops info@phoops.it"
LABEL environment="production"
LABEL project="transitland-route-geometry-generator"
LABEL service="transitland-route-geometry-generator"

WORKDIR /app
COPY --from=builder /app/build/transitland-route-geometry-generator /app

ENTRYPOINT ["./transitland-route-geometry-generator"]