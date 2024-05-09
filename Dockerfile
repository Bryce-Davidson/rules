#build stage
FROM golang:alpine
WORKDIR /app
COPY ./.. /app
RUN go build -o /app/battlesnake -v /app/cli/battlesnake/main.go


EXPOSE 8000

# TODO
CMD ["/app/battlesnake", "-h"]
