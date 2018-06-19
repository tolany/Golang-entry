package main
import(
    "bufio"
    "enconding/csv"
    "os"
    )
    
    func main(){
        // create file
        file, err := os.create("./output.csv")
        if err != nil{
            panic(err)
        }
        
        // create csv writer 
        wr := csv.NewWriter(bufio.NewWriter(file))
        
        // write csv 
        wr.Write([]string{"A", "0.25"})
        wr.Write([]string{"B", "55.70"})
        wr.Flush
    }