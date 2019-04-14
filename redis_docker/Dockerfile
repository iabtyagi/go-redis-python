FROM redis:4.0.11-alpine

RUN mkdir /redisdata
RUN chown redis:redis -R /redisdata

COPY ./redis.conf /

CMD ["redis-server", "/redis.conf"]
