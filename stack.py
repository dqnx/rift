"""stack

Implements a simple LIFO data structure.
"""

class Stack:
    def __init__(self):
        self._stack = []
    
    def empty(self) -> bool:
        """Returns if the Stack is empty."""
        return not self._stack
    
    def size(self) -> int:
        """Returns the number of items in the stack."""
        return len(self._stack)
    
    def top(self) -> object:
        """Returns the top object in the stack."""
        return self._stack[-1]
    
    def push(self, element: object):
        """Adds an item to the top of the stack."""
        self._stack.append(element)
    
    def pop(self) -> object:
        """Removes the top item from the stack and returns it."""
        return self._stack.pop()
        