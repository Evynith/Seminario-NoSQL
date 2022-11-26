DynamoDB

DynamoDB es una base de datos no relacional de clave valor de AWS, esta puede escalar a grandes niveles y que tiene una gran performance y velocidad. El servicio ya se encuentra implementado dentro del paquete de AWS, no hay que hacer nada extra y tecnicamente casi no hay que configurar, ademas tiene muchos modos que automatizan procesos que antes se hacian manualmente como resguardos, dynamo streams, que cada vez que ocurre un cambio a una linea manda un mensaje, ayuda a arquitecturas orientadas a eventos, contiene metricas de logs y threads; y cloud watch de seguridad, etc.
DynamoDB se puede pagar por el uso (almacenamiento, escrituras y lecturas, etc) o por un plan fijo, tiene replicacion automatica sobre las multiples regiones de AWS e integracion con otros servicios de AWS.
Ejemplos de uso son disney plus, netflix, zoom y mercadolibre.

Se puede ver mas sobre este tema en el siguiente video, con enlaces en la descripcion a otros videos interesantes como la presentacion de ML explicando como fue su migracion de 5000 bases de datos de otros servicios a DynamoDB. https://www.youtube.com/watch?v=ybG2Qnucmts

Para este caso yo probe Dynamo de manera local gracias al instalador que provee el sitio de AWS, para ello se requiere algunas consideraciones antes de iniciar como tener instlado el JRE de java y tener credenciales de AWS activas.
Como pso inicial segui ls instrucciones del mismo sitio para aprender la configurcion basica del ambiente y luego los pasos iniciales creando una tabla y modificandola, ademas de cargar un archivo JSON.
Luego de esto cree mi esquema a subir identificando los elementos que debia contener, las claves y el contenido, para ello cree archivos JSON con los items para cada tabla.

El esquema simboliza el uso de un sitio de streaming de musica el cual deberia estar en varias regiones, ademas de permitirse alta escalabilidad y soporte de usuarios. Es comparable al uso de streaming de video como los cuales ya son clientes de esta base de datos.

video presentacion (sin editar): https://youtu.be/y4-Y-wepqRc

Otros enlaces de interes:
AWS-Dynamo:
Descarga local p1: https://docs.aws.amazon.com/es_es/amazondynamodb/latest/developerguide/DynamoDBLocal.html
Descarga local p2: https://docs.aws.amazon.com/es_es/amazondynamodb/latest/developerguide/DynamoDBLocal.DownloadingAndRunning.html#DynamoDBLocal.DownloadingAndRunning.title
Configuracion AWS CLI: https://docs.aws.amazon.com/cli/latest/userguide/getting-started-quickstart.html
Uso p1: https://docs.aws.amazon.com/es_es/amazondynamodb/latest/developerguide/Tools.CLI.html
Uso p2: https://docs.aws.amazon.com/es_es/amazondynamodb/latest/developerguide/getting-started-step-1.html
Tipos atributo: https://docs.aws.amazon.com/es_es/amazondynamodb/latest/APIReference/API_AttributeValue.html
Creacion y atributos de tabla: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateTable.html
Ejemplo datos con JSON: https://docs.amazonaws.cn/en_us/amazondynamodb/latest/developerguide/SampleData.CreateTables.html
Uso de scan: https://docs.aws.amazon.com/es_es/amazondynamodb/latest/developerguide/Scan.html

GO:
SDK: https://aws.amazon.com/sdk-for-go/
Ejemplo de uso nro1: https://docs.aws.amazon.com/es_es/amazondynamodb/latest/developerguide/example_dynamodb_Scenario_GettingStartedMovies_section.html
Ejemplo de uso nro2: https://tutuz-tech.hatenablog.com/entry/2020/06/19/231534
Configuracion: https://docs.aws.amazon.com/es_es/sdk-for-go/v1/developer-guide/configuring-sdk.html
