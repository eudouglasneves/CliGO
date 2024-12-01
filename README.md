# Load Tester CLI

## Descrição

Uma ferramenta CLI em Go para realizar testes de carga em serviços web. Permite configurar a URL, o número total de requisições e o nível de concorrência. O relatório ao final exibe informações como tempo total, número de requisições bem-sucedidas e distribuição de códigos de status HTTP.

---

## Parâmetros CLI

- `--url`: URL do serviço a ser testado. **Obrigatório.**
- `--requests`: Número total de requisições. (Padrão: 100)
- `--concurrency`: Número de requisições simultâneas. (Padrão: 10)

---

## Uso com Docker

### Build da imagem  
    docker build -t loadtester .

### Execução
    docker run loadtester --url=http://google.com --requests=1000 --concurrency=10


### Relatório de Saída
    Ao final da execução, o sistema exibe um relatório como o exemplo abaixo:

    --- Relatório de Teste de Carga ---
    URL Testada: http://google.com
    Número Total de Requests: 1000
    Concorrência: 10
    Tempo Total: 15.3s
    Requests com Status 200: 950

    Outros Status HTTP:
    404: 30
    500: 20

### Licença
    Este projeto é open-source e pode ser utilizado e modificado livremente.

### Testes
    Para validar o funcionamento, experimente executar o Docker localmente com URLs reais e diferentes configurações de concorrência. Isso ajudará a garantir que a aplicação lide corretamente com cargas altas e erros HTTP.