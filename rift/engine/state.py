"""state

Implements a state machine.
"""

from abc import ABC, abstractmethod
from enum import Enum
from engine.stack import Stack

class StateTransition(Enum):
    # Do nothing and continue.
    NONE = 0
    # Adds a new state and switches to it, pausing the previous state.
    ADD = 1
    # Adds a new state and swithces to it, popping the previous state.
    NEXT = 2
    # Pops the current state, returning to the previous state, if exists.
    POP = 3
    # Pops all states.
    EXIT = 4

class State(ABC):

    @abstractmethod
    def run(self, state_input):
        """Returns the transition and next state, if applicable"""
        pass

    @abstractmethod
    def __str__(self) -> str:
        pass

    def on_enter(self):
        """Runs when state is resumed or entered."""
        pass

    def on_exit(self):
        """Runs when state is exited or popped."""
        pass

class StateMachine(ABC):
    """State machine with a LIFO model of the states."""
    def __init__(self):
        # states is a stack here, LIFO. This allows states to be "paused."
        self.states = Stack()

    def run(self, state_input) -> bool:
        """Gets transition from current state.

        Closes state machine if no states remain or exit state transition.
        
        Returns boolean indicating state machine status (False=empty, True=running).
        """
        op, new_state = self.states.top().run(state_input)

        if op == StateTransition.NONE:
            return True

        elif op == StateTransition.ADD:
            # Clean up past state if it exists
            if not self.states.empty():
                self.states.top().on_exit()
                # Pop past state
                self.states.pop()
            
            # Add new state
            self.states.push(new_state)
            self.states.top().on_enter()

        elif op == StateTransition.NEXT:
            # Clean up past state if it exists
            if not self.states.empty():
                self.states.top().on_exit()
            
            # Add new state
            self.states.push(new_state)
            self.states.top().on_enter()

        elif op == StateTransition.POP:
            # Clean up past state
            self.states.top().on_exit()
            self.states.pop()

            # Enter new state, if exists
            if not self.states.empty():
                self.states.top().on_enter()

        elif op == StateTransition.EXIT:
            self._cleanup()
            return False

        else:
            print("ERROR: unrecognized transition.")
            return False

        if self.states.empty():
            return False

        return True

    def _cleanup(self):
        """Calls on_exit on each remaining state and pops it."""
        while not self.states.empty():
            self.states.pop().on_exit()


    
