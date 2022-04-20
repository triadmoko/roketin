# roketin


## POST localhost:8080/api/v1/create
**Body**
```json
{
  "title": "trasdf",
  "description": "description",
  "duration": "asdasd",
  "artist": "sdf",
  "genre": "asdf",
  "filename": "asdfsdf"
}
```
**Response** 
```json
{
    "message": "Insert Success",
    "code": 200,
    "status": "success",
    "data": {
        "title": "trasdf",
        "description": "",
        "duration": "asdasd",
        "artist": "sdf",
        "genre": "asdf",
        "filename": "asdfsdf"
    }
}
```
