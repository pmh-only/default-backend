FROM scratch

ARG TARGETARCH

ARG user=1000
ARG group=1000

USER $user:$group
WORKDIR /app

COPY --chown=$user:$group default-${TARGETARCH} /app/default
COPY --chown=$user:$group ./skins /app/skins

ENTRYPOINT ["/app/default"]
