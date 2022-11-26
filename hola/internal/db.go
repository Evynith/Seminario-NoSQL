 package internal

 import (
    "log"
    "os"
    "fmt"

     "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
    db *dynamodb.DynamoDB

    userTable string
)

func init() {
    dbEndpoint := os.Getenv("DYNAMO_ENDPOINT")
    region := os.Getenv("AWS_REGION")

    userTable = os.Getenv("DYNAMO_TABLE_USER")
    if userTable == "" {
        log.Fatal(`env variable "DYNAMO_TABLE_USER" is required`)
    }

    sess := session.Must(session.NewSession(&aws.Config{
        Endpoint: aws.String(dbEndpoint),
        Region:   aws.String(region),
    }))
    db = dynamodb.New(sess)
}

 func ListTables() {
    // Initialize a session that the SDK will use to load
    // credentials from the shared credentials file ~/.aws/credentials
    // and region from the shared configuration file ~/.aws/config.
    // sess := session.Must(session.NewSessionWithOptions(session.Options{
    //     SharedConfigState: session.SharedConfigEnable,
    // }))

    // Create DynamoDB client
    //svc := dynamodb.New(sess)
    svc := db

    // create the input configuration instance
    input := &dynamodb.ListTablesInput{}

    fmt.Printf("Tables:\n")

    for {
        // Get the list of tables
        result, err := svc.ListTables(input)
        if err != nil {
            if aerr, ok := err.(awserr.Error); ok {
                switch aerr.Code() {
                case dynamodb.ErrCodeInternalServerError:
                    fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
                default:
                    fmt.Println(aerr.Error())
                }
            } else {
                // Print the error, cast err to awserr.Error to get the Code and
                // Message from an error.
                fmt.Println(err.Error())
            }
            return
        }

        for _, n := range result.TableNames {
            fmt.Println(*n)
        }

        // assign the last read tablename as the start for our next call to the ListTables function
        // the maximum number of table names returned in a call is 100 (default), which requires us to make
        // multiple calls to the ListTables function to retrieve all table names
        input.ExclusiveStartTableName = result.LastEvaluatedTableName

        if result.LastEvaluatedTableName == nil {
            break
        }
    }

}
