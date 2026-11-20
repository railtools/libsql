FROM ghcr.io/tursodatabase/libsql-server:v0.24.32

# The image's entrypoint reads SQLD_* variables (not LIBSQL_*).
# Store the database on the Railway volume mounted at /data so it survives redeploys.
# The entrypoint creates this directory and chowns it to the sqld user at startup.
ENV SQLD_DB_PATH=/data/libsql.db
# Single node: don't open the replication gRPC port.
ENV SQLD_NODE=standalone

EXPOSE 8080

# Listen on Railway's $PORT (default 8080) so the public domain and healthcheck hit sqld.
ENTRYPOINT ["/bin/bash", "-c", "export SQLD_HTTP_LISTEN_ADDR=\"${SQLD_HTTP_LISTEN_ADDR:-0.0.0.0:${PORT:-8080}}\"; exec /usr/local/bin/docker-wrapper.sh \"$@\"", "--"]
CMD ["/bin/sqld"]
