[![CI pipeline](https://github.com/fiap-postech-soat1-group21-stage4/product-api/actions/workflows/github-ci.yml/badge.svg)](https://github.com/fiap-postech-soat1-group21-stage4/product-api/actions/workflows/github-ci.yml) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=fiap-postech-soat1-group21-stage4_product-api&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=fiap-postech-soat1-group21-stage4_product-api) [![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=fiap-postech-soat1-group21-stage4_product-api&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=fiap-postech-soat1-group21-stage4_product-api) [![Bugs](https://sonarcloud.io/api/project_badges/measure?project=fiap-postech-soat1-group21-stage4_product-api&metric=bugs)](https://sonarcloud.io/summary/new_code?id=fiap-postech-soat1-group21-stage4_product-api)

[![CD pipeline](https://github.com/fiap-postech-soat1-group21-stage4/product-api/actions/workflows/github-cd.yml/badge.svg)](https://github.com/fiap-postech-soat1-group21-stage4/product-api/actions/workflows/github-cd.yml)

# Software Architecture - Tech Challenge

<details>

<summary>Entrega FASE 4 -Arquitetura de Microsserviços</summary>

## Requisitos

|Recurso|Versão|Obrigatório|Nota|
|-|-|-|-|
|Docker Desktop| >=4.21 |Sim|Necessário para rodar containers das APIs e banco de dados|
|SonarQube Cli| >=5.0 |Não|Apenas testes ShiftLeft locais|
|Golang| 1.21 |Não|Necessário apenas executar o projeto local e rodar testes|

## O que esse projeto faz e possui
### O que esse projeto faz
Através da API é possível criar e gerir produtos, tendo seu registro persistido no banco de dados.

#### O que esse projeto possui
 - [x] Workflow/ Pipeline para Validação e Deploy da Imagem
 - [x] Dockerfile e DockerCompose
 - [x] Documentação para Consumo das API
 - [x] Testes Unitários e BDD
 - [x] Banco de dados

## O que esse projeto não faz e débitos técnicos
#### Débitos técnicos
- [ ] Comunicação integral com outras aplicações.
- [ ] Cobertura completa dos testes

## Como executar o projeto
### Criar Variáveis de Ambiente
Criar um arquivo nomedo como `.env` na raiz do projeto contendo os seguintes valores.
~~~bash
POSTGRES_USER=puser
POSTGRES_PASSWORD=ppass
POSTGRES_DB=product
POSTGRES_HOST_PORT=5432
POSTGRES_CONTAINER_PORT=5432
POSTGRES_HOST=database-postgres
POSTGRES_DSN=user=puser password=ppass dbname=product host=database-postgres port=5432 sslmode=disable
API_HOST_PORT=8080
API_CONTAINER_PORT=8080
~~~

### Executar o projeto
É possivel executar o projeto através do Makefile, a partir da linha de comando.
~~~bash
make run-project
~~~
Notas: o comando deve ser efetuado na pasta raiz do projeto e deve possuir o .env com as variáveis preenchidas

### Executar o Docker
Para executar o projeto, é necessário ter o `Docker Desktop` instalado. Com isso será possível criar as instancias usando o comando `docker compose` via IDE ou linha de comando conforme a seguir:
~~~bash
docker compose -f "docker-compose.yml" up -d --build
~~~
Notas: o comando deve ser efetuado na pasta raiz do projeto e deve possuir o .env com as variáveis preenchida

### Executar testes
Para executar os testes contidos nos projetos, poderá ser aplicado o seguinte comando:
~~~bash
go test -coverprofile=coverage.out ./... ;    go tool cover -func=coverage.ou
~~~
Notas: o comando executará todos os testes e gerará um relatório de cobertura dos testes

### Utilizar Aplicação & Documentação API
1. Crie um produto `[POST] localhost:8080/api/v1/product` 
2. Retorne uma lista de produtos `[GET] localhost:8080/api/v1/product`
3. Atualize um produto específico `[PUT] localhost:8080/api/v1/product/:id`
4. Remove um produto específico `[DELETE] localhost:8080/api/v1/product/:id` 

A documentação está disponível via Postman com os casos de consumo. É possivel rodar pelo link abaixo, ou copiando a coleção que esta dentro da pasta `docs`.

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/16227218-ad366006-d6e5-41a8-8b14-0e5b79002ac0?action=collection%2Ffork&collection-url=entityId%3D16227218-ad366006-d6e5-41a8-8b14-0e5b79002ac0%26entityType%3Dcollection%26workspaceId%3De76668fb-982b-4d15-ab75-26131dab7174#?env%5BDEV%5D=W3sia2V5IjoiYmFzZV91cmwucmVzdGF1cmFudCIsInZhbHVlIjoibG9jYWxob3N0OjgwODAvYXBpL3YxIiwiZW5hYmxlZCI6dHJ1ZSwidHlwZSI6ImRlZmF1bHQifV0=)
