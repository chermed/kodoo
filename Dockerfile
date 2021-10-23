FROM alpine:3.14.1

WORKDIR /app

COPY . ./ 

RUN chmod a+x kodoo \
    && mkdir -p /.kodoo \
    && chown -R 1000:1000 /.kodoo

USER 1000

ENTRYPOINT ["./kodoo"] 