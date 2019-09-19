/* 1. expressモジュールをロードし、インスタンス化してappに代入。*/
var express = require("express");

// get the client
const mysql = require("mysql2");

var app = express();

/* 3. 以後、アプリケーション固有の処理 */

// create the connection to database
const connection = mysql.createConnection({
  host: "localhost",
  user: "root",
  database: "ama",
  password: "password"
});

app.post("/messages", function(req, res, next) {
  console.log(req);
  console.log(res);
  // res.json(photoList);

  // simple query
  // connection.query("INSERT INTO `messages` VALUES (`0`, `0`, ``)", function(
  //   err,
  //   results,
  //   fields
  // ) {
  //   console.log(results); // results contains rows returned by server
  //   console.log(fields); // fields contains extra meta data about results, if available
  // });

  res.json({ result: "ok" });
});

/* 2. listen()メソッドを実行して3000番ポートで待ち受け。*/
var server = app.listen(3001, function() {
  console.log("Node.js is listening to PORT:" + server.address().port);
});
