import os

temp = os.popen('peer chaincode invoke -n mycc -c \'{"Args":["updateAvailableRandomarray","peer0","12998,2326,3498977,943678,234987"]}\' -C myc').readlines()
print(temp)
