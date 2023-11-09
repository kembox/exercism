"""Functions to automate Conda airlines ticketing system."""


def generate_seat_letters(number):
    letters = 'ABCD'* (int(number/4) + 1)
    #Maybe have some extra letters but ok

    for i in range(number):
        yield letters[i]

def generate_seats(number):
    for i,v in enumerate(generate_seat_letters(number)):
            skip = 0
            if (int(i/4) + 1) < 13:
                yield str(int(i/4)+1) + v
            else:
                yield str(int(i/4)+2) + v

def assign_seats(passengers):
    d=dict()
    for i,v in enumerate(generate_seats(len(passengers))):
        d[passengers[i]] = v
    return d

def generate_codes(seat_numbers, flight_id):
    for x in seat_numbers:
        code=x + flight_id
        yield f'{code:0<12}'