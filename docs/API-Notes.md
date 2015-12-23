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

## Repositories

show return this:

```
{
  "name": "",
  "path": "/user/name",
  "repo_type": null,
  "created_at": "2015-12-17T11:41:28.000Z",
  "updated_at": "2015-12-17T11:41:28.000Z",
  "paths": {
    "create_package": "/api/v1/repos/user/name/packages.json",
    "package_contents": "/api/v1/repos/user/name/packages/contents",
    "master_tokens": "/api/v1/repos/user/name/master_tokens",
    "create_master_token": "/api/v1/repos/user/name/master_tokens",
    "self": "/api/v1/repos/user/name"
  },
  "urls": {
    "install_script": "https://token:@packagecloud.io/install/repositories/user/name/script.:package_type.sh"
  }
}
```

not the documented:

```
{
  "name": "myrepo",
  "created_at": "2014-02-26T00:03:28.000Z",
  "url": "https://packagecloud.io/cooluser/myrepo",
  "last_push_human": "3 days ago",
  "package_count_human": "17 packages",
  "private": false,
  "fqname": "cooluser/myrepo"
}
```

## PackageDetail

JSON is missing s `"` at the end of the line `"repository_url": "/api/v1/repos/julio/testrepo,`
