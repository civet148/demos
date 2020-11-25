package main

import (
	"database/sql"
	"fmt"
	"github.com/elliotchance/sshtunnel"
	_ "github.com/go-sql-driver/mysql" //mysql golang driver
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"time"
)

func main() {

	// Setup the tunnel, but do not yet start it yet.
	tunnel := sshtunnel.NewSSHTunnel(
		// User and host of tunnel server, it will default to port 22
		// if not specified.
		"ubuntu@192.168.124.162",

		// Pick ONE of the following authentication methods:
		// sshtunnel.PrivateKeyFile("path/to/private/key.pem"), // 1. private key
		ssh.Password("lilobin148"), // 2. password

		// The destination host and port of the actual server.
		"127.0.0.1:3306",

		// The local port you want to bind the remote port to.
		// Specifying "0" will lead to a random port.
		"6033",
	)

	// You can provide a logger for debugging, or remove this line to
	// make it silent.
	tunnel.Log = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds)

	// Start the server in the background. You will need to wait a
	// small amount of time for it to bind to the localhost port
	// before you can start sending connections.
	go tunnel.Start()
	time.Sleep(100 * time.Millisecond)

	// NewSSHTunnel will bind to a random port so that you can have
	// multiple SSH tunnels available. The port is available through:
	//   tunnel.Local.Port

	// You can use any normal Go code to connect to the destination server
	// through localhost. You may need to use 127.0.0.1 for some libraries.
	//
	// Here is an example of connecting to a PostgreSQL server:
	conn := fmt.Sprintf("root:123456@tcp(127.0.0.1:6033)/chatmgr?charset=utf8mb4")
	if db, err := sql.Open("mysql", conn); err != nil {
		tunnel.Log.Printf("connect to mysql error %s", err)
		return
	} else {
		var rows *sql.Rows
		if rows, err = db.Query("SHOW TABLES"); err != nil {
			tunnel.Log.Printf("SHOW TABLES query error %s", err.Error())
			return
		}
		defer rows.Close()
		for rows.Next() {
			var strTableName string
			rows.Scan(&strTableName)
			tunnel.Log.Printf(strTableName)
		}
	}

}
