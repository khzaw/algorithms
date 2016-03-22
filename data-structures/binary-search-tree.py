class Node:
    def __init__(self, value=None):
        self.left, self.right, self.val = None, None, value


INFINITY = float("infinity")
NEG_INFINITY = float("-infinity")


def is_binary_search_tree(tree, minimum=NEG_INFINITY, maximum=INFINITY):
    if tree is None:
        return True

    if not minimum <= tree.value <= maximum:
        return False

    return is_binary_search_tree(tree.left, minimum, tree.value) and \
           is_binary_search_tree(tree.right, tree.value, maximum)