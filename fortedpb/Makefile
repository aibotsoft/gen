gen:
	protoc -I C:/Users/admin/go/ -I=.  --gogo_out=plugins=grpc,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,paths=source_relative:. *.proto

gen_py:
	python -m grpc_tools.protoc -I../../protos --python_out=. --grpc_python_out=. ../../protos/helloworld.proto
