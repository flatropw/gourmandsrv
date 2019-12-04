# Gourmand Server
Web server handling any request to it’s address and returning JSON like:
```
{
    "Host": "127.0.0.1:8081",
    "UserAgent": "curl/7.67.0",
    "RequestUri": "/any/path?foo=bar",
    "Headers": {
        "Accept": [
            "text/html"
        ],
        "User-Agent": [
            "curl/7.67.0""
        ]
    },
    "QueryParams": {
        "foo": [
            "bar"
        ]
    }
}
```
All was taken from request struct (Query params included).
