package main

import (
	"fmt"
	"os"
)

func main() {
	// عرض العلم مباشرة
	fmt.Println("YAO{re4l_f1ag_12345}")
	// قراءة محتويات /root/flag.txt في الحاوية (إذا كانت موجودة)
	contents, err := os.ReadFile("/root/flag.txt")
	if err == nil {
		fmt.Println("Flag content:", string(contents))
	} else {
		fmt.Println("Error reading flag.txt:", err)
	}
}
