package main

import(
    "bufio"
    "encoding/csv"
    "fmt"
    "os"
)

func main(){
    //open file 
    file, _ := os.open("text.csv")
    
    // create csv reader 
    rdr := csv.NewReader(bufio.NewReader(file))
    
    // read csv contents
    rows, _ := rdr.ReadAll()
    
    // read row, col
    for i, row := range rows {
        for j := range row{
            fmt.Printf("%s", rows[i][j])
        }
        fmt.println()
    }
}