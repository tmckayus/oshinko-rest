#/bin/bash
# Create a cluster named $2 with $1 workers using the oshinko-rest at $3
#curl -H "Content-Type: application/json" -X POST -d '{"name": "'$2'", "config": {"workerCount": '$1', "masterCount": 1, "metrics": {"enable": "true", "carbon": "docker.io/elmiko/carbon", "graphite": "172.30.100.159:5000/tmckay-test/tmckay-graphite"}, "sparkmasterconfig": "metricsconfig"}}' $3/clusters

curl -H "Content-Type: application/json" -X POST -d '{"name": "'$2'", "config": {"name": "clusterconfig"}}' $3/clusters
