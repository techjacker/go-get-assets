package main

import (
// "encoding/json"
// "log"
// "os"
)

type Renamer struct {
	in     string
	out    string
	rel    string
	needle string
}

func (r Renamer) Run(results []Res) error {
	// change to file reader & file writer
	// ioutil.WriteFile
	// dec := json.NewDecoder(os.Stdin)
	// enc := json.NewEncoder(os.Stdout)
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
	return nil
}

// func (d *Downloader) RewriteUrlsInJson(data *struct{}) *struct{} {
// 	for _, v := range d.Results {
// 		fmt.Printf("%q", v)
// 	}
// 	return data
// }
