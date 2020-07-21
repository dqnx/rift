from enum import Enum, unique

import tcod
import tcod.event

from engine.state import StateTransition
import engine.userinterface as ui

"""event

Implements an event handler for user input.

modified from: https://python-tcod.readthedocs.io/en/latest/tcod/event.html#tcod.event.EventDispatch
"""

MOVE_KEYS = {  # key_symbol: (x, y)
    # Arrow keys.
    tcod.event.K_LEFT: (-1, 0),
    tcod.event.K_RIGHT: (1, 0),
    tcod.event.K_UP: (0, -1),
    tcod.event.K_DOWN: (0, 1),
    tcod.event.K_HOME: (-1, -1),
    tcod.event.K_END: (-1, 1),
    tcod.event.K_PAGEUP: (1, -1),
    tcod.event.K_PAGEDOWN: (1, 1),
    tcod.event.K_PERIOD: (0, 0),
    # Numpad keys.
    tcod.event.K_KP_1: (-1, 1),
    tcod.event.K_KP_2: (0, 1),
    tcod.event.K_KP_3: (1, 1),
    tcod.event.K_KP_4: (-1, 0),
    tcod.event.K_KP_5: (0, 0),
    tcod.event.K_KP_6: (1, 0),
    tcod.event.K_KP_7: (-1, -1),
    tcod.event.K_KP_8: (0, -1),
    tcod.event.K_KP_9: (1, -1),
    tcod.event.K_CLEAR: (0, 0),  # Numpad `clear` key.
    # Vi Keys.
    tcod.event.K_h: (-1, 0),
    tcod.event.K_j: (0, 1),
    tcod.event.K_k: (0, -1),
    tcod.event.K_l: (1, 0),
    tcod.event.K_y: (-1, -1),
    tcod.event.K_u: (1, -1),
    tcod.event.K_b: (-1, 1),
    tcod.event.K_n: (1, 1),
}

SELECT_KEYS = [
    tcod.event.K_SPACE,
    tcod.event.K_RETURN
]

class InputHandler(tcod.event.EventDispatch[None]):
    """A state-based superclass that converts `events` into `commands`.

    The configuration used to convert events to commands are hard-coded
    in this example, but could be modified to be user controlled.

    Subclasses will override the `cmd_*` methods with their own
    functionality.  There could be a subclass for every individual state
    of your game.
    """

    def __init__(self):
        super().__init__()

    def ev_quit(self, event: tcod.event.Quit) -> None:
        """The window close button was clicked or Alt+F$ was pressed."""
        print(event)
        return self.cmd_quit()

    def ev_keydown(self, event: tcod.event.KeyDown) -> None:
        """A key was pressed."""
        print(event)
        if event.sym in MOVE_KEYS:
            # Send movement keys to the cmd_move method with parameters.
            return self.cmd_move(*MOVE_KEYS[event.sym])
        elif event.sym == tcod.event.K_ESCAPE:
            return self.cmd_escape()
        elif event.sym in SELECT_KEYS:
            return self.cmd_select()

    def ev_mousebuttondown(self, event: tcod.event.MouseButtonDown) -> None:
        """The window was clicked."""
        #print(event)

    def ev_mousemotion(self, event: tcod.event.MouseMotion) -> None:
        """The mouse has moved within the window."""
        #print(event)

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



