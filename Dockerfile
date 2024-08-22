# syntax=docker/dockerfile:1

FROM golang:alpine3.20

# PORT ARG/DEFAULT VALUE
ARG PORT=3000

WORKDIR /app

# INSTALL NPM
RUN apk add --update npm

# INSTALL GIT
RUN apk update && apk add --no-cache git

# COPY OVER FILES
COPY . .

RUN npm install

# DOWNLOAD GO MODULES
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

# GET GO DEPENDENCIES
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go install github.com/joho/godotenv/cmd/godotenv@latest

# GENERATE TAILWINDCSS
RUN npx tailwindcss -i ./css/input.css -o ./css/output.css

# GENERATE TEMPL GO FILES
RUN templ generate

# BUILD
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-todo-app

# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can (optionally) document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE ${PORT}

# Run
CMD [ "/go-todo-app" ]
