#/bin/bash
docker volume create --name $2
docker run --rm -it -v $1:/from -v $2:/to alpine ash -c "cd /from ; cp -av . /to"
[ $? -eq 0 ] && docker volume rm $1
