# Se baseia na imagem https://hub.docker.com/_/golang/
FROM golang:1.8

# Copia o diretorio local para o diretorio do container
ADD . $GOPATH/src/github.com/michelaquino/rabbit_mq_test/producer

# Instala a aplicacao
RUN go install github.com/michelaquino/rabbit_mq_test/producer

# Executa a aplicacao quando o container for iniciado
ENTRYPOINT $GOPATH/bin/producer

# Expoe a porta 8080
EXPOSE 8080