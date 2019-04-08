#!/bin/bash

workDir=$(pwd)
singleDir=$(dirname $(readlink -f $0))
cd $singleDir

rm -rf single.go
cp lib/singleLinux.go single.go

retVal=$?
if [ $retVal -ne 0 ]
then
    echo "fail to do preBuild.sh in lib [SingleInstance]!"
    exit 2
fi

exit 0