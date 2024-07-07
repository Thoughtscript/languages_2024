#!/usr/bin/env bash

java -jar "review/target/JavaReviewTopics-1.0.0.jar" &

java "java_algos/src/main/java/io/thoughtscript/Main.java" &

wait