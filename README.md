[![Build Status](https://travis-ci.org/allisonverdam/best-credit-card.svg?branch=master)](https://travis-ci.org/allisonverdam/best-credit-card) [![Coverage Status](https://coveralls.io/repos/github/allisonverdam/best-credit-card/badge.svg?branch=master)](https://coveralls.io/github/allisonverdam/best-credit-card?branch=master)

# Best Credit Card
Adiministre melhor os seus cartões de crédito, com nossa api você usará sempre o melhor cartão de crédito para a sua compra.

A api atualmente está rodando em um servidor no heroku no seguinte endereço: https://best-credit-card.herokuapp.com


## Para rodar o projeto no seu computador

É necessário ter o Go instalado no seu computador,
se essa é a primeira vez que você vai usar go, siga [as instruções](https://golang.org/doc/install) para
instalar. Utilizamos a versão 1.8.

Depois de instalar execute os seguintes comandos:
```shell
# pegando o projeto
go get https://github.com/allisonverdam/best-credit-card

# utilizaremos o glide para fazer o controle de versão das dependencias do projeto
go get -u github.com/Masterminds/glide

# entre na pasta do projeto e baixe as dependencias
cd $GOPATH/allisonverdam/best-credit-card
make depends   # ou "glide up"
```
Agora temos que criar um database no postgres, escolhemos o nome `best_credit_card`, tem um script pronto para gerar nosso banco, está na pasta `testdata/db.sql`.

Você pode configurar a conexão com o banco no arquivo `config/app.yaml` ou alterando a variavel de ambiente `API_DSN` assim: 
```
postgres://<username>:<password>@<server-address>:<server-port>/<db-name>
```

Agora já podemos executar o projeto, execute o comando abaixo.

```shell
go run main.go
```

Ou simplemente `make` se estiver disponível no seu computador.

```shell
make
```

A aplicação vai iniciar na porta 8080 por padrão.


## Documentação

Todos os endpoints estão documentados com exemplos de request e response, a documentação se encontra aqui no projeto, mas também é possível acessar [clicando aqui.](https://allisonverdam.github.io/best-credit-card/doc)

E também temos uma coleção publicada pronta para utilizar no postman, [clique aqui.](https://documenter.getpostman.com/view/659591/best-credit-card-unique-wallet/77h84Nt)
