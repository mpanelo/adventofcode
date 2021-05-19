import os

def main():
    total_fuel_requirement = 0
    for mass in read_input():
        total_fuel_requirement += calculate_fuel_requirement(mass)
    print(f"Total Fuel Requirement: {total_fuel_requirement}")


def read_input():
    curdir = os.path.dirname(__file__) 
    datapath = os.path.join(curdir, './../../puzzledata/day01/input.txt')

    with open(datapath, 'r') as file:
        for line in file:
            yield int(line.strip())


def calculate_fuel_requirement(mass, total_fuel=0):
    fuel = mass // 3 - 2
    if fuel <= 0:
        return total_fuel
    return calculate_fuel_requirement(fuel, total_fuel+fuel) 


if __name__ == "__main__":
    main()