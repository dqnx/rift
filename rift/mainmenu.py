"""mainmenu

Implements the main menu user interface, game state, and input handling methods.
"""

from engine.gamestate import GameState
from state import State


class MenuState(State, GameState):
    def __init__(self):
        super().__init__()

        self.selected = 0
        self.options = ['Start', 'Exit']
        self.exit = True

        self._ui = MainMenuUI()

    def __str__(self):
        return "MenuState"

    def transition(self):
        if self.exit:
            return (StateTransition.EXIT, None)
        return (StateTransition.NONE, None)
    
    def update(self):
        pass

class MainMenuUI(UserInterface):
    def __init__(self):
        super().__init__()

        self.selection = 0
        self.menu_items = ['Start', 'Options', 'Exit']

    def render(self, console: tcod.console.Console):
        sets = Settings()
        screen_x, screen_y = sets.screen_size
        center_x = int(screen_x / 2)
        center_y = int(screen_y / 2)

        console.print(center_x, center_y, Settings().title, sets.colors['default_fg'], sets.colors['default_bg'], alignment=tcod.CENTER)

        row = 1
        for item in self.menu_items:
            console.print(center_x, center_y + row, item, sets.colors['default_fg'], sets.colors['default_bg'], alignment=tcod.CENTER)
            row += 1
    
    def select(self):
        pass

    def navigate(self, d: Direction):
        if d == Direction.UP:
            selection = max(0, selection - 1)
        elif d == Direction.DOWN:
            selection = min(len(self.menu_items), selection + 1)
