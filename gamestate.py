"""gamestate

Implements a state machine and states for the game.
"""

from state import State, StateMachine, StateTransition

class MenuState(State):
    def __init__(self):
        self.selected = 0
        self.options = ['Start', 'Exit']
        self.exit = True

    def name(self):
        return "MenuState"

    def transition(self):
        if self.exit:
            return (StateTransition.EXIT, None)
        return (StateTransition.NONE, None)

class GameStateMachine(StateMachine):
    def __init__(self):
        super().__init__()
        self.states.push(MenuState())

    def update(self):
        print("Updating:", self.states.top().name())

    def render(self):
        print("Rendering:", self.states.top().name())
