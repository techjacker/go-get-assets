# go-get-assets

Downloads all public google drive file links contained in a JSON file. Use Google Drive as your CMS.

## Installation

`Shell
go get github.com/techjacker/go-get-assets
`


## Usage

### CLI

`Shell
$ go-get-assets <input-json> <download-dir>
`

#### CLI Options

stem to replace downloaded URLs in rewritten JSON file
`
-r <rewrite-urls-stem>
`

path to write rewritten JSON file
`
-o <output-JSON-path>
`

needle to use to look for download URLs (defaults to Google drive stem)
`
-n <needle>
`



<output-json>
#### Example

`Shell
$ go-get-assets cms.json $PWD/images /images cms.output.json
`
##### BEFORE
`
$ cat cms.json
{
    "clients": [{
        "name": "Samsung",
        "photourl": "https://drive.google.com/file/d/city.jpg"
    }],
    "photourl": "https://drive.google.com/file/d/freedom.jpg, https://drive.google.com/file/d/traffic.jpg"
}
`

##### AFTER
`
$ ls -l $PWD/images
images/city.jpg
images/freedom.jpg
images/traffic.jpg
`

`
$ cat cms.output.json
{
    "clients": [{
        "name": "Samsung",
        "photourl": "/images/city.jpg"
    }],
    "photourl": "/images/freedom.jpg, /images/traffic.jpg"
}
`


----------------------------------

### Methodology

Downloads public files from gdrive by constructing URLs based on the following logic.

Copy a google drive share link of a public file, eg:
`https://drive.google.com/file/d/<ID>`
Eg:
`https://drive.google.com/file/d/0ByPfUp1fLihSNm5SSjZoalhPQ3M/view?usp=sharing`

Extract the ID to construct the download URL.
`https://googledrive.com/host/<ID>`
Eg:
`https://googledrive.com/host/0ByPfUp1fLihSNm5SSjZoalhPQ3M`

WGET or similar the file. Make sure you enable following redirects.

The downloaded file is named after the ID, eg:
`/path/to/file/0ByPfUp1fLihSNm5SSjZoalhPQ3M`