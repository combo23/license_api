# license-api API Documentation

## Endpoints

## End-point: GetLicense
### Method: GET
>```
>/license?key=
>```

### Response Example

```json
{
    "license": "XXXXX-XXXXX-XXXXX-XXXXX-XXXXX",
    "username": "combo23",
    "createdAt": "1723018924",
    "updatedAt": "1723018924",
    "expiresAt": "1723018924",
    "hwid":"cdd31a3e-d3c7-45cc-b271-e3ce4d1e568e",
    "status":"active"
}
```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: CreateLicense
### Method: POST
>```
>/license
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|API_KEY|

### Body (**json**)

```json
{
    "username":"combo23",
    "expires_at":"1723288604"
}
```

### Example Response

```json
{
    "message":"license created",
    "license": {
        "license": "XXXXX-XXXXX-XXXXX-XXXXX-XXXXX",
        "username": "combo23",
        "createdAt": "1723018924",
        "updatedAt": "1723018924",
        "expiresAt": "1723018924",
        "hwid":"cdd31a3e-d3c7-45cc-b271-e3ce4d1e568e",
        "status":"active"
    }
}
```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃


## End-point: Verify
### Method: POST
>```
>/verify
>```

### Headers

|Content-Type|Value|
|---|---|
|Authorization|ADMIN_TOKEN|

### Body (**json**)

```json
{
    "license_key":"XXXXX-XXXXX-XXXXX-XXXXX-XXXXX",
    "hwid":"XXXXX"
}
```

### Example Response

```json
{
    "message": "license verified"
}
```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃


## End-point: Unbind
### Method: POST
>```
>/unbind?key=
>```

### Headers

|Content-Type|Value|
|---|---|
|Authorization|ADMIN_TOKEN|


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃


## End-point: Ban
### Method: GET
>```
>/ban?key=
>```

### Headers

|Content-Type|Value|
|---|---|
|Authorization|ADMIN_TOKEN|

### Example Response

```json
{
    "message": "license banned"
}
```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃


## End-point: Unbind
### Method: GET
>```
>/unbind?key=
>```

### Headers

|Content-Type|Value|
|---|---|
|Authorization|ADMIN_TOKEN|

### Example Response

```json
{
    "message": "license unbound"
}
```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃


## End-point: Update
### Method: POST
>```
>/update
>```

### Headers

|Content-Type|Value|
|---|---|
|Authorization|ADMIN_TOKEN|

### Body (**json**)

```json
{
    "license": "XXXXX-XXXXX-XXXXX-XXXXX-XXXXX",
    "username": "combo23",
    "createdAt": "1723018924",
    "updatedAt": "1723018924",
    "expiresAt": "1723018924",
    "hwid":"cdd31a3e-d3c7-45cc-b271-e3ce4d1e568e",
    "status":"active"
}
```

### Example Response

```json
{
    "message": "license updated"
}
```
