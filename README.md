
# GoExpert Weather 🌤️

Projeto desenvolvido em Go para consulta de clima atual com base em um CEP. O sistema retorna a temperatura em graus Celsius, Fahrenheit e Kelvin. Desenvolvido pelo **Paulo Nunes**.

## Funcionalidades 📋

- Receber um CEP válido de 8 dígitos.
- Consultar a API ViaCEP para identificar a localização do CEP.
- Utilizar a API WeatherAPI para consultar a temperatura na localização encontrada.
- Converter e retornar a temperatura nos formatos Celsius, Fahrenheit e Kelvin.

## Requisitos 📦

- Docker e Docker Compose instalados.
- Configuração do ambiente com as variáveis:
  - `WEATHER_API_KEY`: Chave da API WeatherAPI para consulta de clima.

## Endereço ativo 🌐

A aplicação está disponível no Google Cloud Run no seguinte endereço:  
[https://goexpert-weather-206501290178.us-central1.run.app](https://goexpert-weather-206501290178.us-central1.run.app)

---

## Exemplos de uso 🛠️

### Com `curl`

#### Consulta local

```bash
curl "http://localhost:8080/weather?cep=01214000"
# Saída esperada:
# {"temp_C":30.2,"temp_F":86.36,"temp_K":303.34999999999997}

curl "http://localhost:8080/weather?cep=01001000"
# Saída esperada:
# {"temp_C":22.1,"temp_F":71.78,"temp_K":295.25}
```

#### Consulta no Google Cloud Run

```bash
curl "https://goexpert-weather-206501290178.us-central1.run.app/weather?cep=01214000"
# Saída esperada:
# {"temp_C":30.2,"temp_F":86.36,"temp_K":303.34999999999997}

curl "https://goexpert-weather-206501290178.us-central1.run.app/weather?cep=01001000"
# Saída esperada:
# {"temp_C":22.1,"temp_F":71.78,"temp_K":295.25}
```

---

## Como executar o projeto 🚀

### Localmente

1. Configure as variáveis de ambiente:
   - `WEATHER_API_KEY`: Insira sua chave de API da WeatherAPI.

2. Execute o servidor:
   ```bash
   go run cmd/main.go
   ```

3. Acesse `http://localhost:8080` para validar que o servidor está rodando.

### Com Docker

1. Configure as variáveis no `docker-compose.yml`:
   ```yaml
   services:
     app:
       environment:
         WEATHER_API_KEY: "sua_chave_api_weather"
   ```

2. Execute:
   ```bash
   sudo podman compose up --build
   ```

3. Acesse `http://localhost:8080` ou a URL publicada.

---

## Estrutura do projeto 📂

```
├── cmd/
│   └── main.go         # Arquivo principal
├── configs/
│   └── config.go       # Configurações do projeto
├── internal/
│   ├── delivery/
│   │   └── rest/
│   │       └── handler.go  # Rota HTTP
│   ├── repository/
│   │   └── zip_code_repository.go  # Consulta ViaCEP
│   └── usecase/
│       └── weather_usecase.go      # Lógica principal
├── docker/
│   └── Dockerfile      # Configuração Docker
├── docker-compose.yml  # Configuração Docker Compose
└── README.md           # Este arquivo
```

## Testes automatizados ✅

1. Configure o ambiente:
   ```bash
   go mod tidy
   ```

2. Execute os testes:
   ```bash
   go test ./internal/repository/... ./internal/usecase/... -v
   ```
---

## 👨‍💻 Autor

**Paulo Henrique Nunes Vanderley**  
- 🌐 [Site Pessoal](https://www.paulonunes.dev/)  
- 🌐 [GitHub](https://github.com/paulnune)  
- ✉️ Email: [paulo.nunes@live.de](mailto:paulo.nunes@live.de)  
- 🚀 Aluno da Pós **GoExpert 2024** pela [FullCycle](https://fullcycle.com.br)

---

## 🎉 Agradecimentos

Este repositório foi desenvolvido com muita dedicação para a **Pós GoExpert 2024**. Agradeço à equipe da **FullCycle** por proporcionar uma experiência incrível de aprendizado! 🚀
