FROM alpine as setup
RUN addgroup --gid 123456 -S appgroup && \
    adduser --uid 123456 -S appuser -G appgroup

FROM scratch as production
COPY --from=setup /etc/passwd /etc/passwd
COPY conventional-commit /conventional-commit
USER appuser
ENTRYPOINT ["/conventional-commit"]
