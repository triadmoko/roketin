# Roketin

## POST localhost:8080/api/v1/create
**Body**

![](doc/create.png)

**Response** 
```json
{
    "message": "Insert Success",
    "code": 200,
    "status": "success",
    "data": {
        "title": "asdfasdf",
        "description": "",
        "duration": "sdfasfg",
        "artist": "asdfasdf",
        "genre": "dfhsdfdg",
        "filename": "video/1f1d26b34700053b7518.mp4"
    }
}
```

## PUT localhost:8080/api/v1/update
![](doc/update.png)

```json
{
    "message": "Insert Success",
    "code": 200,
    "status": "success",
    "data": {
        "title": "Good Action",
        "description": "Long time",
        "duration": "03:20:20",
        "artist": "My Sister",
        "genre": "Comedy",
        "filename": "video/0f6fdfc7781ebbc45693.mp4"
    }
}
```