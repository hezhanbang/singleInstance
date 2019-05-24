#!/bin/bash

option=$1
case "$option" in
	dst=*) value=`echo "$option" | sed -e 's/[-_a-zA-Z0-9]*=//'` ;;
	DST=*) value=`echo "$option" | sed -e 's/[-_a-zA-Z0-9]*=//'` ;;
	   *) value="" ;;
esac
echo $value

dstDir=
if [ -n "$value" ]; then
	dstDir=$value
fi

workDir=$(pwd)
singleDir=$(dirname $(readlink -f $0))
cd $singleDir

if [ -n "$dstDir" ]; then
	rm -rf $dstDir
	mkdir -p $dstDir
	cp lib/singleLinux.go $dstDir/single.go
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
