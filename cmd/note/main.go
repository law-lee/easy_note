package main

import (
	demonote "github.com/law-lee/easy_note/kitex_gen/demonote/noteservice"
	"log"
)

func main() {
	svr := demonote.NewServer(new(NoteServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
