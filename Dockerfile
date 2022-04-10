FROM golang:1.18-bullseye as base
  ENV APP_DIR /usr/app

  WORKDIR ${APP_DIR}

  RUN apt-get update && apt-get upgrade -y

  COPY . .

  RUN go mod vendor

FROM base as development
  RUN apt-get install curl
  RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

  EXPOSE 3001

  ENTRYPOINT [ "air", "-c", ".air.toml" ]

FROM base as production
  RUN apt-get install make
  RUN make build

  ENTRYPOINT [ "./bin/server" ]