# Server-DB-GoLang-


**Job evaluation:**

Download and install PostgreSQL – 1 hours.

create DB on PostgreSQL and init - 15 min.

Open VsCode and up service  that connects to the DB that we opened in the previous section (Server side) – 2.5 hours
•	implement middleware (API)
1)	POST
2)	GRT 
3)	DELETE
4)	PUT
•	Need to check through postman agent if all the middleware we made works.

STACK REVIEW – 45 min

**REVIEW**

![תמונה](https://user-images.githubusercontent.com/57719538/124265487-6a21fe00-db3e-11eb-8e04-8c0b3528d171.png)



**Comment** 

psql -U postgres – open shell Postgres.
\L – All Databases in computer.
\C [database] – connect to DB.


•	Create new database:

![תמונה](https://user-images.githubusercontent.com/57719538/124265620-99386f80-db3e-11eb-8581-845048f876f9.png)


CREATE DATABASE users;


•	By \l commend we can see the new database call "users"

![תמונה](https://user-images.githubusercontent.com/57719538/124265731-cd139500-db3e-11eb-9bc6-493db1d80296.png)

 
•	now I connect to "users" database with \c users commend.

•	By commend "\dt" can see all relation.

•	You can see in a screenshot our table we created.

![תמונה](https://user-images.githubusercontent.com/57719538/124265743-d4d33980-db3e-11eb-8967-60ff5f355fb6.png)

 
 **GO**
 
 Install Go:

 apt install go-glang
 
 
 You must import github lib and therefore you need to write down the following command:

  go get github.com/lib/pq
  
 Note:
The server is listening to port 8080 and by postadmin and by HTTP requests I did tests.
 
 
**Example:**

*POST - /users*
![תמונה](https://user-images.githubusercontent.com/57719538/126353290-47958e21-b76f-40cc-a69f-82f0240a029d.png)


*GET - /users*

![תמונה](https://user-images.githubusercontent.com/57719538/126353624-6501dc1e-1144-495c-93de-94d3c4849b87.png)


*DELETE - /users/0*

![תמונה](https://user-images.githubusercontent.com/57719538/126353868-84093750-2cfe-4a26-afed-54f77fc1c708.png)

*PUT - /users/0*

![תמונה](https://user-images.githubusercontent.com/57719538/126354604-f9008797-c762-4f6c-a0a8-4882102635e6.png)








