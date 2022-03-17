db = db.getSiblingDB('fdk-harvest-admin');
db.createCollection('datasources');
db.datasources.insert([
    {
        "id": "test-id",
        "dataSourceType": "DCAT-AP-NO",
        "dataType": "dataset",
        "url": "http://url.com",
        "acceptHeaderValue": "text/turtle",
        "publisherId": "123456789",
        "description": "test source",
        "authHeader": {
            "name": "X-API-KEY",
            "value": "MyApiKey"
        }
    },
    {
        "id": "test-id-2",
        "dataSourceType": "CPSV-AP-NO",
        "dataType": "publicService",
        "url": "http://url2.com",
        "acceptHeaderValue": "text/turtle",
        "publisherId": "111222333",
        "description": "test source 2"
    },
    {
        "id": "test-id-3",
        "dataSourceType": "CPSV-AP-NO",
        "dataType": "publicService",
        "url": "http://url3.com",
        "acceptHeaderValue": "text/turtle",
        "publisherId": "123456789",
        "description": "test source 3"
    },
    {
        "id": "to-be-deleted",
        "dataSourceType": "SKOS-AP-NO",
        "dataType": "concept",
        "url": "http://example.com",
        "acceptHeaderValue": "text/turtle",
        "publisherId": "987654321",
        "description": "source to be deleted"
    },
    {
        "id": "to-be-updated",
        "dataSourceType": "SKOS-AP-NO",
        "dataType": "concept",
        "url": "http://example.com",
        "acceptHeaderValue": "text/turtle",
        "publisherId": "987654321",
        "description": "source to be updated",
        "authHeader": {
            "name": "X-API-KEY",
            "value": "MyApiKey"
        }
    }
]);
