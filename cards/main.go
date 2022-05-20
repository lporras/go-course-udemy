package main

func main() {
  // var card string = "Ace of Spades"
  // card := "Ace of Spades"
  // card = "Five of Diamonds"
  //card := newCard()
  //fmt.Println(card)

  //cards := deck{"Ace of Diamonds", newCard()}
  //cards := newDeck()
  //fmt.Println(cards)
  //cards = append(cards, "Six of Spades")
  //hand, remainingDeck := deal(cards, 5)

  //fmt.Println("hand:")
  //hand.print()
  //fmt.Println("remaining deck:")
  //remainingDeck.print()
  // cards.saveToFile("my_cards")

  cards := newDeckFromFile("my_cards")
  cards.shuffle()
  cards.print()
}