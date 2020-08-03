"""test_universe

Test routines for the universe module and orbital mechanics."""

import math
import unittest
import universe as uni
import engine.coordinate as coord

class TestBody(unittest.TestCase):
    def test_init(self):
        name = "test"
        mass = 133.7
        pos = coord.Cartesian(3,4,5)
        parent = uni.Body(name, mass, pos)

        b = uni.Body(name, mass, pos, parent=parent)
        self.assertEqual(b.name, name)
        self.assertEqual(b.mass, mass)
        self.assertEqual(b.position.x, pos.x)
        self.assertEqual(b.position.y, pos.y)
        self.assertEqual(b.position.z, pos.z)
        self.assertIs(b.parent, parent)

class TestKepler(unittest.TestCase):
    def test_keplers_eq(self):
        e = 0.2
        M_rad = math.pi/3 #radians
        test_E = 1.2361

        E, _ = uni._keplers_eq(e, M_rad, hist=True)
        
        self.assertAlmostEqual(E, test_E, places=3)

if __name__ == '__main__':
    unittest.main()