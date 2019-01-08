import fileinput
from collections import deque, defaultdict
from typing import Optional
import copy

class Unit:
    reading_order = [(-1, 0), (0, -1), (0, 1), (1, 0)]
    def __init__(self, x, y, kind, field: "Field", ap = 3, hp = 200):
        self.x = x
        self.y = y
        self.kind = kind
        self.ap = ap
        self.hp = hp
        self.field = field

    def process_round(self):
        if self.hp <= 0:
            return

        direction = self.find_closest_direction()
        if direction != None:
            dx, dy = direction
            self.field.field[self.x][self.y] = "."
            unit = self.field.units[(self.x, self.y)]
            del self.field.units[(self.x, self.y)]
            self.x += dx
            self.y += dy
            self.field.field[self.x][self.y] = self.kind
            unit.x, unit.y = self.x, self.y
            self.field.units[(self.x, self.y)] = unit

        enemy = self.find_enemy_to_attack()
        if enemy != None:
            enemy.reduce_hp(self.ap)

    def get_enemy_kind(self):
        if self.kind == "G":
            return "E"

        return "G"

    def find_enemy_to_attack(self) -> Optional["Unit"]:
        min_hp = min((self.field.units[(self.x + dx, self.y + dy)].hp for dx, dy in self.reading_order if self.field.field[self.x + dx][self.y + dy] == self.get_enemy_kind()), default=0)
        for dx, dy in self.reading_order:
            if self.field.field[self.x + dx][self.y + dy] == self.get_enemy_kind() and self.field.units[(self.x + dx, self.y + dy)].hp == min_hp:
                return self.field.units[(self.x + dx, self.y + dy)]

        return None

    def find_closest_direction(self):
        field = self.field.field
        prev = {(self.x, self.y): (self.x, self.y)}
        visited = defaultdict(bool)
        visited[(self.x, self.y)] = True
        queue = deque([(self.x, self.y)])
        target = None
        while len(queue) > 0:
            cur = queue.pop()
            for dx, dy in self.reading_order:
                x, y = cur[0] + dx, cur[1] + dy
                place = field[x][y]
                if visited[(x, y)]:
                    continue

                if place == "." :
                    visited[(x, y)] = True
                    queue.appendleft((x, y))
                    if (x, y) not in prev:
                        prev[(x, y)] = cur
                elif place == self.get_enemy_kind():
                    target = cur
                    break

            if target is not None:
                break

        if target is None:
            return None 

        cur = target
        while prev[cur] != (self.x, self.y):
            cur = prev[cur]

        return (cur[0] - self.x, cur[1] - self.y)
    
    def reduce_hp(self, ap):
        self.hp = max(self.hp - ap, 0)
        if self.hp == 0:
            self.field.field[self.x][self.y] = "."
            del self.field.units[(self.x, self.y)]


class Field:
    def __init__(self, field, elf_ap):
        self.field = copy.deepcopy(field)
        self.units = {(x, y): Unit(x, y, cell, field=self, ap=elf_ap if cell == "E" else 3) for x, row in enumerate(field) for y, cell in enumerate(row) if cell in ["G", "E"]}
    
    def process_round(self):
        unit_list = sorted(self.units.items(), key=lambda item: (item[1].x, item[1].y) )
        for i, (_, unit) in enumerate(unit_list):
            unit.process_round()
            if self.is_complete() and i < len(unit_list) - 1:
                return True

        return False

    def is_complete(self):
        kinds = set([unit.kind for unit in self.units.values()])
        return len(kinds) == 1

    def get_sum_hp(self):
        return sum(unit.hp for unit in self.units.values())

    def print_field(self):
        for row in self.field:
            print("".join(row))

        print()

def run_battle(field_list, elf_ap):
    field = Field(field_list, elf_ap)
    initial_elf_count = sum(1 for unit in field.units.values() if unit.kind == "E")

    while True:
        field.process_round()
        if field.is_complete():
            elf_count = sum(1 for unit in field.units.values() if unit.kind == "E")
            if initial_elf_count == elf_count:
                return True

            break
    
    return False



if __name__ == "__main__":
    field_list = []
    for line in fileinput.input():
        row = list(line.strip())
        field_list.append(row)

    for elf_ap in range(4, 100):
        if run_battle(field_list, elf_ap):
            print(elf_ap)
            break
