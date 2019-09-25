package internal

import (
	"fmt"
	"grpc-tutorial/api"
	"time"
)

func (w *FH) ListFiles() {
	api.ListFiles()

	go func() {
		for {
			w.Runtime.Events.Emit("cpu_usage", [4]int{1, 2, 3, 4})
			time.Sleep(1 * time.Second)
		}
	}()

	w.Runtime.Events.On("filesDropped", func(data ...interface{}) {
		// You should probably do better error checking
		fmt.Printf("I received the 'filesDropped' event with the message '%s'!\n", data[0])
	})
}

// func (w *FH) ListFiles() {
// 	api.ListFiles()

// 	go func() {
// 		for {
// 			w.Runtime.Events.Emit("cpu_usage", [4]int{1, 2, 3, 4})
// 			time.Sleep(1 * time.Second)
// 		}
// 	}()

// 	w.Runtime.Events.On("filesDropped", func(data ...interface{}) {
// 		// You should probably do better error checking
// 		fmt.Printf("I received the 'filesDropped' event with the message '%s'!\n", data[0])
// 	})
// }
