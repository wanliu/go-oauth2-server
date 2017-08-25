# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Contact maintainer with any issues you encounter
MAINTAINER Richard Knop <risoknop@gmail.com>

# Set environment variables
ENV PATH /go/bin:$PATH

# Create a new unprivileged user
RUN useradd --user-group --shell /bin/false app

# Cd into the api code directory
WORKDIR /go/src/github.com/wanliu/go-oauth2-server

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/wanliu/go-oauth2-server

# Chown the application directory to app user
RUN chown -R app:app /go/src/github.com/wanliu/go-oauth2-server/

# Use the unprivileged user
USER app

# Install the api program
RUN go install github.com/wanliu/go-oauth2-server

USER root

RUN curl https://s3.cn-north-1.amazonaws.com.cn/jiejie-tools/realize.tar.gz | tar -xzf - -C /go/bin

USER app

# User docker-entrypoint.sh script as entrypoint
ENTRYPOINT ["./docker-entrypoint.sh"]

# Document that the service listens on port 8080.
EXPOSE 8080
