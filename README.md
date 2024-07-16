# Go User CRUD com GORM Genérico

Este projeto demonstra como criar uma simples app em Go para realizar operações CRUD (Create, Read, Update, Delete) em usuários, utilizando o GORM como ORM (Object-Relational Mapping) e aproveitando os generics do Go para criar um adaptador de repositório genérico e reutilizável.

## Recursos

### CRUD completo para usuários:
- Criação de novos usuários.
- Busca de usuários por ID.
- Atualização de usuários existentes.
- Exclusão de usuários.

### GORM:
- Interação simplificada com o banco de dados.
- Suporte a migrações para gerenciar o esquema do banco de dados.

### Generics:
- Repositório genérico que pode ser usado com qualquer modelo GORM.
- Código mais limpo, conciso e reutilizável.

### Arquitetura Limpa:
- Separação clara de responsabilidades entre domínio, aplicação e infraestrutura.
- Facilidade de manutenção e testabilidade.

## Estrutura do Projeto

```bash
go-user-create/
├── adapter/
│   └── gorm_repository.go   // Adaptador GORM genérico
├── domain/
│   ├── user.go              // Modelo de usuário
│   └── user_repository.go  // Interface do repositório
├── services/
│   └── user_service.go     // Lógica de negócio do usuário
├── main.go                 // Ponto de entrada da aplicação
└── go.mod  
```

## Como usar

- Clone o Repositório:
  ```bash 
    git clone git@github.com:br4tech/go-user-create.git 
  ```
- Instale as Dependências:
  ```bash
    go mod download
  ```

- Execute a Aplicação:
  ```bash
    go run main.go
  ```
