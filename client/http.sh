#!/bin/bash

ip=`ipconfig getifaddr en0`
cat experiment-b.html | sed "s/localhost/$ip/" > index.html

python -m SimpleHTTPServer 8000
