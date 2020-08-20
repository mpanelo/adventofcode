OPCODE_OFFSET = 4
OPCODE_ADD = 1
OPCODE_MULT = 2
OPCODE_HALT = 99


def main():
    sequence = read_input()
    process(sequence)
    print(sequence)


def read_input():
    with open('input.txt', 'r') as file:
        return [int(i) for i in file.read().strip().split(',')]



def process(sequence):
    for i in range(0, len(sequence), OPCODE_OFFSET):
        opcode, p1, p2, dest = sequence[i:i+OPCODE_OFFSET]

        if opcode == OPCODE_ADD:
            sequence[dest] = sequence[p1] + sequence[p2]
        elif opcode == OPCODE_MULT:
            sequence[dest] = sequence[p1] * sequence[p2]
        elif opcode == OPCODE_HALT:
            return
        else:
            raise Exception(f"Unknown opcode {opcode}!")


if __name__ == "__main__":
    main()