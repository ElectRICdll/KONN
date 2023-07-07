package basic

import "konn/entity/prop"

type Builder struct {
	list map[string]*prop.Structure
}

func (b *Builder) Build() {

}

func (b *Builder) Cancel() {

}

func (b *Builder) List() map[string]*prop.Structure {
	return b.list
}
