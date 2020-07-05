"""singleton

Defines a singleton class for subclassing.
Use the @singleton decorator.
Referenced from: https://www.python.org/dev/peps/pep-0318/
"""

def singleton(cls):
    instances = {}
    def getinstance():
        if cls not in instances:
            instances[cls] = cls()
        return instances[cls]
    return getinstance