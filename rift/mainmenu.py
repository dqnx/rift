"""mainmenu

Implements the main menu user interface, game state, and input handling methods.
"""
import tcod

from engine.gamestate import GameState
from engine.state import State, StateTransition
from engine.input_handlers import SelectorMenu, Orientation
from engine.settings import Settings
from worldgen import WorldGenState

import engine.userinterface as ui

class MenuState(State, GameState):
    def __init__(self):
        super().__init__()

        sets = Settings()
        self._ui = MainMenuUI(*sets.screen_size, None) # TODO: more flexible, work for all menus?
        self._input_handler = SelectorMenu(self._ui, Orientation.VERTICAL)
        self.exit = False

    def __str__(self):
        return "MenuState"

    def run(self, state_input):
        transition = self._input_handler.dispatch(state_input)
        if transition:
            return transition
        return (StateTransition.NONE, None)
    
class MainMenuUI(ui.UserInterface):
    def __init__(self, width: int, height: int,  parent: ui.UserInterface, x: int = 0, y: int = 0):
        super().__init__(width, height, parent, x=x, y=y)

        sets = Settings()
        origin_x = int(width/2) - 5
        origin_y = int(height/2) - 5

        # Add text element for the title
        self._children.append(ui.Text("RIFT", sets.colors['default_fg'], 10, 1, self,
            x=origin_x, y=origin_y))
        
        # Define start selector as activated
        def _start():
            return (StateTransition.NONE, None)
        self._children.append(ui.Selector("Start", 10, 1, self,
            action=_start, x=origin_x, y=origin_y+2))
        self._children[1].toggle()

        # Define world gen activator
        def _worldgen():
            return (StateTransition.NEXT, WorldGenState())
        self._children.append(ui.Selector("World Generation", 10, 23, self,
            action=_worldgen, x=origin_x, y=origin_y+3))

        # Define options selector
        def _options():
            return (StateTransition.NONE, None)
        self._children.append(ui.Selector("Options", 10, 1, self,
            action=_options, x=origin_x, y=origin_y+4))

        # Define exit selector
        def _exit():
            return (StateTransition.EXIT, None)
        self._children.append(ui.Selector("Exit", 10, 1, self,
            action=_exit, x=origin_x, y=origin_y+5))

    def _render(self, console: tcod.console.Console, offset: (int, int) = (0, 0)):
        # MainMenuUI only renders children.
        pass

