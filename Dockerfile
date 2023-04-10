FROM alpine as setup
RUN addgroup --gid 123456 -S appgroup
RUN adduser --uid 123456 -S appuser -G appgroup
RUN apk add ca-certificates

FROM scratch as production
COPY --from=setup /etc/passwd /etc/passwd
COPY --from=setup /etc/ssl /etc/ssl
COPY conventional-commit /conventional-commit
USER appuser
ENTRYPOINT ["/conventional-commit"]
