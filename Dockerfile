FROM ghcr.io/tursodatabase/libsql-server:latest

ENV LIBSQL_DB_PATH=/data/libsql.db
ENV LIBSQL_HTTP_LISTEN_ADDR=0.0.0.0:8080
ENV LIBSQL_AUTH_TOKEN=${LIBSQL_AUTH_TOKEN}

CMD ["sh", "-c", "libsql-server --db-file $LIBSQL_DB_PATH --enable-http --http-listen-addr $LIBSQL_HTTP_LISTEN_ADDR --auth-token $LIBSQL_AUTH_TOKEN"]

