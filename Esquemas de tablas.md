aws --endpoint-url http://localhost:8000 dynamodb create-table \
--table-name songplay \
    --attribute-definitions \
        AttributeName=id,AttributeType=N \
    --key-schema \
        AttributeName=id,KeyType=HASH \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --table-class STANDARD

aws --endpoint-url http://localhost:8000 dynamodb create-table \
    --table-name song \
    --attribute-definitions \
        AttributeName=id,AttributeType=N \
        AttributeName=title,AttributeType=S \
    --key-schema \
        AttributeName=id,KeyType=HASH \
        AttributeName=title,KeyType=RANGE \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --table-class STANDARD

aws --endpoint-url http://localhost:8000 dynamodb create-table \
  --table-name artist \
    --attribute-definitions \
        AttributeName=id,AttributeType=N \
        AttributeName=name,AttributeType=S \
    --key-schema \
        AttributeName=id,KeyType=HASH \
        AttributeName=name,KeyType=RANGE \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --table-class STANDARD

aws --endpoint-url http://localhost:8000 dynamodb create-table \
  --table-name album \
    --attribute-definitions \
        AttributeName=id,AttributeType=N \
        AttributeName=title,AttributeType=S \
    --key-schema \
        AttributeName=id,KeyType=HASH \
        AttributeName=title,KeyType=RANGE \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --table-class STANDARD

aws --endpoint-url http://localhost:8000 dynamodb create-table \
  --table-name user \
    --attribute-definitions \
        AttributeName=id,AttributeType=N \
        AttributeName=username,AttributeType=S \
    --key-schema \
        AttributeName=id,KeyType=HASH \
        AttributeName=username,KeyType=RANGE \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --table-class STANDARD
