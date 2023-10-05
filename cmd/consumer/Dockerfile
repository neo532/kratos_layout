FROM hub.com/standard/golang:1.16-debian


ADD ./bin/consumer /home/www/github.com/neo532/kratos_layout/bin/consumer
ADD ./configs /home/www/github.com/neo532/kratos_layout/configs
ADD ./cmd/consumer/supervisord.conf /home/www/github.com/neo532/kratos_layout/supervisord.conf

WORKDIR /home/www/github.com/neo532/kratos_layout

CMD ["supervisord","-c","/home/www/github.com/neo532/kratos_layout/supervisord.conf"]
