package errpkg

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)


var Answersheet *answersheet
var General *general
var Auth *auth
var Test *test

func LoadError() {
	root := rootErr{}

	file, err := os.ReadFile("error.yml")

	if err != nil {
		log.Fatalln("error load error", err)
	}

	err = yaml.Unmarshal(file, &root)

	if err != nil {
		log.Fatalln("error unmarshal file", err)
	}

	//General = root.General
	
    Answersheet = root.Answersheet 
    General = root.General 
    Auth = root.Auth 
    Test = root.Test 
}

func LoadErrorFromPath(path string) {
    root := rootErr{}
    if len(path) == 0 {
        path = "error.yml"
    }
    file,err := os.ReadFile(path)
    if err != nil {
        log.Fatalln("error load error",err)
    }
    err = yaml.Unmarshal(file,&root)
    if err != nil {
        log.Fatalln(err)
    }

    	//General = root.General
    	
        Answersheet = root.Answersheet 
        General = root.General 
        Auth = root.Auth 
        Test = root.Test 
}