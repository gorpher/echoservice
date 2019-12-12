docker build -t gorpher/echoservice:v1 docker
echo "docker build success!!!"

docker push gorpher/echoservice:v1
echo "docker push success!!!"


rm -f docker/echoservice-linux-amd64
rm -f docker/healthchecker-linux-amd64

echo "clean success!!!"

