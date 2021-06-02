FROM golang:alpine AS build-env
COPY . /src
RUN cd /src && go build -o ssp main.go

FROM alpine
COPY --from=build-env /src/ssp /ssp
ENTRYPOINT /ssp
