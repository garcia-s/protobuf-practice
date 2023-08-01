#! /bin/bash
protoc --proto_path=./proto_models  --go_out=./client/  --go_out=./go --js_out=./node --dart_out=./dartlang user.proto

