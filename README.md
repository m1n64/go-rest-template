## Go REST API template project

****
Backend: Go 1.21, Gin-gonic, GORM, air and delve

Database: Postgresql, Redis

Logs, metrics: Prometheus, Grafana

Server, proxy: nginx

*****

Dev-mode:
```shell
mkdir grafana
```

```shell
sudo chmod 777 grafana/ -R
```


```shell
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
```

Production:
```shell
mkdir grafana
```

```shell
sudo chmod 777 grafana/ -R
```

```shell
docker-compose up -d
```