#!/usr/bin/env python3
#shit i may need

import pandas as pd
import io
import requests

#I struggled a lot to get the data, first I was using .readlines() but that was giving me strings and i failed to convert. 
#Then I discoveres this would give me a data frame
data = pd.read_fwf('input.txt', sep="/t", header=None, names =['depths'])

#my data eheh
print(data)

#ok i figured out how to read data from the row
print(data.iloc[5])

#trying to figure out the FOR loop now. Nevermind I think I dont need this in a data frame, i can use count.
print(data['depths'].count())

#Nope i need some kind of loop because i need to compare. Im gonna create a variable to keep track and try WHILE loop.
#And yes I know the 2000 implies that I know the size of the sample and i should use some kind of function to figure out but I'm tired.
count = 0
while (count < 2000):
    count = count + 1
print(count)

# #discovered this FOR loop with index. But now it tells me either 2000  is ut of range or -1 so i gotta limit somehow
#  count_if = 0
#  for i in data.index:
#      if (data['depths'][i] > data['depths'][i-1]):
#          count_if = count_if + 1
 
#  #meh i dont know
#  count_if = 0
#  for i in data.index:
#      if (data['depths'][i] < data['depths'][i+1]):
#          count_if = count_if + 1
#   if i == 1999:
# break
# print(count_if)
