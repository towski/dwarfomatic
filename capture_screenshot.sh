#!/bin/bash

#echo "import -window $1 /home/towski/code/dwarfomatic/public/$2;" >> ./screenshot_log.log

DIR=/home/towski/code/dwarfomatic

import -window $1 $DIR/public/$2;

THUMBNAIL=$(echo $2 | sed s/\.png//)
convert $DIR/public/$2 -thumbnail 96x96 $DIR/public/$THUMBNAIL\_thumb.png
