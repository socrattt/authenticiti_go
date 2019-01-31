package main

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
    "encoding/xml"
)

type order struct {
    Id        string  `json:"id" xml:"id,omitempty"`
    Data      string  `json:"data" xml:"data"`
    CreatedAt string  `json:"createdAt" xml:"createdAt,omitempty"`
    UpdatedAt string  `json:"updatedAt" xml:"updatedAt,omitempty"`
}

func readBody(w http.ResponseWriter, r *http.Request) (b []byte) {
    b, err := ioutil.ReadAll(r.Body)

    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    return
}

func handleErr(err error) {
    if err != nil {
        panic(err)
    }
}

func convertXmlToJson(xml_input []byte) (js []byte) {
    var bufferOrder order
    xml.Unmarshal(xml_input, &bufferOrder)

    js, err := json.Marshal(bufferOrder)
    handleErr(err)
    return
}

func buildResponse(w http.ResponseWriter, js []byte) {
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}
