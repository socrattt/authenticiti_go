package main

import (
    "net/http"
    "bytes"
    "io/ioutil"
    "fmt"
)

const URL = "http://52.52.253.24:8888"

func postXmlServer(xml_input []byte) (b []byte) {
    resp, err := http.Post("http://52.52.253.24:8888/orders", "application/xml", bytes.NewBuffer(xml_input))
    handleErr(err)

    defer resp.Body.Close()
    b, err = ioutil.ReadAll(resp.Body)
    return
}

func getXmlServer(order_id []byte) (b []byte) {
    resp, err := http.Get(fmt.Sprintf("%s/orders/%s", URL, order_id))
    handleErr(err)

    defer resp.Body.Close()
    b, err = ioutil.ReadAll(resp.Body)
    return
}
