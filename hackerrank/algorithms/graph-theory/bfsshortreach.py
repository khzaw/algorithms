#!/bin/python
from collections import deque

class Node:
    def __init__(self, value):
        self.value = value
        self.adjacentNodes = []
        self.visited = False

def breadth_first_traversal(start, graph):
    q = deque([start])
    start.visited = True

    while len(q) > 0:
        node = q.pop()
        for neighbour in node.adjacentNodes:
            if not neighbour.visited:
                q.appendleft(neighbour)
                graph[neighbour.value] = 6 + graph[node.value] if graph[node.value] > 0 else 6
                neighbour.visited = True

t = int(raw_input().strip())
for g in xrange(t):
    n, m = map(int, raw_input().strip().split(' '))
    graph = {}
    nodes = []
    for i in xrange(n):
        nodes.append(Node(i))
        graph[i] = -1

    for _ in range(m):
        m1, m2  = map(int, raw_input().strip().split(' '))
        nodes[m1-1].adjacentNodes.append(nodes[m2-1])
        nodes[m2-1].adjacentNodes.append(nodes[m1-1])
    s = int(raw_input().strip())
    breadth_first_traversal(nodes[s-1], graph)
    del graph[s-1]
    distance = ' '.join(str(x) for x in graph.values())
    print distance