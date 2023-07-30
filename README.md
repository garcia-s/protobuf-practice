# Creating protobufs in different languanges


## For every language you would need to instal the protoc plugin

### in Node you would need to install it with npm

`npm install -g protoc-gen-js`

## in GO you can install it directly from apt in debian based linux distros

`sudo apt install protoc-gen-go`

## in Dart you have to install it with pub

`dart pub global activate protoc_plugin`

you also have to update your path to add pub-cache/bin directory should look something like this in your .bashrc file

`export PUB_DIR = ~/.pub-cache/bin 
export PATH=$PATH:$PUB_DIR`
