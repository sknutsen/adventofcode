inputFile = open("input.txt", "r")

lines = inputFile.readlines()

sumOfValidGames: int = 0
sumOfPowers: int = 0

maxDict = {
    "red": 12,
    "green": 13,
    "blue": 14,
}

minDict: dict = {}

for line in lines:
    game = line.split(": ")
    print(game)
    
    num = int(game[0].split(" ")[1])
    
    minDict[num] = {
        "red": 0,
        "green": 0,
        "blue": 0,
    }
    
    gameIsValid: bool = True
    
    for gameRound in game[1].strip("\n").split("; "):
        for cubes in gameRound.split(", "):
            key = cubes.split(" ")[1]
            val = int(cubes.split(" ")[0])
            
            if maxDict[key] < val:
                gameIsValid = False
                
            if not minDict[num][key] or minDict[num][key] < val:
                minDict[num][key] = val
                
    if gameIsValid:
        sumOfValidGames += num
        
    print(minDict[num])
    sumOfPowers += minDict[num]["red"] * minDict[num]["green"] * minDict[num]["blue"]
        
print("total games " + str(sumOfValidGames))
print("sum of powers " + str(sumOfPowers))
