# Software Architecture - Tech Challenge

<details>

<summary>Entrega FASE 4 -Arquitetura de Microsserviços</summary>

# Software Architecture - FASE 4 - Tech Challenge

## Requisitos

|Recurso|Versão|Obrigatório|Nota|
|-|-|-|-|
|Docker Desktop| 4.21 ou mais atual|Sim|Necessário para rodar containers das APIs e banco de dados|
|Golang| 1.20|Não|Necessário apenas no caso de rodar localmente sem container|

## O que esse projeto faz e possui
### O que esse projeto faz
Através da API é possível criar um pedido e acompanhar seu status, tendo seu registro persistido no banco de dados.

#### O que esse projeto possui
 - [x] Dockerfile e DockerCompose
 - [x] Documentação para Consumo das API
 - [x] Testes Unitários
 - [x] Banco de dados

## O que esse projeto não faz e débitos técnicos
#### O que esse projeto não faz
- Não se comunica com outros microsserviços;

#### Débitos técnicos
- [ ] Remoção paramêtros *hard coded*, como portas das aplicações.
- [ ] Comunicação com outras aplicações.
- [ ] Algumas partes da aplicação não estão com testes unitários

## Como executar o projeto
### Criar Variáveis de Ambiente
Criar um arquivo nomedo como `.env` na raiz do projeto contendo os seguintes valores.
~~~bash
POSTGRES_USER=postuser
POSTGRES_PASSWORD=postpass
POSTGRES_DB=product
POSTGRES_HOST_PORT=5432
POSTGRES_CONTAINER_PORT=5432
POSTGRES_HOST=database-postgres
POSTGRES_DSN=user=puser password=ppass dbname=order host=database-postgres port=5432 sslmode=disable
~~~
Notas: dada a natureza desse projeto, o arquivo ".env" já está na pasta raiz, assim como, intencionalmente, há valores ***hard coded*** no código.

### Executar o projeto
É possivel executar o projeto através do Makefile, a partir da linha de comando. 
~~~bash
make run-project
~~~
Notas: o comando deve ser efetuado na pasta raiz do projeto

### Executar o Docker
Para executar o projeto, é necessário ter o `Docker Desktop` instalado. Com isso será possível criar as instancias usando o comando `docker compose` via IDE ou linha de comando conforme a seguir:
~~~bash
docker compose -f "docker-compose.yml" up -d --build
~~~
Notas: o comando deve ser efetuado na pasta raiz do projeto

### Utilizar Aplicação & Documentação API
1. Crie um produto `[POST] localhost:8080/api/v1/product` 
2. Retorne uma lista de produtos `[GET] localhost:8080/api/v1/product`
3. Atualize um produto específico `[PUT] localhost:8080/api/v1/product/:id`
4. Remove um produto específico `[DELETE] localhost:8080/api/v1/product/:id` 

A documentação está disponível via Postman com os casos de consumo. É possivel rodar pelo link abaixo, ou copiando a coleção que esta dentro da pasta `docs`.

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/16227218-ad366006-d6e5-41a8-8b14-0e5b79002ac0?action=collection%2Ffork&collection-url=entityId%3D16227218-ad366006-d6e5-41a8-8b14-0e5b79002ac0%26entityType%3Dcollection%26workspaceId%3De76668fb-982b-4d15-ab75-26131dab7174#?env%5BDEV%5D=W3sia2V5IjoiYmFzZV91cmwucmVzdGF1cmFudCIsInZhbHVlIjoibG9jYWxob3N0OjgwODAvYXBpL3YxIiwiZW5hYmxlZCI6dHJ1ZSwidHlwZSI6ImRlZmF1bHQifV0=)