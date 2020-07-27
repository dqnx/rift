from enum import Enum, unique

from bearlibterminal import terminal

from engine.state import StateTransition
import engine.userinterface as ui

"""event

Implements an event handler for user input.

modified from: https://python-tcod.readthedocs.io/en/latest/tcod/event.html#terminal.TEventDispatch
"""

MOVE_KEYS = {  # key_symbol: (x, y)
    # Arrow keys.
    terminal.TK_LEFT: (-1, 0),
    terminal.TK_RIGHT: (1, 0),
    terminal.TK_UP: (0, -1),
    terminal.TK_DOWN: (0, 1),
    terminal.TK_HOME: (-1, -1),
    terminal.TK_END: (-1, 1),
    terminal.TK_PAGEUP: (1, -1),
    terminal.TK_PAGEDOWN: (1, 1),
    terminal.TK_PERIOD: (0, 0),
    # Numpad keys.
    terminal.TK_KP_1: (-1, 1),
    terminal.TK_KP_2: (0, 1),
    terminal.TK_KP_3: (1, 1),
    terminal.TK_KP_4: (-1, 0),
    terminal.TK_KP_5: (0, 0),
    terminal.TK_KP_6: (1, 0),
    terminal.TK_KP_7: (-1, -1),
    terminal.TK_KP_8: (0, -1),
    terminal.TK_KP_9: (1, -1),
    # Vi Keys.
    terminal.TK_H: (-1, 0),
    terminal.TK_J: (0, 1),
    terminal.TK_K: (0, -1),
    terminal.TK_L: (1, 0),
    terminal.TK_Y: (-1, -1),
    terminal.TK_U: (1, -1),
    terminal.TK_B: (-1, 1),
    terminal.TK_N: (1, 1),
}

SELECT_KEYS = [
    terminal.TK_SPACE,
    terminal.TK_RETURN
]

class InputHandler():
    """A state-based superclass that converts `events` into `commands`.

    The configuration used to convert events to commands are hard-coded
    in this example, but could be modified to be user controlled.

    Subclasses will override the `cmd_*` methods with their own
    functionality.  There could be a subclass for every individual state
    of your game.
    """
    def dispatch(self, event: int):
        if (event == terminal.TK_CLOSE):
            return self.cmd_quit()
        # Key checking
        elif event in MOVE_KEYS:
            # Send movement keys to the cmd_move method with parameters.
            return self.cmd_move(*MOVE_KEYS[event])
        elif event == terminal.TK_ESCAPE:
            return self.cmd_escape()
        elif event in SELECT_KEYS:
            return self.cmd_select()

    # TODO: Implement mouse controls
    #def ev_mousebuttondown(self, event: tcod.event.MouseButtonDown) -> None:
    #def ev_mousemotion(self, event: tcod.event.MouseMotion) -> None:

    def cmd_move(self, x: int, y: int):
        """Intent to move: `x` and `y` is the direction, both may be 0."""
        #print("Command move: " + str((x, y)))
    
    def cmd_select(self):
        """Intent to make a selection."""
        #print("Command select.")

    def cmd_escape(self):
        """Intent to exit this state."""
        print("Command escape.")
        return self.cmd_quit()

    def cmd_quit(self):
        """Intent to exit the game."""
        print("Command quit.")
        return (StateTransition.EXIT, None)

@unique
class Orientation(Enum):
    VERTICAL = 0
    HORIZONTAL = 1

class SelectorMenu(InputHandler):
    """Input handler for a horizontal or vertical arrangement of UI Interactables."""
    def __init__(self, element: ui.UserInterface, orientation: Orientation):
        super().__init__()
        self._element = element
        self._orientation = orientation

    def cmd_move(self, x: int, y: int):
        """Move ui selector around."""
        if self._orientation is Orientation.VERTICAL:
            if x == 0: 
                active = ui.find_active(self._element)
                if active is not None:
                    if y < 0:
                        active.previous()
                    else:
                        active.next()
        elif self._orientation is Orientation.HORIZONTAL:
            if y == 0: 
                active = ui.find_active(self._element)
                if active is not None:
                    if x < 0:
                        active.previous()
                    else:
                        active.next()

        return None

    def cmd_select(self):
        """Intent to make a selection."""
        active = ui.find_active(self._element)
        if active is not None:
            return active.select()
        return None

    def cmd_escape(self):
        """Intent to exit this state."""
        print("Command escape.")
        return self.cmd_quit()

    def cmd_quit(self):
        """Intent to exit the game."""
        print("Command quit.")
        return (StateTransition.EXIT, None)



