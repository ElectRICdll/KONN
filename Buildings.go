package main

type MainController struct {
	Constructor factory
	Producer
	health int
	BelongsTo string
}

func (self *MainController) setBody() {
	self.health = 1000
}

func (self *MainController) construct() {

}

func (self *MainController) produce() {

}
