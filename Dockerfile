FROM golang:1.16.5-buster AS builder

ENV REPO_URL=github.com/alidevjimmy/user_microservice_t
ENV GOPATH=/app
ENV APP_PATH=$GOPATH/src/$REPO_URL
WORKDIR $APP_PATH/src
ADD src .
RUN mkdir /out
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /out/main .

FROM scratch
COPY --from=builder /out/main /
CMD ["/main"]