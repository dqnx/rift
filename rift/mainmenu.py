"""mainmenu

Implements the main menu user interface, game state, and input handling methods.
"""
import tcod

from engine.gamestate import GameState
from engine.state import State, StateTransition
from engine.userinterface import UserInterface, Direction
from engine.input_handlers import InputHandler
from settings import Settings

class MenuState(State, GameState):
    def __init__(self):
        super().__init__()

        self._ui = MainMenuUI() # TODO: more flexible, work for all menus?
        self._input_handler = MenuInputHandler(self._ui)
        self.exit = False

    def __str__(self):
        return "MenuState"

    def run(self, state_input):
        ui_action = self._input_handler.dispatch(state_input)
        if ui_action:
            return ui_action
        return (StateTransition.NONE, None)
    
class MainMenuUI(UserInterface):
    def __init__(self):
        super().__init__()

        self.selection = 0
        self.menu_items = [
            ('Start', (StateTransition.NONE, None)),
            ('Options', (StateTransition.NONE, None)),
            ('Exit', (StateTransition.EXIT, None))
        ]

    def render(self, console: tcod.console.Console):
        sets = Settings()
        screen_x, screen_y = sets.screen_size
        center_x = int(screen_x / 2)-4
        center_y = int(screen_y / 2)-5

        # print title
        console.print(center_x, center_y, Settings().title, sets.colors['default_fg'], sets.colors['default_bg'], alignment=tcod.LEFT)

        # print menu items
        menu_offset = 2
        row = menu_offset
        for item in self.menu_items:
            console.print(center_x, center_y + row, item[0], 
                sets.colors['default_fg'], sets.colors['default_bg'], alignment=tcod.LEFT)
            row += 1
        
        # print selector
        console.print(center_x - 1, center_y + menu_offset + self.selection, '*', 
            sets.colors['default_fg'], sets.colors['default_bg'], alignment=tcod.CENTER)
        
    def select(self):
        return self.menu_items[self.selection][1]

    def navigate(self, x: int, y: int):
        if x == 0:
            self.selection += y
            self.selection = max(0, self.selection)
            self.selection = min(self.selection, len(self.menu_items)-1)

class MenuInputHandler(InputHandler):
    def __init__(self, ui: UserInterface):
        super().__init__()
        self._ui = ui

    def cmd_move(self, x: int, y: int) -> (int, int):
        """Move ui selector around."""
        self._ui.navigate(x, y)

    def cmd_select(self) -> None:
        """Intent to make a selection."""
        return self._ui.select()

    def cmd_escape(self) -> None:
        """Intent to exit this state."""
        print("Command escape.")
        return self.cmd_quit()

    def cmd_quit(self) -> None:
        """Intent to exit the game."""
        print("Command quit.")
        return (StateTransition.EXIT, None)
