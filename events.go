
package main

import (
	"./filemanager"
	"./menu"
)
  func enter(mn *menu.Menu) {
		activePath := mn.GetActive()
		fm.Path().Push(activePath)
		if filemanager.IsDir(fm.Path().Current()) {
			err := refresh(mn, &fm, fm.Path().Current())
			if err != nil {
				fm.Path().Pop()
				mn.ShowMsg(err.Error())
			}
		} else {
			c := make(chan string)
			worker := &Worker{Command: "xdg-open", Args: fm.Path().Current(), Output: c}
			go worker.Run()
			mn.ShowMsg(fm.Path().Current() + " is opening...")
			fm.Path().Pop()
		}
	}
func back(mn *menu.Menu) {
		curr := fm.Path().Current()
		fm.Path().Pop()
		err := refresh(mn, &fm, fm.Path().Current())
		if err != nil {
			fm.Path().Set(curr)
			mn.ShowMsg(curr + " is the first directory")
		}
	}

  func exit(m *menu.Menu){
    menu.Close()
  }
