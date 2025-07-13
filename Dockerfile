FROM ghcr.io/tursodatabase/libsql-server:latest

# Set Railway-compatible config
ENV LIBSQL_DB_PATH=/data/libsql.db
ENV LIBSQL_HTTP_LISTEN_ADDR=0.0.0.0:8080
ENV LIBSQL_ENABLE_HTTP=true

RUN mkdir -p /data && chown -R 1000:1000 /data

EXPOSE 8080

