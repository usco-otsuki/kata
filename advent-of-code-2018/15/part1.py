import fileinput
from typing import Optional

class Unit:
    reading_order = [(-1, 0), (0, -1), (0, 1), (1, 0)]
    def __init__(self, x, y, kind, field: "Field", ap = 3, hp = 200):
        self.x = x
        self.y = y
        self.kind = kind
        self.ap = ap
        self.hp = hp
        self.field = field

    def proess_round(self):
        if self.hp <= 0:
            return

        direction = self.find_closest_direction()
        if direction != None:
            dx, dy = direction
            self.field[self.x][self.y] = "."
            self.x += dx
            self.y += dy
            self.field[self.x][self.y] = self.kind
        enemy = self.find_enemy_to_attack()
        if enemy != None:
            enemy.reduce_hp(self.ap)

    def get_enemy_kind(self):
        if self.kind == "G":
            return "E"
        else:
            return "G"

    def find_enemy_to_attack(self) -> Optional["Unit"]:
        for dx, dy in self.reading_order:
            if self.field.field[self.x + dx][self.y + dy] == self.get_enemy_kind():
                return self.field.units[(self.x + dx, self.y + dy)]

        return None

    def find_closest_direction(self):
        # TODO: to implement
        return (0,0)
    
    def reduce_hp(self, ap):
        self.hp = max(self.hp - ap, 0)
        if self.hp == 0:
            self.field.field[self.x][self.y] = "."


class Field:
    def __init__(self, field):
        self.field = field
        self.units = {(x, y): Unit(x, y, cell, field=self) for y, cell in enumerate(row) for x, row in enumerate(field) if cell in ["G", "E"]}
    
    def process_round(self):
        unit_list = sorted(self.units.items(), lambda item: (item[1].unit.x, item[1].unit.y) )
        for x, y, unit in unit_list:
            unit.proces_round()
            if unit.hp == 0:
                del self.units[(x, y)]

    def is_complete(self):
        kinds = set([unit.kind for unit in self.units])
        return len(kinds) == 1

    def get_sum_hp(self):
        return sum(unit.hp for unit in self.units)


if __name__ == "__main__":
    field_list = []
    for line in fileinput.input():
        row = list(line.strip())
        field_list.append(row)

    field = Field(field_list)

    round = 0
    while True:
        field.process_round()
        round += 1 
        if field.is_complete():
            sum_hp = field.get_sum_hp()
            result = sum_hp * round
            print(result)
            break