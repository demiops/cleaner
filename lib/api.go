# API
# Central brain of cleaner
# Takes care of storing rules/policies in the sql databse 
# Takes care of instructing executor which objects needs to delete
# Takes care of instructing notifier to send notifications to heraldo on objects scheduled for deletion
# Takes care of keeping track of which objects are blacklisted and disabled for deletion


import "database/sql"
import _ "github.com/go-sql-driver/mysql"

db, err := sql.Open("mysql", "user:password@/dbname")

if err != nil {
    panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
}
defer db.Close()

// Open doesn't open a connection. Validate DSN data:
err = db.Ping()
if err != nil {
    panic(err.Error()) // proper error handling instead of panic in your app
}

// Prepare statement for inserting data
    stmtIns, err := db.Prepare("INSERT INTO squareNum VALUES( ?, ? )") // ? = placeholder
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

    // Prepare statement for reading data
    stmtOut, err := db.Prepare("SELECT squareNumber FROM squarenum WHERE number = ?")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtOut.Close()

    // Insert square numbers for 0-24 in the database
    for i := 0; i < 25; i++ {
        _, err = stmtIns.Exec(i, (i * i)) // Insert tuples (i, i^2)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
    }

    var squareNum int // we "scan" the result in here

    // Query the square-number of 13
    err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    fmt.Printf("The square number of 13 is: %d", squareNum)

    // Query another number.. 1 maybe?
    err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    fmt.Printf("The square number of 1 is: %d", squareNum)
