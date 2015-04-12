package main

import (
	"os"
	// "io"
	"encoding/json"

	// "log"
)

type Renamer struct {
	in     string
	out    string
	rel    string
	needle string
}

// func NewDecoder(r io.Reader) *Decoder {
// func NewDecoder(r io.Writer) *Decoder {

// func (r Renamer) Run(results []Res, w io.Writer) error {
// func (r Renamer) Run(results []Res) error {
func (r Renamer) Run() error {

	// change to file reader & file writer
	// dec := json.NewDecoder(r)
	// enc := json.NewEncoder(w)

	var err error

	// inFile, err := ioutil.ReadFile(string(r.in))

	// const inFile = `
	// {
	//   "nophotos": "here",
	//   "mapofphotos": {
	//     "photourl": "http://gdrive.com/diff.jpg",
	//     "nastyarray": ["sdfsd"],
	//     "noaphoto": "sdfsd"
	//   },
	//   "clients": [{
	//     "name": "Samsung",
	//     "photourl": "http://gdrive.com/city.jpg"
	//   }, {
	//     "name": "Sony",
	//     "photourl": "http://gdrive.com/sf.jpg"
	//   }],
	//   "projects": [{
	//     "descriptionsub": "Launch Project",
	//     "descriptionsublink": "http://themummyofmulberryavenue.com",
	//     "photourl": [{
	//       "deep": "http://gdrive.com/highway.jpg, http://gdrive.com/freedom.jpg"
	//     }]
	//   }, {
	//     "descriptionsub": "Launch Project",
	//     "descriptionsublink": "http://themummyofmulberryavenue.com",
	//     "photourl": "http://gdrive.com/freedom.jpg, http://gdrive.com/traffic.jpg"
	//   }]
	// }
	// `

	outFile, err := os.OpenFile(string(r.out), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)

	// err = json.NewEncoder(outFile).Encode(inFile)
	err = json.NewEncoder(outFile).Encode(string(r.in))
	err = outFile.Close()
	// err = inFile.Close()

	return err
	// dec := json.NewDecoder(strings.NewReader(jsonStream))

	// _, err := w.Write([]byte("hello"))

	// ioutil.WriteFile

	// for {
	// 	var v map[string]interface{}
	// 	if err := dec.Decode(&v); err != nil {
	// 		log.Println(err)
	// 		return err
	// 	}
	// 	for k := range v {
	// 		if k != "Name" {
	// 			delete(v, k)
	// 		}
	// 	}
	// 	if err := enc.Encode(&v); err != nil {
	// 		log.Println(err)
	// 		return err
	// 	}
	// }
}

// func (d *Downloader) RewriteUrlsInJson(data *struct{}) *struct{} {
// 	for _, v := range d.Results {
// 		fmt.Printf("%q", v)
// 	}
// 	return data
// }
