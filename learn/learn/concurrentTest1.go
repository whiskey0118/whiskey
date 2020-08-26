//go程序设计语言8.5 并行循环
package main

import (
	"gopl.io/ch8/thumbnail"
	"log"
)

func main() {
	makeThumbnails([]string{"F:\\我的\\go\\whiskeybee\\whiskeybee\\websockTest\\learn\\1.jpg", "F:\\我的\\go\\whiskeybee\\whiskeybee\\websockTest\\learn\\2.jpg"})
}

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f)
	}
}

func makeThumbnails2(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}
	for range filenames {
		<-ch
	}
}

func test(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

//func test1()  {
//	a,err :=ioutil.ReadFile("F:\\我的\\go\\whiskeybee\\whiskeybee\\websockTest\\learn\\1.png")
//	if err != nil{
//		fmt.Println(err)
//	}
//	fmt.Println("a:",a)
//
//}
