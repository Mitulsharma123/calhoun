package deck

type Suit uint8

const (
	Spade Suit = iota
	Heart Suit
	Diamond Suit
	Club Suit
)

type Rank uint8

const(
	_ Rank = iota
	Ace 
	Two 
	Three
	Four
	Five
	Six
	Seven
	Eight 
	Nine 
	Ten
	Jack
	Queen
	King
)

type Card struct{
	Suit
	Rank
}

