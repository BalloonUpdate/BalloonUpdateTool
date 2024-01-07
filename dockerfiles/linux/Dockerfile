LABEL authors="mraodragon"
# Compile stage
FROM golang:1.21-alpine AS build-env
# Build Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest
ENTRYPOINT ["top", "-b"]