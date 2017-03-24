# Se baseia na imagem https://hub.docker.com/_/golang/
FROM golang:1.8

# Copia o diretorio local para o diretorio do container
ADD . $GOPATH/src/github.com/michelaquino/rabbit_mq_test/consumer

# Instala a aplicacao
RUN go install github.com/michelaquino/rabbit_mq_test/consumer

# Executa a aplicacao quando o container for iniciado
ENTRYPOINT $GOPATH/bin/consumer