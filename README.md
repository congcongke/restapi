# restapi

```
swagger generate server -f swagger/swagger.yaml --template-dir=swagger/templates --exclude-main --additional-initialism=CVE --additional-initialism=GC -A harbor --existing-models=github.com/goharbor/harbor/src/server/v2.0/models --server-package=.

rm -rf models
```