package engine

func getCurrentAttributeName(g *Game) string {
	return "Set your " + g.characterCreator.GetCurrentAttributeName() + " to"
}
