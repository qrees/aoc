#!/bin/python3

import math
import os
import random
import re
import sys
from collections import defaultdict

groups = {}

def get_leader(city):
    if groups[city] == city:
        return city
    else:
        leader = get_leader(groups[city])
        set_leader(city, leader)
        return leader

def set_leader(city, leader):
    groups[city] = leader

# Complete the roadsAndLibraries function below.
def roadsAndLibraries(n, c_lib, c_road, cities):
    global groups
    groups = {}
    for city in range(1, n+1):
        groups[city] = city
    for city_pair in cities:
        city_pair = sorted(city_pair)
        from_city, to_city = city_pair
        from_city = get_leader(from_city)
        to_city = get_leader(to_city)
        if from_city != to_city:
            set_leader(from_city, to_city)
    road_count = defaultdict(lambda: -1)

    print(groups)
    for city in groups:
        leader = get_leader(city)
        road_count[leader] = road_count[leader] + 1
    suma = 0
    print(dict(road_count), c_road, c_lib)
    for count in road_count.values():
        if (count * c_road) + c_lib > (count + 1) * c_lib:
            # build libraries
            suma += (count + 1) * c_lib
            print("libraries", suma)
        else:
            # build roads
            suma += (count * c_road) + c_lib
            print("roads", suma)

    return suma


if __name__ == '__main__':
    fptr = sys.stdout
    # fptr = open(os.environ['OUTPUT_PATH'], 'w')

    q = int(input())

    for q_itr in range(q):
        nmC_libC_road = input().split()

        n = int(nmC_libC_road[0])

        m = int(nmC_libC_road[1])

        c_lib = int(nmC_libC_road[2])

        c_road = int(nmC_libC_road[3])

        cities = []

        for _ in range(m):
            cities.append(list(map(int, input().rstrip().split())))

        result = roadsAndLibraries(n, c_lib, c_road, cities)

        fptr.write(str(result) + '\n')

    # fptr.close()
