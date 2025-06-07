# pyright: basic

from node import Node


def depth_first_search(root, search_value):
    stack = [root]
    root.visited = True

    while len(stack) > 0:
        node = stack.pop()
        for neighbour in node.adjacentNodes:
            if not neighbour.visited:
                stack.append(neighbour)

        if node.value == search_value:
            return True
        else:
            node.visited = True

    return False


if __name__ == "__main__":
    node0 = Node(0)
    node1 = Node(1)
    node2 = Node(2)
    node3 = Node(3)
    node4 = Node(4)
    node5 = Node(5)
    node6 = Node(6)

    node0.adjacentNodes = [node1, node2, node3]
    node1.adjacentNodes = [node5, node6]
    node2.adjacentNodes = [node4]
    node3.adjacentNodes = [node2, node4]
    node4.adjacentNodes = [node1]
    node5.adjacentNodes = []
    node6.adjacentNodes = [node4]

    assert depth_first_search(node0, 4) == True
