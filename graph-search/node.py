# pyright: basic


class Node:
    def __init__(self, value):
        self.value = value
        self.adjacentNodes = []
        self.visited = False

    def __str__(self):
        return str(self.value)

    def __repr__(self):
        return str(self.value) + " " + str(self.visited)

    def __eq__(self, other):
        return self.value == other.value
