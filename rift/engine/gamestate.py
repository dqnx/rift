"""gamestate

Implements a state machine and states for the game.
"""
from abc import ABC, abstractmethod
from engine.state import StateMachine


class GameState(ABC):
    """Abstract base class defining render (with a UI object) and update methods."""
    def __init__(self):
        self._ui = None
        self._input_handler = None

    @property
    def ui(self):
        return self._ui

    def render(self):
        self.ui.render()

class GameStateMachine(StateMachine):
    def __init__(self, initial_state: GameState):
        super().__init__()
        self.states.push(initial_state)

    def render(self):
        self.states.top().render()

