FROM hub.com/standard/golang:1.21-debian


ADD ./bin/api /home/www/github.com/neo532/kratos_layout/bin/api
ADD ./configs /home/www/github.com/neo532/kratos_layout/configs
ADD ./cmd/api/supervisord.conf /home/www/github.com/neo532/kratos_layout/supervisord.conf

WORKDIR /home/www/github.com/neo532/kratos_layout

EXPOSE 8501
EXPOSE 9501

CMD ["bin/api","-conf","configs/config.yaml"]
