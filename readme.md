# 🚀 go-bootstrap

![Go Version](https://img.shields.io/badge/Go-1.20%2B-00ADD8?style=for-the-badge&logo=go)
![Fiber](https://img.shields.io/badge/Fiber-v2-00C7B7?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Status](https://img.shields.io/badge/status-em%20desenvolvimento-yellow?style=for-the-badge)

Template inicial para APIs REST em Go utilizando Fiber, GORM e autenticação com JWT.

Ideal para iniciar projetos backend com uma estrutura organizada e pronta para escalar.

---

## 📦 Tecnologias utilizadas

- **Go**
- **Fiber**
- **GORM**
- **SQLite**
- **JWT**
- **UUID**

---

## 📁 Estrutura do projeto

```bash
go-bootstrap/
│
├── main.go
├── go.mod
├── go.sum
├── readme.md
│
├── controllers/
│
├── database/
│   ├── migrations/
│   │   └── migration.go
│   └── database.go
│
├── files/
│   └── database.sqlite
│
├── models/
│   └── example.go
│
├── server/
│   ├── routes/
│   │   └── router.go
│   └── server.go
│
├── utils/
    ├── Date.go
    └── SecurityJWT.go
```

---

## ⚙️ Pré-requisitos

- Go 1.20+
- Git

---

## ▶️ Como rodar o projeto

### 1. Clone o repositório

```bash
git clone https://github.com/RafaelMoreira96/go-bootstrap.git
cd go-bootstrap
```

### 2. Instale as dependências

```bash
go mod tidy
```

### 3. Execute o projeto

```bash
go run main.go
```

A aplicação estará disponível em:

```
http://localhost:8000
```

---

## 🗄️ Banco de dados

Por padrão o projeto usa **SQLite**, mas pode ser alterado facilmente para:

- PostgreSQL
- MySQL

Arquivo responsável:

```bash
database/database.go
```

---

## 🧱 Arquitetura

Separação inspirada em Clean Architecture:

- Controllers → entrada HTTP e regras de negócio
- Models → entidades  
- Server → configurações do servidor e rotas  

---

## 🧪 Melhorias futuras

- [ ] Docker  
- [ ] Testes automatizados  
- [ ] Swagger  
- [ ] Variáveis de ambiente (.env)  
- [ ] CI/CD  

---

## 🤝 Contribuição

Pull requests são bem-vindos!

---

## 📄 Licença

MIT

---

## 👨‍💻 Autor

Rafael Moreira  
Backend Developer — Go | Java | Python