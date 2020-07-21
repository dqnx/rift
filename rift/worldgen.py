"""worldgen

Implements a world generation algorithm and visualization.
"""

import tcod

from engine.input_handlers import InputHandler
from engine.settings import Settings
from engine.gamestate import GameState
from engine.state import State, StateTransition
import engine.userinterface as ui

class WorldGenState(State, GameState):
    def __init__(self):
        super().__init__()

        sets = Settings()
        self._ui = WorldGenUI(*sets.screen_size, None)
        self._input_handler = InputHandler()
        self.exit = False

    def __str__(self):
        return "WorldGenState"

    def run(self, state_input):
        transition = self._input_handler.dispatch(state_input)
        if transition:
            return transition
        return (StateTransition.NONE, None)

class WorldGenUI(ui.UserInterface):
    def __init__(self, width: int, height: int,  parent: ui.UserInterface, x: int = 0, y: int = 0):
        super().__init__(width, height, parent, x=x, y=y)

        sets = Settings()
        # Add text element for the title
        self._children.append(ui.Text("World Generation", sets.colors['default_fg'], 20, 1, self,
            x=0, y=0))

    def _render(self, console: tcod.console.Console, offset: (int, int) = (0,0)):
        # WorldGen only renders children.
        pass