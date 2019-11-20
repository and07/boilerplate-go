FROM alpine

COPY bin/boilerplate-go .
COPY tpl tpl
CMD ["/boilerplate-go"]
