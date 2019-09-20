import React, { useState } from "react";
import Router from "next/router";
import axios from "axios";
import { Container } from "../components/Container";

export default () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async e => {
    e.preventDefault();
    const data = {
      name: name,
      email: email,
      password: password
    };
    await axios.post("http://localhost:8080/signin", data).then(response => {
      console.log(response.data);
      window.localStorage.setItem("token", response.data);
      Router.push("/");
    });
  };

  const handleNameChange = e => {
    setName(e.target.value);
  };

  const handleEmailChange = e => {
    setEmail(e.target.value);
  };

  const handlePasswordChange = e => {
    setPassword(e.target.value);
  };

  return (
    <Container>
      <h1>ログイン</h1>
      <form
        onSubmit={e => {
          handleSubmit(e);
        }}
      >
        <div>ユーザー名</div>
        <input
          type="text"
          name="name"
          value={name}
          onChange={e => {
            handleNameChange(e);
          }}
        />
        <div>メールアドレス</div>
        <input
          type="text"
          name="email"
          value={email}
          onChange={e => {
            handleEmailChange(e);
          }}
        />
        <div>パスワード</div>
        <input
          type="password"
          name="password"
          value={password}
          onChange={e => {
            handlePasswordChange(e);
          }}
        />
        <button type="submit">ログインする</button>
      </form>
    </Container>
  );
};
