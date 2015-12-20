## Master Tokens

master_token POST returns this:

```json
{
	"id":1111,
	"repository_id":1111,
	"name":"name",
	"value":"value",
	"created_at":"2015-12-20T21:12:09.781Z",
	"updated_at":"2015-12-20T21:12:09.781Z"
}
```

not 

```json
{
    "value" : "value",
    "name" : "name",
    "paths" : {
        "self" : "/api/v1/repos/username/reponame/master_tokens/1"
    }
}
```

In addition example curl does not work as `Content-Type` needs to be set to `application/json` and the `DELETE` method is not documented for deleting tokens?

## PackageDetail

JSON is missing s `"` at the end of the line `"repository_url": "/api/v1/repos/julio/testrepo,`

