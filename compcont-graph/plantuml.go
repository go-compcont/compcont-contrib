package compcontgraph

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const ENCODING = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"
const PLANTUML_URL = "http://www.plantuml.com/plantuml/png"

func plantUMLBase64(input []byte) []byte {
	encoding := base64.NewEncoding(ENCODING)
	return []byte(encoding.EncodeToString(input))
}

func deflateCompress(content []byte) ([]byte, error) {
	var comp bytes.Buffer
	w, _ := flate.NewWriter(&comp, flate.HuffmanOnly)
	_, _ = w.Write(content)
	_ = w.Flush()
	_ = w.Close()
	return comp.Bytes(), nil
}

func DeflateEncodedURL(content []byte) string {
	comp, err := deflateCompress(content)
	if err != nil {
		log.Println()
	}
	encoded := plantUMLBase64(comp)
	return fmt.Sprintf("%s/%s", PLANTUML_URL, string(encoded))
}

func renderToPNG(puml string, file string) (err error) {
	resp, err := http.Get(DeflateEncodedURL([]byte(puml)))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bs := &bytes.Buffer{}
	if _, err = io.Copy(bs, resp.Body); err != nil {
		return
	}
	err = os.WriteFile(file, bs.Bytes(), 0644)
	return
}
