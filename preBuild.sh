#!/bin/bash

workDir=$(pwd)
singleDir=$(dirname $(readlink -f $0))
cd $singleDir

rm -rf single.go
cp lib/singleLinux.go single.go

retVal=$?
cd $workDir

if [ $retVal -ne 0 ]
then
    echo "fail to do preBuild.sh in lib [singleInstance]!"
    exit 2
fi

exit 0