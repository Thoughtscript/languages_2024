#!/usr/bin/env bash

rm -rf kotlin_2024 
rm -rf python_processing_2024
rm -rf rrr_2024
rm -rf node/async_error_node_2024
rm -rf java/java_algos
rm -rf php_2024

git clone https://github.com/Thoughtscript/kotlin_2024.git &
git clone https://github.com/Thoughtscript/python_processing_2024.git &
git clone https://github.com/Thoughtscript/rrr_2024.git &
git clone https://github.com/Thoughtscript/php_2024.git &
cd java && git clone https://github.com/Thoughtscript/java_algos.git &
cd node && git clone https://github.com/Thoughtscript/async_error_node_2024.git &

# If using Docker Compose Engine V2 uncomment:
# sleep 30 && docker compose up
sleep 30 && docker-compose up

wait