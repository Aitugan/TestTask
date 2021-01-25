# Documentation

## **_Controllers_**

### **Get Exchange Rate**

### Description

Users can get all possible courses. Depending on the request, handler returns the currency rate for a specific pair

#### Sample Request

```
GET /api/exchange-rate?base=USD&quote=RUB
```

#### Sample Response

```
{
  "data": "74.55RUB",
  "message": "Success",
  "status": true
}
```

### **Get Rate History**

### Description

Users can see history of changes in any currency pair

#### Sample Request

```
GET /api/history?base=USD&quote=RUB
```

#### Sample Response

```
{
  "data": [
    {
      "Date": "25.01.2021",
      "Value": "74.55"
    }
  ],
  "message": "Success",
  "status": true
}

```

### **Make conversion**

### Description

Users can convert for any currency pair

#### Sample Request

```
GET /api/history?base=USD&quote=RUB
```

Body:

```
{
"convertFrom":"KZT",
"amount":1000,
"convertTo":"USD"
}
```

#### Sample Response

```
{
  "data": "2.38",
  "message": "Success",
  "status": true
}

```
