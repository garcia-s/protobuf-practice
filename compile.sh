#! /bin/bash
protoc --proto_path=./proto --go_out=./go/proto/ --js_out=./node/proto/ --dart_out=./dartlang/proto user.proto

