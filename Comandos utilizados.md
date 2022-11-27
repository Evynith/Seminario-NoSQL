### Iniciar base de datos desde archivo

```sh
java -Djava.library.path=./DynamoDBLocal_lib -jar DynamoDBLocal.jar -sharedDb
```

### Crear tabla
```sh
aws --endpoint-url http://localhost:8000 dynamodb create-table \
    --table-name Music \
    --attribute-definitions \
        AttributeName=Artist,AttributeType=S \
        AttributeName=SongTitle,AttributeType=S \
    --key-schema \
        AttributeName=Artist,KeyType=HASH \
        AttributeName=SongTitle,KeyType=RANGE \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --table-class STANDARD
```

### Mostrar estado de la tabla
```sh
aws --endpoint-url http://localhost:8000 dynamodb describe-table --table-name Music | grep TableStatus
```

### Insertar elementos en una tabla con API de Dynamo
```sh
aws --endpoint-url http://localhost:8000 dynamodb put-item \
    --table-name Music  \
    --item \
        '{"Artist": {"S": "No One You Know"}, "SongTitle": {"S": "Call Me Today"}, "AlbumTitle": {"S": "Somewhat Famous"}, "Awards": {"N": "1"}}'

aws --endpoint-url http://localhost:8000 dynamodb put-item \
    --table-name Music  \
    --item \
        '{"Artist": {"S": "No One You Know"}, "SongTitle": {"S": "Howdy"}, "AlbumTitle": {"S": "Somewhat Famous"}, "Awards": {"N": "2"}}'

aws --endpoint-url http://localhost:8000 dynamodb put-item \
    --table-name Music \
    --item \
        '{"Artist": {"S": "Acme Band"}, "SongTitle": {"S": "Happy Day"}, "AlbumTitle": {"S": "Songs About Life"}, "Awards": {"N": "10"} }'

aws --endpoint-url http://localhost:8000 dynamodb put-item \
    --table-name Music \
    --item \
        '{"Artist": {"S": "Acme Band"}, "SongTitle": {"S": "PartiQL Rocks"}, "AlbumTitle": {"S": "Another Album Title"}, "Awards": {"N": "8"} }'
```

### Insertar elementos en una tabla con partiQL
```sh
aws --endpoint-url http://localhost:8000 \
                dynamodb execute-statement --statement "INSERT INTO Music  \
                VALUE  \
                {'Artist':'No One You Know','SongTitle':'Call Me Today', 'AlbumTitle':'Somewhat Famous', 'Awards':'1'}"

aws --endpoint-url http://localhost:8000 dynamodb execute-statement --statement "INSERT INTO Music  \
                VALUE  \
                {'Artist':'No One You Know','SongTitle':'Howdy', 'AlbumTitle':'Somewhat Famous', 'Awards':'2'}"

aws --endpoint-url http://localhost:8000 dynamodb execute-statement --statement "INSERT INTO Music  \
                VALUE  \
                {'Artist':'Acme Band','SongTitle':'Happy Day', 'AlbumTitle':'Songs About Life', 'Awards':'10'}"

aws --endpoint-url http://localhost:8000 dynamodb execute-statement --statement "INSERT INTO Music  \
                VALUE  \
                {'Artist':'Acme Band','SongTitle':'PartiQL Rocks', 'AlbumTitle':'Another Album Title', 'Awards':'8'}"
```

### Mostrar todos los elementos de una tabla
```sh
aws --endpoint-url http://localhost:8000 dynamodb scan \
     --table-name songplay
```

### Hacer una consulta a una tabla con partiQL
```sh
aws --endpoint-url http://localhost:8000 \
    dynamodb execute-statement --statement "SELECT * FROM Music   \
                                            WHERE Artist='Acme Band' AND SongTitle='Happy Day'"
```

### Modificar elemento de una tabla con API de Dynamo
```sh
aws --endpoint-url http://localhost:8000 \
    dynamodb update-item \
    --table-name Music \
    --key '{ "Artist": {"S": "Acme Band"}, "SongTitle": {"S": "Happy Day"}}' \
    --update-expression "SET AlbumTitle = :newval" \
    --expression-attribute-values '{":newval":{"S":"Updated Album Title"}}' \
    --return-values ALL_NEW
```

### Modificar elemento de una tabla con partiQL
```sh
aws --endpoint-url http://localhost:8000 \
    dynamodb execute-statement --statement "UPDATE Music  \
                                            SET AlbumTitle='Updated Album Title'  \
                                            WHERE Artist='Acme Band' AND SongTitle='Happy Day' \
                                            RETURNING ALL NEW *"
```

### Hacer una consulta a una tabla con API de Dynamo
```sh
aws --endpoint-url http://localhost:8000 \
    dynamodb query \
    --table-name Music \
    --key-condition-expression "Artist = :name" \
    --expression-attribute-values  '{":name":{"S":"Acme Band"}}'
```

### Hacer una consulta a una tabl con partiQL
```sh
aws --endpoint-url http://localhost:8000 \
    dynamodb execute-statement --statement "SELECT * FROM Music   \
                                            WHERE Artist='Acme Band'"
```

### Agregar un indice a una tabla
```sh
aws --endpoint-url http://localhost:8000 \
    dynamodb update-table \
    --table-name Music \
    --attribute-definitions AttributeName=AlbumTitle,AttributeType=S \
    --global-secondary-index-updates \
        "[{\"Create\":{\"IndexName\": \"AlbumTitle-index\",\"KeySchema\":[{\"AttributeName\":\"AlbumTitle\",\"KeyType\":\"HASH\"}], \
        \"ProvisionedThroughput\": {\"ReadCapacityUnits\": 10, \"WriteCapacityUnits\": 5      },\"Projection\":{\"ProjectionType\":\"ALL\"}}}]"
```

### Consultar un indice con API de Dynamo
```sh
 aws --endpoint-url http://localhost:8000 \
    dynamodb query \
    --table-name Music \
    --index-name AlbumTitle-index \
    --key-condition-expression "AlbumTitle = :name" \
    --expression-attribute-values  '{":name":{"S":"Somewhat Famous"}}'
```
### Consultar un indice con partiQL
```sh
aws --endpoint-url http://localhost:8000 dynamodb execute-statement --statement "SELECT * FROM \"Music\".\"AlbumTitle-index\"  \
                                            WHERE AlbumTitle='Somewhat Famous'"
```
### Eliminar tabla 
```sh
aws --endpoint-url http://localhost:8000 dynamodb delete-table --table-name Music
```

### Listar todas las tablas existentes
```sh
aws dynamodb list-tables --endpoint-url http://localhost:8000
```

### Cargar un JSON
```sh
aws --endpoint-url http://localhost:8000 dynamodb batch-write-item --request-items file://ProductCatalog.json
```
