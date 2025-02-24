# Client-Server API Challenge - "May the dollar rate be forever in your favor!" 💸

Hi!! I`m Jessica... or sometimes Milena!
Welcome to this repository, where Go, context with timeouts and SQLite come together to create something cool! 😎

## 🎯 The Challenge
Build an API that enables client-server communication efficiently, securely, and – most importantly – without losing your sanity or breaking everything in the process (fingers crossed! 🤞).

## 🔨 How to Run

 ### 1. In Terminal #1: Start the server:
 `go run server.go`

 If everything’s peachy, you’ll see... 
 well, you`ll see.

 ### 2. In Terminal #2: Run the client:
 `go run client.go`

 And check the file **_cotacao.txt_** to see the current dollar rate. 💰
 
 ### 3. Where’s the Database?
 The server creates/uses cotacao.db to store the exchange rate history.
 **To take a peek, use SQLite:**
 `sqlite3 cotacao.db` or `sqlite3 /server/cotacao.db`
 
 **Inside the SQLite shell:**  
 `.tables` and `SELECT * FROM cotacoes;`

## 🚨 Possible Errors
 
- If the external API call takes longer than 200ms, the server logs an error and responds with an error message.
- If inserting data into the database takes more than 10ms, there’s another glorious error.
- If the client fails to get a response within 300ms, an error about timeouts will appear.

Nothing like a dash of timeout-induced drama to keep our code interesting! 🍿

## 🚀 Technologies Used

- **Go** - Because performance matters 🚀
- **SQLite** – A lightweight database for storing exchange rate history 📊

## 📜 License
 This project was created with learning objectives. Feel free to use and modify. Take the opportunity to check the dollar exchange rate and have fun learning Go!