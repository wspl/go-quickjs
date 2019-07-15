#!/bin/sh
go_quickjs="github.com/wspl/go-quickjs"
install_dir="$GOPATH/src/$go_quickjs"
quickjs="https://bellard.org/quickjs/quickjs-2019-07-09.tar.xz"

old_pwd=$(pwd)

if test -f "quickjs.go"; then
  install_dir=$(pwd)
else
  go get -d $go_quickjs
  cd $install_dir
fi

rm -rf ./libquickjs.a
rm -rf ./quickjs-source

wget -O quickjs-source.tar.xz $quickjs
mkdir quickjs-source
tar zxvf quickjs-source.tar.xz -C quickjs-source

cd quickjs-source/quickjs*
sed -i -e '37d' ./Makefile
make
cp libquickjs.a $install_dir

cd $install_dir
go install

cd $old_pwd