FROM scratch
ARG TARGETARCH
COPY ./dist/randgen_linux_${TARGETARCH}*/ /
WORKDIR /workspace
ENTRYPOINT ["/randgen"]