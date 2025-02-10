
# runtime container
FROM alpine:3.18

WORKDIR /app
COPY ./build/app-main /app/

EXPOSE 8080
ENTRYPOINT ["/app/app-main"]