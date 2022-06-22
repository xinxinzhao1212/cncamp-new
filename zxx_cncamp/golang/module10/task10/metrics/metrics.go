package metrics

import "fmt"

func Registry() {

	err := prometheus.Register(functionLatency)

	if err != nil {

		fmt.Println(err)

	}
}
