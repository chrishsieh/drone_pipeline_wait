FROM alpine
ADD ./drone_pipeline_wait /bin/
RUN chmod 755 /bin/drone_pipeline_wait \
    && apk -Uuv add ca-certificates
ENTRYPOINT /bin/drone_pipeline_wait
