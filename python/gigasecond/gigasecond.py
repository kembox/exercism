one_gigasec = 1_000_000_000
from datetime import timedelta
def add(moment):
        return moment + timedelta(seconds=one_gigasec)