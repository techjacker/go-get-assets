https://drive.google.com/file/d/0ByPfUp1fLihSNm5SSjZoalhPQ3M/view?usp=sharing


I found the solution. Use: https://googledrive.com/host/ID

eg:

https://googledrive.com/host/0ByPfUp1fLihSNm5SSjZoalhPQ3M
-> need to follow redirects


wget

saves into file: 0ByPfUp1fLihSNm5SSjZoalhPQ3M



`Go
    // Create the directories in the path
    file := filepath.Join(d.dir, key)
    if err := os.MkdirAll(filepath.Dir(file), 0775); err != nil {
        panic(err)
    }
`