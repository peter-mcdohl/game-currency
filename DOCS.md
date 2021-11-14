# API

## Add Currency

URL : /v1/currency

Method : POST

Request Body:
```
{
    "id": 1,
    "name": "Knut"
}
```

Success Response (code: 201)
```
{
    "message": "Succuessfully add currency"
}
```

Fail Response (code: 400)
```
{
    "message": "Failed to add currency"
}
```

## Get Currency

URL : /v1/currency

Method : GET

Success Response (code: 200)
```
{
    "data": [
        {
            "id": 1,
            "name": "Knut"
        },
        {
            "id": 2,
            "name": "Sickle"
        },
        {
            "id": 3,
            "name": "Galleon"
        }
    ]
}
```

## Add Conversion Rate

URL : /v1/currency/conversion-rate

Method : POST

Request Body:
```
{
    "currency_id_from": 2,
    "currency_id_to": 1,
    "currency_rate": 29
}
```

Success Response (code: 201)
```
{
    "message": "Succuessfully add conversion rate"
}
```

Fail Response (code: 400)
```
{
    "message": "Failed to add conversion rate"
}
```

## Convert Currency

URL : /v1/currency/convert

Method : POST

Request Body:
```
{
    "currency_id_from": 1,
    "currency_id_to": 2,
    "amount": 580
}
```

Success Response (code: 200)
```
{
    "result": 20
}
```

Fail Response (code: 400)
```
{
    "message": "Failed to convert"
}
```


# Database

## Currency

Table name: currency

Schema:
- id: bigint (PK)
- name: text
- created_at: timestamp
- updated_at: timestamp
- deleted_at: timestamp (nullable)

## Coversion Rate

Table name: conversion_rate

Schema:
- id: bigint (PK)
- currency_id_from: bigint (index)
- currency_id_to: bigint (index)
- rate: numeric
- created_at: timestamp
- updated_at: timestamp
- deleted_at: timestamp (nullable)
