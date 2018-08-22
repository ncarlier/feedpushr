#########################################
# Build stage
#########################################
FROM golang:1.10 AS builder

# Repository location
ARG REPOSITORY=github.com/ncarlier

# Artifact name
ARG ARTIFACT=feedpushr

# Copy sources into the container
ADD . /go/src/$REPOSITORY/$ARTIFACT

# Set working directory
WORKDIR /go/src/$REPOSITORY/$ARTIFACT

# Build the binary
RUN make

# Build plugins
ENV REPOSITORY=${REPOSITORY}
ENV ARTIFACT=${ARTIFACT}
RUN git clone https://${REPOSITORY}/${ARTIFACT}-contrib.git plugins \
      && make -C plugins plugins

#########################################
# Distribution stage
#########################################
FROM alpine

# Repository location
ARG REPOSITORY=github.com/ncarlier

# Artifact name
ARG ARTIFACT=feedpushr

# Fix lib dep
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# Install root certificates
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Install binary and default scripts
COPY --from=builder /go/src/$REPOSITORY/$ARTIFACT/release/$ARTIFACT-linux-amd64 /usr/local/bin/$ARTIFACT

# Install plugins
COPY --from=builder /go/src/$REPOSITORY/$ARTIFACT/plugins/release/*.so .
RUN for file in *.so; do mv $file `echo $file | sed 's/-linux-amd64//'`; done

# Define command
CMD feedpushr

