#!/bin/bash

## e.g. sh preBuild.sh --dst=/home/heb/install

opt=

for option
do
    opt="$opt `echo $option | sed -e \"s/\(--[^=]*=\)\(.* .*\)/\1'\2'/\"`"

    case "$option" in
        -*=*) value=`echo "$option" | sed -e 's/[-_a-zA-Z0-9]*=//'` ;;
           *) value="" ;;
    esac

    case "$option" in
        --dst=*)                       DST_DIR="$value"             ;;
	esac
done

workDir=$(pwd)
singleDir=$(dirname $(readlink -f $0))
cd $singleDir

if [ -n "$DST_DIR" ]; then
	rm -rf $DST_DIR
	mkdir -p $DST_DIR
	cp lib/singleLinux.go $DST_DIR/single.go
else
	rm -rf single.go
	cp lib/singleLinux.go single.go
fi

retVal=$?
cd $workDir

if [ $retVal -ne 0 ]
then
    echo "fail to do preBuild.sh in lib [singleInstance]!"
    exit 2
fi

exit 0
