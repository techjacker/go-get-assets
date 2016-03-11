# go-get-assets

Download the google drive assets contained in a JSON file. Use Google Drive as your CMS.


----------------------------------

## Installation

```Shell
go get github.com/techjacker/go-get-assets
```

----------------------------------

## Usage

### CLI

```Shell
$ go-get-assets <input-json> <download-dir>
```

#### CLI Options
```
$ go-get-assets -r <rewrite-urls-stem> -o <output-JSON-path> -n <needle> <input-json> <download-dir>
```
stem to replace downloaded URLs in rewritten JSON file
```
-r <rewrite-urls-stem>
```

path to write rewritten JSON file
```
-o <output-JSON-path>
```

needle to use to look for download URLs (defaults to Google drive stem)
```
-n <needle>
```
show help
```
$ go-get-assets -h
```


----------------------------------

## Examples

### Example 1

Just downloads the assets from Google Drive.

```Shell
$ go-get-assets $PWD/cms.json $PWD/images
```

### Example 2

Downloads the assets from Google Drive and saves new JSON file with paths to assets rewritten. 

```Shell
$ go-get-assets -r /images -o $PWD/cms.output.json $PWD/cms.json $PWD/images
```

#### BEFORE
```
$ cat cms.json
{
    "clients": [{
        "name": "Samsung",
        "photourl": "https://drive.google.com/file/d/city.jpg"
    }],
    "photourl": "https://drive.google.com/file/d/freedom.jpg, https://drive.google.com/file/d/traffic.jpg"
}
```

#### AFTER
```
$ ls -l images
 images/city.jpg
 images/freedom.jpg
 images/traffic.jpg
```

```
$ cat cms.output.json
{
    "clients": [{
        "name": "Samsung",
        "photourl": "/images/city.jpg"
    }],
    "photourl": "/images/freedom.jpg, /images/traffic.jpg"
}
```


----------------------------------

## How it works

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
