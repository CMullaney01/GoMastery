package parsers

//cool way t do polymorphism in go, We embed the default parser and then when we want to change anything we can just write our own version of the function
type FoodHTMLParser struct {
	DefaultHTMLParser
}

func (p *FoodHTMLParser) ParseBody(content string) {}

func NewFoodParser() HTMLParser {
	return &FoodHTMLParser{}
}
