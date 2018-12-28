package main

import "time"
import "strconv"
// import "log"
import "github.com/gin-gonic/gin"
import "github.com/olivere/elastic"
// import "github.com/teris-io/shortid"
import "net/http"
import "./collectors"
//import "fmt"
//import "context"


const (
    elasticIndexName    = "assets"
    elasticTypeName     = "document"
)

var (
    elasticClient *elastic.Client
)

func main() {
    var err error
    //ctx := context.Background()
    for {
        elasticClient, err = elastic.NewClient(
            elastic.SetURL("http://localhost:9200"),
            elastic.SetSniff(false),
        )
        if err != nil {
            time.Sleep(3 * time.Second)
        } else {
            break
        }
    }
}
