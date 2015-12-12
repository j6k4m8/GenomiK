''' Find the gc_content '''
# gc_content = 0
# total = len(genome)
# for x in range(len(genome)):
#   if genome[x] == 'G' or genome[x] == 'C':
#     gc_content += 1

# gc_content = gc_content/float(total)

''' Generate the random unitigs '''
# for x in range(20):
#   num = rand.randint(0, len(genome)-1)
#   genome = genome[:num] + nt_map[genome[num]] + genome[num+1:]

# rands = []
# for x in range(20):
#   rands.append(rand.randint(0, len(genome)-1))

# rands.sort()
# rands.append(len(genome))

# unitigs = []

# for x in range(len(rands)-1):
#   if x is 0:
#     unitigs.append(genome[:rands[x+1]])
#   else:
#     unitigs.append(genome[rands[x-1]:rands[x+1]])


# output = ""
# for x in range(len(unitigs)):
#   if x == len(unitigs)-1:
#     output += unitigs[x]
#   else:
#     output += unitigs[x] + ","

# print(output)