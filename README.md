# juridigo_api_interacao
Servidor


# Obtenção e inicialização do projeto

```
cd go/src/github.com/
mkdir juridigo
cd juridigo
git clone https://github.com/andre250/juridigo_api_interacao.git
cd juridigo_api_interacao/
dep ensure 
gin -i main.go
```

# Endpoints

## Usuario
```
http://.../usuario?lat={lat}&long={long}&dist={dist}&rank={rank}

METHOD: GET

Descrição: Lista usuarios baseados nos critérios passados
Parametros:
 - lat = float64 (Latitude do usuario)
 - long = float64 (Longitude do usuario)
 - dist = float64 (Distância em KM de referencia)
 - rank = []int (Rankings de retorno) 
```

## Trabalho
```
http://.../trabalho

METHOD: POST

Descrição: Criação de um trabalho

Body: 
{
  "tipoTrabalho": [String],
  "categoriaTrabalho": [String],
  "usuarioAtribuido": [String],
  "descricao": [String],
  "rotulo": [String],
  "prazo": [String | Unix timestamp],
  "valor": [Float64],
  "taxa": [Float64],
  "classificacao": [uint],
  "multiplicador": [uint],
  "localizacao": {
    "pais": [String],
    "estado": [String],
    "cidade": [String],
    "regiao": [String],
    "rua": [String],
    "numero": [String],
    "complemento": [String],
    "latitude": [Float64],
    "longitude": [Float64]
  },
  "audiencia": {
    "tipo": [String],
    "numero": [String],
    "partes": [
      [String],
      [String]
    ],
    "vara": [String],
    "nomeResponsavel": [String],
    "telefoneResponsavel": [String],
    "forum": [String]
  }
}
```
```
http://.../trabalho?id={id}

METHOD: PUT

Descrição: Atualização de um trabalho

Parametros:
 - id = string (id do trabalho atualizado)

Body: 
{
  "campo": "valor"
}

OBS. A Atualização deve seguir regra de escala de objeto, por exemplo:

- Atualizar Audiencia

Body:
{
    "audiencia": {
        "tipo": [String]
    }
}

- Autalizar Descricao

Body:
{
    "descricao": [String]
}
```

```
http://.../trabalho?usuario={id}&status={status}

METHOD: GET

Descrição: Lista trabalhos de um usuario

Parametros:
    - id = string (id do usuario) *obrigatório
    - status = array de string (status do trabalho) *opcional
```

```
http://.../trabalho/{id}

METHOD: GET

Descrição: Lista detalhes de um trabalho

Parametros:
    - id = string (id do trabalho)
```

## Proposta

```
http://.../proposta

METHOD: POST

Descrição: Cria uma proposta

Body:
{
  "idTrabalho": [String],
  "usuarioRelacionado": [String],
  "tipoTrabalho": [String],
  "categoriaTrabalho": [String],
  "empresa": [String],
  "rotulo": [String],
  "prazo": [String | unix timestamp],
  "localizacao": {
    "pais": [String],
    "estado": [String],
    "cidade": [String],
    "regiao": [String],
    "rua": [String],
    "numero": [String],
    "complemento": [String],
    "latitude": [Float64],
    "longitude": [Float64]
  },
  "audiencia": {
    "tipo": [String],
    "numero": [String]
  }
}
```

```
https://.../proposta?usuario={id}&status={status}

METHOD: GET

Descrição: Busca propostas de um usuario

Parametros:
    - usuario = string (id do usuario)
    - status = array de string (status do trabalho) *opcional 
```

```
http://.../proposta?id={id}

METHOD: PUT

Descrição: Atualização de uma proposta

Parametros:
 - id = string (id da proposta atualizada)

Body: 
{
  "campo": "valor"
}

OBS. A Atualização deve seguir regra de escala de objeto, por exemplo:

- Atualizar Tipo da Audiencie

Body:
{
    "audiencia": {
        "tipo": [String]
    }
}

- Autalizar Rótulo

Body:
{
    "rotulo": [String]
}
```

## Fluxo

```
http://.../fluxo

METHOD: POST

Descrição: Cria um fluxo

Body:
{
  "idTrabalho": [String],
  "etapas": [
    {
      "etapaId": [String],
      "prazo": [String | unix timestamp],
      "status": [String],
      "nome": [String],
      "descricao": [String],
      "usuario": [
        {
          "id": [String],
          "nome": [String],
          "imagemPerfil": [String],
          "dataAtribuido": [String | unix timestamp],
          "dataAtualizado": [String | unix timestamp]
        }
      ],
      "localizacao": {
        "pais": [String],
        "estado": [String],
        "cidade": [String],
        "regiao": [String],
        "rua": [String],
        "numero": [String],
        "complemento": [String],
        "latitude": [Float64],
        "longitude": [Float64]
      }
    }
  ]
```


```
http://.../fluxo?id={id}

METHOD: PUT

Descrição: Atualização de um fluxo

Parametros:
 - id = string (id do fluxo atualizado)

Body: 
{
  "campo": "valor"
}

OBS. A Atualização deve seguir regra de escala de objeto, por exemplo:

- Atualizar Uma etapa

Body:
{
     "etapas": [
        {
        "prazo": [String | unix timestamp]
        }
    ]
}

```


```
https://.../fluxo?trabalho={id}

METHOD: GET

Descrição: Busca fluxos de um trbaalho

Parametros:
    - trabalho = string (id de trabalho)
```