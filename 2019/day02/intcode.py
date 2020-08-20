import itertools

TARGET_SOLUTION = 19690720


def main():
    program = IntcodeProgram(memory=VirtualMemory(src='input.txt'))
    for noun, verb in itertools.product(range(100), repeat=2):
        solution = program.simulate([(1, noun), (2, verb)])

        if solution == TARGET_SOLUTION:
            print("Solution Found!")
            print(f"100 * noun + verb = {100 * noun + verb}")
            break


class VirtualMemory:
    def __init__(self, src='input.txt'):
        self.src = src
        self.memory = None
        self.reset()

    def reset(self):
        with open(self.src, 'r') as file:
            self.memory = [int(value) for value in file.read().strip().split(',')]

    def get(self, address):
        return self.memory[address]

    def set(self, address, value):
        self.memory[address] = value


class IntcodeProgram:
    INSTR_OFFSET = 4
    OPCODE_ADD = 1
    OPCODE_MULT = 2
    OPCODE_HALT = 99

    def __init__(self, memory: VirtualMemory):
        self.memory = memory
        self.instr_pointer = 0

    def simulate(self, memory_values):
        self._prepare_memory(memory_values)

        while True:
            instr_opcode = self._read_pointer()

            if instr_opcode == self.OPCODE_ADD:
                p1 = self._read_pointer()
                p2 = self._read_pointer()
                dest = self._read_pointer()

                self.memory.set(dest, self.memory.get(p1) + self.memory.get(p2))
            elif instr_opcode == self.OPCODE_MULT:
                p1 = self._read_pointer()
                p2 = self._read_pointer()
                dest = self._read_pointer()

                self.memory.set(dest, self.memory.get(p1) * self.memory.get(p2))
            elif instr_opcode == self.OPCODE_HALT:
                break
            else:
                raise Exception(f"Unknown opcode {instr_opcode}!")

        return self.memory.get(0)

    def _read_pointer(self):
        value = self.memory.get(self.instr_pointer)
        self.instr_pointer += 1
        return value

    def _prepare_memory(self, memory_values):
        self.memory.reset()
        self.instr_pointer = 0
        for address, value in memory_values:
            self.memory.set(address, value)


if __name__ == "__main__":
    main()
