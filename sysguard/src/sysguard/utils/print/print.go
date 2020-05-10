package print

import "fmt"

func Info(m string) {
	fmt.Println("Info      | ", m)
}

func Critical(m string) {
	fmt.Println("Critical  | ", m)
}

func Warning(m string) {
	fmt.Println("Warning   | ", m)
}

func Welcome() {
	fmt.Println(`          | SYSGUARD - PROXY - V DEV 0.0.1        
          | `)
}
