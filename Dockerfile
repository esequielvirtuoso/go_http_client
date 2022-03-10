# --- Base ----
FROM golang:1.16-stretch AS base
WORKDIR $GOPATH/src/github.com/esequielvirtuoso/go_http_client

# ---- Dependencies ----
FROM base AS dependencies
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN	go mod tidy
RUN	go mod vendor

# ---- Test ----
FROM dependencies AS test
COPY . .
RUN go get -u github.com/axw/gocov/gocov
RUN GO111MODULE=off go get -u github.com/matm/gocov-html
RUN go test -v -cpu 1 -failfast -coverprofile=coverage.out -covermode=set ./...
RUN gocov convert coverage.out | gocov-html > /index.html
RUN grep -v "_mock" coverage.out >> filtered_coverage.out
RUN go tool cover -func filtered_coverage.out

# ---- Lint ----
FROM dependencies AS lint
RUN go get github.com/golangci/golangci-lint/cmd/golangci-lint
COPY . .
RUN golangci-lint run -c ./.golangci.yml

# ---- audit ----
FROM dependencies AS audit
COPY go.mod .
COPY .nancy-ignore .
RUN wget https://github.com/sonatype-nexus-community/nancy/releases/download/v1.0.9/nancy-v1.0.9-linux-amd64 -qO /bin/nancy  && chmod +x /bin/nancy
RUN go list -m all | nancy -x $GOPATH/src/github.com/esequielvirtuoso/go_http_client/.nancy-ignore sleuth