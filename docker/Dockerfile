FROM golang:1.20.1-bullseye

RUN apt update
RUN apt install -y unzip

WORKDIR /workspace
COPY go.mod .
# COPY go.sum .
RUN go mod download

# for vscode-plugin
RUN go install -v golang.org/x/tools/gopls@latest
