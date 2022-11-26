java -Djava.library.path=./DynamoDBLocal_lib -jar DynamoDBLocal.jar -sharedDb

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

aws --endpoint-url http://localhost:8000 dynamodb describe-table --table-name Music | grep TableStatus

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


aws --endpoint-url http://localhost:8000 dynamodb scan \
     --table-name songplay


aws --endpoint-url http://localhost:8000 \
    dynamodb execute-statement --statement "SELECT * FROM Music   \
                                            WHERE Artist='Acme Band' AND SongTitle='Happy Day'"


aws --endpoint-url http://localhost:8000 \
    dynamodb update-item \
    --table-name Music \
    --key '{ "Artist": {"S": "Acme Band"}, "SongTitle": {"S": "Happy Day"}}' \
    --update-expression "SET AlbumTitle = :newval" \
    --expression-attribute-values '{":newval":{"S":"Updated Album Title"}}' \
    --return-values ALL_NEW


aws --endpoint-url http://localhost:8000 \
    dynamodb execute-statement --statement "UPDATE Music  \
                                            SET AlbumTitle='Updated Album Title'  \
                                            WHERE Artist='Acme Band' AND SongTitle='Happy Day' \
                                            RETURNING ALL NEW *"


aws --endpoint-url http://localhost:8000 \
    dynamodb query \
    --table-name Music \
    --key-condition-expression "Artist = :name" \
    --expression-attribute-values  '{":name":{"S":"Acme Band"}}'


aws --endpoint-url http://localhost:8000 \
    dynamodb execute-statement --statement "SELECT * FROM Music   \
                                            WHERE Artist='Acme Band'"


aws --endpoint-url http://localhost:8000 \
    dynamodb update-table \
    --table-name Music \
    --attribute-definitions AttributeName=AlbumTitle,AttributeType=S \
    --global-secondary-index-updates \
        "[{\"Create\":{\"IndexName\": \"AlbumTitle-index\",\"KeySchema\":[{\"AttributeName\":\"AlbumTitle\",\"KeyType\":\"HASH\"}], \
        \"ProvisionedThroughput\": {\"ReadCapacityUnits\": 10, \"WriteCapacityUnits\": 5      },\"Projection\":{\"ProjectionType\":\"ALL\"}}}]"

 aws --endpoint-url http://localhost:8000 \
    dynamodb query \
    --table-name Music \
    --index-name AlbumTitle-index \
    --key-condition-expression "AlbumTitle = :name" \
    --expression-attribute-values  '{":name":{"S":"Somewhat Famous"}}'


aws --endpoint-url http://localhost:8000 dynamodb execute-statement --statement "SELECT * FROM \"Music\".\"AlbumTitle-index\"  \
                                            WHERE AlbumTitle='Somewhat Famous'"



aws --endpoint-url http://localhost:8000 dynamodb delete-table --table-name Music

aws dynamodb list-tables --endpoint-url http://localhost:8000

aws --endpoint-url http://localhost:8000 dynamodb batch-write-item --request-items file://ProductCatalog.json
