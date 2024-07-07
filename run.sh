#!/usr/bin/env bash

rm -rf kotlin_2024 
rm -rf python_processing_2024
rm -rf rrr_2024
rm -rf node/async_error_node_2024
rm -rf java/java_algos

git clone https://github.com/Thoughtscript/kotlin_2024.git &
git clone https://github.com/Thoughtscript/python_processing_2024.git &
git clone https://github.com/Thoughtscript/rrr_2024.git &
cd java && git clone https://github.com/Thoughtscript/java_algos.git &
cd node && git clone https://github.com/Thoughtscript/async_error_node_2024.git &

sleep 30 && docker-compose up

wait