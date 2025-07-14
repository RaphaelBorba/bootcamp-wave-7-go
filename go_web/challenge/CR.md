# Desafio Go Web - Code Review Luiz

## Pontos Positivos

- **Tratamento de erros:** Erros de leitura e conversão de dados do CSV são tratados com mensagens descritivas.
- **Bom uso de interfaces:** Interfaces bem definidas para repositórios e serviços.
- **Uso de contexto:** Integração do `context.Context` em métodos que acessam dados.

---

## Pontos de Melhoria


- **Mensagens de erro repetitivas:** As mensagens retornadas pelo handler são semelhantes. Adicionar error de Not Found quando o parametro não existir, seria uma boa.
- **Inconsistência nas interfaces:** A interface `RepositoryTicketMap` define `GetTicketByDestinationCountry`, mas o método implementado no repository é `GetTicketsByDestinationCountry`.
- **Pouco logging:** Incluir logs durante as chamadas HTTP facilitaria o debug e monitoramento da aplicação.
- **Valores mágicos:** O valor `1` passado como `lastId` no repositório não tem explicação e poderia ser extraído dinamicamente do CSV.