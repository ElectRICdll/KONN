package basic

const BLOCK_SIZE = 9

type Node struct {
	id         int
	contains   []Substance
	curTerrain Terrain
	curWeather Weather
	position   struct {
		Y int
		X int
	}
}

func NewNode(curTerrain Terrain) Node {
	return Node{
		curTerrain: curTerrain,
		curWeather: NewCommonWeather(),
	}
}

func (n *Node) ShowContains() []Substance {
	return n.contains
}

func (n *Node) Enter(sub *Substance) {

}

func (n *Node) Exit(sub *Substance) {

}

func (n *Node) Terrain() Terrain {
	return n.curTerrain
}

func (n *Node) SetTerrain(t Terrain) {
	n.curTerrain = t
}

func (n *Node) Weather() Weather {
	return n.curWeather
}

func (n *Node) SetWeather(w Weather) {
	n.curWeather = w
}

func (n *Node) PosY() int {
	return n.position.Y
}

func (n *Node) PosX() int {
	return n.position.X
}
