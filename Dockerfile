FROM ghcr.io/tursodatabase/libsql-server:latest

# Set required environment variables
ENV LIBSQL_DB_PATH=/data/libsql.db
ENV LIBSQL_HTTP_LISTEN_ADDR=0.0.0.0:8080

# Make sure /data is writable (Railway mounts it as persistent volume)
RUN mkdir -p /data && chown -R 1000:1000 /data

# Expose port
EXPOSE 8080

# Use the built-in entrypoint, just pass in the arguments directly
ENTRYPOINT ["/usr/local/bin/libsql-server"]
CMD [ "--db-file", "/data/libsql.db", "--enable-http", "--http-listen-addr", "0.0.0.0:8080", "--auth-token", "${LIBSQL_AUTH_TOKEN}" ]
