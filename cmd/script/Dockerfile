FROM hub.com/standard/golang:1.16-debian

ADD ./configs /home/www/github.com/neo532/kratos_layout/configs
ADD ./bin/script /home/www/github.com/neo532/kratos_layout/bin/script
ADD ./cmd/script/supervisord.conf /home/www/github.com/neo532/kratos_layout/supervisord.conf

WORKDIR /home/www/github.com/neo532/kratos_layout

CMD ["supervisord","-c","/home/www/github.com/neo532/kratos_layout/supervisord.conf"]
