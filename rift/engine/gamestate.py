"""gamestate

Implements a state machine and states for the game.
"""
from abc import ABC, abstractmethod
import tcod
from state import StateMachine


class GameState(ABC):
    """Abstract base class defining render (with a UI object) and update methods."""
    def __init__(self):
        self._ui = None

    @property
    def ui(self):
        return self._ui

    @abstractmethod
    def update(self):
        pass

    def render(self, console: tcod.console.Console):
        self.ui.render(console)

class GameStateMachine(StateMachine):
    def __init__(self, initial_state: GameState):
        super().__init__()
        self.states.push(initial_state)

    def update(self):
        print("Updating:", self.states.top().name())

    def render(self):
        print("Rendering:", self.states.top().name())
