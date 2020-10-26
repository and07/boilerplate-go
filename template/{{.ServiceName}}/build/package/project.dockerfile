FROM alpine

COPY bin/{{.ServiceName}} .
COPY tpl tpl
CMD ["/{{.ServiceName}}"]
