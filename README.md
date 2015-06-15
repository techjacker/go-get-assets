# go-get-assets

Downloads all public google drive file links contained in a JSON file. Use Google Drive as your CMS.

## Installation

`Shell
go get github.com/techjacker/go-get-assets
`


## Usage

`
go-get-assets <input-json> <download-dir> <rewrite-urls-stem> <output-json>
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