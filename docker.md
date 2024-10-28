## To get the ip adress
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' name_of_db 
