
docker build . -t simpleweb_go

#push on server
docker tag simpleweb_go khajer/simplewebgo:1
docker push khajer/simplewebgo:1
