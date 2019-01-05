import fileinput

class Unit:
    def __init__(self, x, y, kind, ap = 3, hp = 200, field = None):
        self.x = x
        self.y = y
        self.kind = kind
        self.ap = ap
        self.hp = hp
        self.field = field

    def set_field(self, field):
        pass

    def proess_round(self):
        pass

    def move(self):
        pass
    
    def attach(self):
        pass


class Field:
    def __init__(self, field):
        self.field = field
        self.units = [Unit(x,y,cell) for y, cell in enumerate(row) for x, row in enumerate(field) if cell in ["G", "E"]]
    
    def proces_round(self):
        self.units = sorted(self.units, lambda unit: (unit.x, unit.y) )
        for unit in self.units:
            unit.proces_round()

    def is_complete(self):
        pass

    def get_sum_hp(self):
        return sum(unit.hp for unit in self.units)


if __name__ == "__main__":
    field = []
    for line in fileinput.input():
        row = list(line.strip())
        field.append(row)

    round = 0
    while True:
        field.proces_round()
        round += 1 
        if field.is_complete():
            sum_hp = field.get_sum_hp()
            result = sum_hp * round
            print(result)
            return