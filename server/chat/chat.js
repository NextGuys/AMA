/* 1. expressモジュールをロードし、インスタンス化してappに代入。*/
const express = require("express");
const bodyParser = require("body-parser");

// get the client
const mysql = require("mysql2");
const cors = require("cors");

const app = express();
app.use(cors());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

/* 3. 以後、アプリケーション固有の処理 */

// create the connection to database
const connection = mysql.createConnection({
  host: "localhost",
  user: "root",
  database: "ama",
  password: "password"
});

app.post("/messages", (req, res, next) => {
  const message = req.body.value;

  // simple query
  const sql = `INSERT INTO messages (user_id, room_id, content) VALUES (0, 0, "${message}")`;
  connection.query(sql, function(err, results, fields) {
    console.log(results); // results contains rows returned by server
    console.log(fields); // fields contains extra meta data about results, if available
  });

  res.json({ result: "ok" });
});

/* 2. listen()メソッドを実行して3000番ポートで待ち受け。*/
const server = app.listen(3001, function() {
  console.log("Node.js is listening to PORT:" + server.address().port);
});
