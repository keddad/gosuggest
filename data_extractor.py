"""
This script transforms the data from the original dataset into a format that can be used by the suggestion engine.
"""


import csv

movie_ratings = {}
movie_names = {}

rate_num = 3_000_000

with open("ratings.csv", "r") as file:
    reader = csv.reader(file)
    next(reader, None)
    for row in reader:
        if rate_num == 0:
            break

        rate_num -= 1
        movie_ratings[int(row[1])] = movie_ratings.get(
            int(row[1]), []) + [row[2]]

for k, v in movie_ratings.items():
    movie_ratings[k] = sum(map(float, v)) / len(v)

with open("movies_metadata.csv", "r") as file:
    reader = csv.reader(file)
    next(reader, None)
    for row in reader:
        try:
            movie_names[int(row[5])] = row[8]
        except:
            pass  # ignore movies with broken id

with open("data.csv", "w") as csvfile:
    writer = csv.writer(csvfile)
    writer.writerow(["id", "name", "rating"])

    for k, v in movie_names.items():
        if k not in movie_ratings:
            continue

        writer.writerow([k, v, int(movie_ratings[k])])
