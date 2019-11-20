FROM swaggerapi/swagger-ui

COPY ./openapi.json /app/openapi.json
COPY ./openapi.yaml /app/openapi.yaml
COPY ./bootstrap.sh /app/bootstrap.sh

CMD ["sh", "/app/bootstrap.sh"]

RUN cd ./

RUN ls -la