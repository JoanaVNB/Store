# Store
Aplicação para cadastro de usuários e seus pedidos.
* Use o comando `docker-compose up -d`. Após os containers estarem de pé, execute a aplicação com o comando: `go run main.go`.
* Para testar: vá até o diretório que deseja testar e use o comando: `go test` ou `go test -v`para o modo verboso.
* Para acessar o Kibana: acesse `elasticsearch:5601`.

## v1
* ✅ API Rest: Gin-Gonic;
* ✅ ORM: GORM;
* ✅ Testes: Gomock; e
* ✅ Persistência: MySQL.

## v2
* ✅ API Rest: Gin-Gonic;
* ✅ ORM: GORM;
* ✅ Testes: Gomock **para UseCases e Handlers**;
* ✅ Persistência: MySQL e **Elastic Search**; e
* ✅ **Docker**.

## v3 - Em construção.
* ✅ API Rest: Gin-Gonic;
* ✅ ORM: GORM;
* ✅ Testes: Gomock para UseCases e Handlers;
* ✅ Docker;
* Persistência: MySQL **com** Elastic Search; e
* Cache: **Redis**.

## Links
Abaixo seguem meus links com explicação do código ou de estudo:
* [Explicação dos testes, citando a função Create](https://joanavidon.notion.site/Elastic-Search-a6443921416f4b558dc267abb76fc675)
* Resumos:
* [Elastic Seach](https://www.notion.so/joanavidon/Testes-92702c53edb34fcca788be34eed89a4f)
* [Gorm e Gin](https://app.diagrams.net/#G180mv-hWo-ncgnk3HMDbdmZFh9T1uzAKN)
* [MySQL](https://joanavidon.notion.site/MySQL-e5e2e66d42ff4d03acb5089dc3c3df17)
* [GoLang](https://www.notion.so/joanavidon/Go-150fab3ab8fc4d60a58a025e28d97051)
* [Docker](https://www.notion.so/joanavidon/Docker-5b93972bba3549a2a64bb557c7eeaae9)
