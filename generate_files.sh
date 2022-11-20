#!/bin/sh
mkdir "files"
for i in 1 2 3 4 5 6 7 8 9 10
do
   touch files/file_$i.txt
   openssl rand -base64 1000 > files/file_$i.txt
done