def transform(legacy_data):
    d = dict()
    for k,v in legacy_data.items():
        for i in v:
            d[i.lower()] = k
    return d