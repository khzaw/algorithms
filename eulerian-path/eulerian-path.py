from collections import defaultdict

class Graph:
    def __init__(self):
        self._graph = defaultdict(list)
        self._odd_vertices = []

    def add_vertex(self, v):
        if not v in self._graph:
            self._graph[v] = list()

    def add_edge(self, v1, v2):
        self._graph[v1].append(v2)
        self._check_odd_vertex(v1)

        self._graph[v2].append(v1)
        self._check_odd_vertex(v2)

    def _check_odd_vertex(self, v):
        return len(self._graph[v]) % 2

    def check_eulerian(self):
        odd_deg = {v:e for (e, v) in self._graph.iteritems() if len(v) % 2 != 0}
        return odd_deg or False

    def check_semi_eulerian(self):
        odd_deg = {v:e for (e, v) in self._graph.iteritems() if len(v) % 2 != 0}
        return odd_deg in [0, 2] or False



