#!/bin/bash

ip=`ipconfig getifaddr en0`
cat visualizingEvidence.html | sed "s/localhost/$ip/" > index.html

python -m SimpleHTTPServer 8000
