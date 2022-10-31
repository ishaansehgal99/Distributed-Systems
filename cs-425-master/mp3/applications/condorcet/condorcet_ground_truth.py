# ground-truth.py FILE
import os
import sys
from collections import defaultdict

args = sys.argv
filename = args[1]

print('Reading...')
with open(filename, 'r') as text_file:
    lines = text_file.read().splitlines() 

print('Read {} lines.'.format(len(lines)))
    
# calculate who beats who how many times
beats = defaultdict(lambda: defaultdict(int))    
for line in lines:
    order = line.split(' ')[1].split(',')
    
    for i in range(len(order)):
        for j in range(i + 1, len(order)):
            winner = order[i]
            loser = order[j]
            
            beats[winner][loser] += 1
            
# print results    
candidates = sorted([k for k in beats])
for a in candidates:
    print(a + "'s net votes over others:")
    for b in candidates:
        if a == b:
            continue
            
        a_over_b = beats[a][b]
        b_over_a = beats[b][a]
        
        print("===>", b + ':', a_over_b - b_over_a)
    print()
