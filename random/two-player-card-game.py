#  Here is the description of the game simulation.

    #  this is a two player card game
    #  the game starts with a deck of cards
    #  the cards are dealt out to both players
    #  on each turn:
        #  both players turn over their top-most card
        #  the player with the higher valued card takes the cards and puts them in their scoring pile (scoring 1 point per card)
    #  this continues until the players have no cards left
    #  the player with the highest score wins

# Extensions
    # extend the game to support more than two players

# Super bonus
#  Should the interviewee manage to add N players support, I still have one additional change. It’s a bit more challenging than the previous requirements: “don’t assume the cards have unique values”. The details are:

    #  Remove the assumption that the cards are unique (that is, go to a standard deck)
    #  If players have the same valued card, they draw an additional card, repeat until one has a higher card
    #  The person with the highest card takes all the cards, scoring one each
    #  Only the players that tie continue drawing cards (1+ players may sit out on the additional draws)

