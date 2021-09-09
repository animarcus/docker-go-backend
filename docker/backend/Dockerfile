FROM golang:1.16

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -v -o webserver

CMD ["./wait-for-it.sh", "db:3306", "--", "./webserver"]



######################################


# FROM golang:1.16

# WORKDIR /app

# COPY . ./

# RUN go mod download

# RUN ["go", "get", "github.com/go-playground/justdoit"]

# ENTRYPOINT justdoit -watch="./" -include="(.+\\.go|.+\\.c)$" -build="go build webserver" -run="./webserver"

# RUN go get github.com/githubnemo/CompileDaemon
# CMD [ "/bin/bash" ]
# RUN go build -v -o webserver
# CMD CompileDaemon -polling -build="go build -v -o webserver" -command="./webserver"
# CMD ["./webserver"]


# FROM golang:1.16

# WORKDIR /app

# COPY . /app

# RUN go mod download


# RUN GO111MODULE=off go get github.com/oxequa/realize

# ENTRYPOINT realize start
