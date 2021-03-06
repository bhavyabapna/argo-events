package main

import (
	"io/ioutil"

	"sigs.k8s.io/yaml"
)

func cleanCRD(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	crd := make(map[string]interface{})
	err = yaml.Unmarshal(data, &crd)
	if err != nil {
		panic(err)
	}
	delete(crd, "status")
	metadata := crd["metadata"].(map[string]interface{})
	delete(metadata, "annotations")
	delete(metadata, "creationTimestamp")
	spec := crd["spec"].(map[string]interface{})
	delete(spec, "validation")
	for _, v := range spec["versions"].([]interface{}) {
		v1 := v.(map[string]interface{})
		delete(v1, "schema")
	}
	data, err = yaml.Marshal(crd)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		panic(err)
	}
}
