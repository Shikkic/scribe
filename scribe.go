package main

import (
	"log"
	"net/http"
)

var PORT = ":4000"
var RootNode *Node

func main() {
	http.HandleFunc("/set", SetHandler)
	http.HandleFunc("/unset", RemoveHandler)
	http.HandleFunc("/get", GetHandler)

	log.Printf("Listing on port %s", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func SetHandler(w http.ResponseWriter, r *http.Request) {
	setKey := r.FormValue("key")
	setValue := r.FormValue("value")
	log.Printf("Recieved Key: %s, Value: %s", setKey, setValue)

	if setKey == "" {
		http.Error(w, "No key provided", http.StatusBadRequest)
		return
	}

	if setValue == "" {
		http.Error(w, "No value provided", http.StatusBadRequest)
		return
	}

	newNode := Node{setKey, setValue, nil, nil}
	AppendNodeToRootNode(&newNode)
	PrintRootTree()
}

func RemoveHandler(w http.ResponseWriter, r *http.Request) {
	removeKey := r.FormValue("key")

	if removeKey == "" {
		http.Error(w, "No key provided", http.StatusBadRequest)
	}

	log.Printf("Recieved remove Key: %s", removeKey)
	error := RemoveKeyFromRootNode(removeKey)

	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	PrintRootTree()
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	getKey := r.FormValue("key")

	if getKey == "" {
		http.Error(w, "No key provided", http.StatusBadRequest)
		return
	}

	nodeValue, err := GetValueFromRootNodeKey(getKey)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Write([]byte(nodeValue))
}
