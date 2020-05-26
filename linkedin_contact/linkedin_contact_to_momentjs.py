#!/usr/bin/env python
import csv
from datetime import datetime
from dateutil.parser import parse

file_path="Invitations.csv"
with open(file_path, newline='') as csvfile:
  spamreader = csv.reader(csvfile, delimiter=',')
  spamreader.__next__()
  l = []
  for row in spamreader:
    incoming = True
    if row[4] != "INCOMING":
      incoming = False
    l.append((parse(row[2]), incoming))

  l.sort(key=lambda x: x[0])
  i = 0
  out = 0
  inc = 0
  outgoing = []
  incoming = []
  print('total = [', end='')
  for row in l:
    print('{x: moment("%s").toDate(), y: %s}, ' % (row[0], i), end='')
    if row[1]:
      incoming.append('{x: moment("%s").toDate(), y: %s},\n' % (row[0], inc))
      inc += 1
    else:
      outgoing.append('{x: moment("%s").toDate(), y: %s},\n' % (row[0], out))
      out += 1
    i+=1
  print('];', end='')
  print('incoming = [' + ''.join(incoming) + '];')
  print('outgoing = [' + ''.join(outgoing) + '];')
