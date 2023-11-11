# Use multi-stage builds to reduce the final image size (https://docs.docker.com/build/building/multi-stage/)
# Why BusyBox: see https://www.ardanlabs.com/blog/2020/02/docker-images-part1-reducing-image-size.html

FROM golang:1.21 as builder

# Set working directory on container
WORKDIR /usr/src/app

# Prepare dependencies
#COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download && go mod verify

# Copy application files and install the app
ADD ./ ./
RUN go build -v -o /bin/


FROM busybox:1.36

COPY --from=builder /bin/extract /bin/extract

EXPOSE 8280

# Run the compiled app by default
CMD [ "extract" ]