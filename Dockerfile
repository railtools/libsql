FROM ghcr.io/tursodatabase/libsql-server:latest

# Use /data as the standard Railway persistent storage location
ENV LIBSQL_DB_PATH=/data/libsql.db
ENV LIBSQL_HTTP_LISTEN_ADDR=0.0.0.0:8080
ENV LIBSQL_AUTH_TOKEN=${LIBSQL_AUTH_TOKEN}

# Create the directory and ensure correct permissions
RUN mkdir -p /data && chown -R 1000:1000 /data

# Expose the default LibSQL HTTP port
EXPOSE 8080

CMD ["sh", "-c", "libsql-server --db-file $LIBSQL_DB_PATH --enable-http --http-listen-addr $LIBSQL_HTTP_LISTEN_ADDR --auth-token $LIBSQL_AUTH_TOKEN"]

