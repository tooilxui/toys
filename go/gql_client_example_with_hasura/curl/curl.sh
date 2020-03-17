# 1. query with json variable 不用引號
# curl -X POST http://localhost:8080/v1/graphql \
# -H "Content-Type: application/json" \
# --data '{ "query": "query ('$spec': jsonb) { erp_json_test(where: {data: {_contains: '$spec'}}) { data } }", "variables": { "spec": { "A": "aaa" } } }'

# 2. simple query
# curl -X POST \
# http://192.1.1.115:8080/v1/graphql \
# -H "Content-Type:application/json" \
# -d '{
# 	"query" : "query { erp_json_test { data }}"
# }'

# 3. query with file
# curl -X POST \
# http://192.1.1.115:8080/v1/graphql \
# -H "Content-Type:application/json" \
# --data-binary "@./curl.json"