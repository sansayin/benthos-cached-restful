// Debug remote go
// dlv debug --headless --listen :4747 . -- run
// dlv connect :4747
package main

import (
    "context"

	"github.com/benthosdev/benthos/v4/public/service"

	// Import all standard Benthos components
	_ "github.com/benthosdev/benthos/v4/public/components/all"

	"fmt"
	"runtime"
	"time"

)

func init() {

	go func() {
		tk := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-tk.C:
				fmt.Printf("Routines: %v\n", runtime.NumGoroutine())
			}
		}
	}()
}
func main() {
	service.RunCLI(context.Background())
}
