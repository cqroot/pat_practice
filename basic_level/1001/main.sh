#!/bin/sh

read -r num
count=0
while [ "$num" -gt 1 ];
do
    if [ $((num%2)) = 1 ]; then
        num=$(((num*3+1)/2))
    else
        num=$((num/2))
    fi
    count=$((count+1))
done
echo $count
