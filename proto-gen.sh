#/bin/sh

# remove generated files
rm -rf gen && mkdir gen

for protoFile in $(find proto/github.com -name *.proto)
do
  protoc -I proto --go_out=gen "$protoFile"
done
