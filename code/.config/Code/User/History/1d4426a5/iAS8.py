print("The Love Calculator is calculating your score...")
name1 = input()  # What is your name?
name2 = input()  # What is their name?
# 🚨 Don't change the code above 👆
# Write your code below this line 👇

# Test: Brad Pitt Jennifer Aniston = 72


def love_calculator(combined_names):
    true = "TRUE"
    love = "LOVE"
    true_counter = 0
    love_counter = 0
    for x in true:
        true_counter += combined_names.count(x.lower())

    for y in love:
        love_counter += combined_names.count(y.lower())

    return int(str(true_counter) + str(love_counter))


combined_names = (name1 + name2).lower()
score = love_calculator(combined_names)

if score < 10 or score > 90:
    print(f"Your score is {score}, you go together like coke and mentos.")
elif 40 <= score <= 50:
    print(f"Your score is {score}, you are alright together.")
else:
    print(f"Your score is {score}.")
