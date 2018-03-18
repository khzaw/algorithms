# Optimize for the time cost of m calls on your queue.
class Queue:

    def __init__(self):
        self.stack1 = []
        self.stack2 = []

    def enqueue(self, item):
        self.stack1.append(item)

    def dequeue(self):
        if len(self.stack2) == 0:
            while self.stack1:
                self.stack2.append(self.stack1.pop())

        if len(self.stack2) == 0:
            raise IndexError('Empty queue. Insert something first.')

        return self.stack2.pop()
