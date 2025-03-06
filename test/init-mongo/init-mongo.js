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
        "url": "http://example1.com",
        "acceptHeaderValue": "text/turtle",
        "publisherId": "987654321",
        "description": "source to be updated",
        "authHeader": {
            "name": "X-API-KEY",
            "value": "MyApiKey"
        }
    }
]);
db.createCollection('reports');
db.reports.insert([
    {
        "id": "data-source-id",
        "reports": {
            "concept": {
                "id": "data-source-id",
                "url": "http://example.com",
                "dataType": "concept",
                "harvestError": false,
                "startTime": "2022-04-06 14:00:07 +0200",
                "endTime": "2022-04-06 14:00:17 +0200",
                "changedCatalogs": [],
                "changedResources": []
            }
        }
    },
    {
        "id": "test-id",
        "reports": {
            "dataset": {
                "id": "test-id",
                "url": "http://url.com",
                "dataType": "dataset",
                "harvestError": false,
                "startTime": "2022-04-06 14:00:07 +0200",
                "endTime": "2022-04-06 14:00:17 +0200",
                "changedCatalogs": [],
                "changedResources": []
            }
        }
    },
    {
        "id": "test-id-2",
        "reports": {
            "publicService": {
                "id": "test-id-2",
                "url": "http://url2.com",
                "dataType": "publicService",
                "harvestError": false,
                "startTime": "2022-04-06 14:00:07 +0200",
                "endTime": "2022-04-06 14:00:17 +0200",
                "changedCatalogs": [],
                "changedResources": []
            },
            "event": {
                "id": "test-id-2",
                "url": "http://url2.com",
                "dataType": "event",
                "harvestError": false,
                "startTime": "2022-04-06 14:00:07 +0200",
                "endTime": "2022-04-06 14:00:17 +0200",
                "changedCatalogs": [],
                "changedResources": []
            }
        }
    },
    {
        "id": "test-id-3",
        "reports": {
            "publicService": {
                "id": "test-id-3",
                "url": "http://url3.com",
                "dataType": "publicService",
                "harvestError": true,
                "startTime": "2022-04-06 14:00:07 +0200",
                "endTime": "2022-04-06 14:00:17 +0200",
                "errorMessage": "error message",
                "changedCatalogs": [],
                "changedResources": []
            },
            "event": {
                "id": "test-id-3",
                "url": "http://url3.com",
                "dataType": "event",
                "harvestError": false,
                "startTime": "2022-04-06 15:00:07 +0200",
                "endTime": "2022-04-06 15:00:17 +0200",
                "changedCatalogs": [],
                "changedResources": []
            }
        }
    },
    {
        "id": "reasoning-test-id",
        "reports": {
            "dataset": {
                "id": "reasoning-test-id",
                "url": "http://url.com",
                "dataType": "dataset",
                "harvestError": false,
                "startTime": "2022-04-06 14:00:17 +0200",
                "endTime": "2022-04-06 14:00:27 +0200",
                "changedCatalogs": [],
                "changedResources": []
            }
        }
    },
    {
        "id": "reasoning-test-id-2",
        "reports": {
            "publicService": {
                "id": "reasoning-test-id-2",
                "url": "http://url2.com",
                "dataType": "publicService",
                "harvestError": false,
                "startTime": "2022-04-06 14:00:17 +0200",
                "endTime": "2022-04-06 14:00:27 +0200",
                "changedCatalogs": [],
                "changedResources": []
            },
            "event": {
                "id": "reasoning-test-id-2",
                "url": "http://url2.com",
                "dataType": "event",
                "harvestError": false,
                "startTime": "2022-04-06 14:00:17 +0200",
                "endTime": "2022-04-06 14:00:27 +0200",
                "changedCatalogs": [],
                "changedResources": []
            }
        }
    },
    {
        "id": "reasoning-test-id-3",
        "reports": {
            "event": {
                "id": "reasoning-test-id-3",
                "url": "http://url3.com",
                "dataType": "event",
                "harvestError": false,
                "startTime": "2022-04-06 14:00:17 +0200",
                "endTime": "2022-04-06 14:00:27 +0200",
                "changedCatalogs": [],
                "changedResources": []
            }
        }
    },
    {
        "id": "ingested",
        "reports": {
            "dataset": {
                "id": "ingested",
                "dataType": "dataset",
                "harvestError": false,
                "startTime": "2022-04-06 14:00:37 +0200",
                "endTime": "2022-04-06 14:00:47 +0200",
                "changedCatalogs": [],
                "changedResources": []
            },
            "publicService": {
                "id": "ingested",
                "dataType": "publicService",
                "harvestError": false,
                "startTime": "2022-04-06 14:01:07 +0200",
                "endTime": "2022-04-06 14:01:17 +0200",
                "changedCatalogs": [],
                "changedResources": []
            },
            "event": {
                "id": "ingested",
                "dataType": "event",
                "harvestError": false,
                "startTime": "2022-04-06 14:00:27 +0200",
                "endTime": "2022-04-06 14:00:37 +0200",
                "changedCatalogs": [],
                "changedResources": []
            }
        }
    }
]);
