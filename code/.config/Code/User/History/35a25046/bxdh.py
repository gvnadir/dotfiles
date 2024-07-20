print("The Love Calculator is calculating your score...")
name1 = input() # What is your name?
name2 = input() # What is their name?
# 🚨 Don't change the code above 👆
# Write your code below this line 👇
true = "TRUE"
love = "LOVE"
name1_score = 0
name2_score = 0

def score_calculator(name):
	for x in true:
		true_counter += name.lower().count(x.lower())

	for y in love:
		love_counter += name.lower().count(y.lower())
	
	return true_counter + love_counter

name1_score = score_calculator(name1)
name2_score = score_calculator(name2)
total = name1_score + name2_score

print(f'${name1}: ${name1_score}')
print(f'${name2}: ${name2_score}')
print(f'total: ${total}')