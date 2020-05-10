# Clean Architecture Implementation

This project *tries* to implement a custom version of Uncle Bob's Clean Architecture.

<!--
![alt-text](https://miro.medium.com/max/1400/1*B7LkQDyDqLN3rRSrNYkETA.jpeg "Uncle Bob's Clean Architecture")-->

## General scaffolding

```
internal
+-- entities
+-- usecase
|   +-- gateway
+-- interface
|   +-- controller
|   +-- dto
+-- infrastructure
|   +-- repository
|   +-- service
|   +-- delivery
```

## Explanation

Inside **internal** folder, you will find the following directories.

### entities
Commonly known as business model. Contains structures (domain entities an domain value objects) with data an functions non-dependant of outter layers.

### usecase
Application-wide logic. Functions for business logic orchestration.

#### gateway
Interfaces representing the invocation to outer layers, needed for the application logic: external services, data repositories, events, etc.

### interface
Adaptation between application input and output.

#### controller
Input data conversion and handling. Converts DTOs into understandable information for inner layers. Calls to usecase/interactors functions. Takes usecase/interactors results and convert them into DTOs. 

#### dtos
Data-transfer objects. Structures for data exchange from/to application.

### infrastructure
Infrastructure layer.

#### repository
Data access logic. Implements usecase/gateway interfaces. Includes DB <=> Model data conversion, but doesn't contain business logic. Tipically CRUD.

#### service
External services (WS, REST, gprc, etc.) access logic. Implements usecase/gateway interfaces. Includes external models <=> internal model data adaptation, but doesn't contain business logic.

#### event
Event raising logic. Implements usecase/gateway interfaces.

#### delivery
Routes and requests handling configuration.
