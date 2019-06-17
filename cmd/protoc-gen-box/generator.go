package main

import (
	"bytes"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

type (
	Generator struct {
		*bytes.Buffer
		Request  *plugin.CodeGeneratorRequest
		Response *plugin.CodeGeneratorResponse
	}
)

func New() *Generator {
	g := new(Generator)
	g.Buffer = new(bytes.Buffer)
	g.Request = new(plugin.CodeGeneratorRequest)
	g.Response = new(plugin.CodeGeneratorResponse)

	return g
}

// Error reports a problem, including an error, and exits the program.
func (g *Generator) Error(err error, msgs ...string) {
	s := strings.Join(msgs, " ") + ":" + err.Error()
	log.Print("protoc-gen-go: error:", s)
	os.Exit(1)
}

// Fail reports a problem and exits the program.
func (g *Generator) Fail(msgs ...string) {
	s := strings.Join(msgs, " ")
	log.Print("protoc-gen-go: error:", s)
	os.Exit(1)
}

func (g *Generator) GenerateAllFiles() {
	log.Println("GetParameter", g.Request.GetParameter())

	for _, file := range g.Request.GetFileToGenerate() {
		log.Println("GetFileToGenerate", file)
	}

	for _, file := range g.Request.GetProtoFile() {
		log.Println(file.GetName())

		if file.GetName() != "api/api.proto" {
			continue
		}

		g.Response.File = append(g.Response.File, &plugin.CodeGeneratorResponse_File{
			Name:    proto.String("api.pb.gin.go"),
			Content: proto.String("aaa"),
		})

		for _, srv := range file.GetService() {
			for _, method := range srv.GetMethod() {
				log.Println(method.GetName(), method.GetOutputType(), method.GetInputType())
			}
		}
	}

	// data, _ := json.Marshal(g.Response)

	// fmt.Print(string(data[:]))
}
