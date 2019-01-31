package main

import (
    "net/http"
    "encoding/json"
    "encoding/xml"
    "github.com/gorilla/mux"
)

func postOrder(w http.ResponseWriter, r *http.Request) {
    // convert input to XML
    var inputOrder order
    json.Unmarshal(readBody(w, r), &inputOrder)
    xml_input, err := xml.Marshal(inputOrder)
    handleErr(err)

    // call out to server
    xml_response := postXmlServer(xml_input)
    js_response := convertXmlToJson(xml_response)

    buildResponse(w, js_response)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
    // call external server
    xml_response := getXmlServer([]byte(mux.Vars(r)["id"]))
    js_response := convertXmlToJson(xml_response)

    buildResponse(w, js_response)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/orders", postOrder).Methods("POST")
    r.HandleFunc("/orders/{id}", getOrder).Methods("GET")
    http.Handle("/", r)

    err := http.ListenAndServe(":9123", nil);
    handleErr(err)
}
