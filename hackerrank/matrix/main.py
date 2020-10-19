#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the minTime function below.
@profile
def minTime(roads, machines):
    machines = set(machines)
    def set_leader(node, leader):
        leaders[node] = leader

    def get_leader(node):
        if node == leaders[node]:
            return node
        leader = get_leader(leaders[node])
        if leader != leaders[node]:
            set_leader(node, leader)
        return leader

    leaders = {}
    for index in range(0, len(roads) + 1):
        leaders[index] = index
    sorted_roads = sorted(roads, key=lambda x: x[2], reverse=True)
    total_cost = 0
    print("start")
    to_destroy = len(machines) - 1
    for source, destination,  cost in sorted_roads:
        if to_destroy == 0:
            break
        source_leader = get_leader(source)
        destination_leader = get_leader(destination)

        if source_leader in machines and destination_leader in machines:
            total_cost += cost
            to_destroy -= 1
            continue

        if source_leader in machines:
            set_leader(destination_leader, source_leader)
            continue

        set_leader(source_leader, destination_leader)
    return total_cost


if __name__ == '__main__':
    fptr = sys.stdout

    nk = input().split()

    n = int(nk[0])

    k = int(nk[1])

    roads = []

    for _ in range(n - 1):
        roads.append(list(map(int, input().rstrip().split())))

    machines = []

    for _ in range(k):
        machines_item = int(input())
        machines.append(machines_item)

    result = minTime(roads, machines)

    fptr.write(str(result) + '\n')
