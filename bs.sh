docker build ./ -t dining_hall_image
docker stop dining_hall_container
docker run -d --rm -p 8001:8001 --name dining_hall_container dining_hall_image go run main http://host.docker.internal
