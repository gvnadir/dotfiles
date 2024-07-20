#Write a program that interprets the Body Mass Index (BMI) based on a user's weight and height.

# 🚨 Don't change the code below 👇
height = float(input("enter your height in m: "))
weight = float(input("enter your weight in kg: "))
# 🚨 Don't change the code above 👆

#Write your code below this line 👇
bmi =  float(round(weight / height ** 2))
text1 = f"Your BMI is {bmi}, "

if bmi < 18.5:
    text2 = "you are underweight."
elif bmi < 25:
    text2 = "you have a normal weight."
elif bmi < 30:
    text2 = "you are slightly overweight."
elif bmi < 35:
    text2 = "you are obese."
else:
    text2 = "you are clinically obese."

print(text1 + text2)
