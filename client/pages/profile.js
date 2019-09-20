import React, { useState } from 'react'
import Header from '../components/Header'
import styled from 'styled-components'
import Link from 'next/link'
import axios from 'axios'

const Button = styled.button`
  width: 150px;
  height: 40px;
  margin: 10px;
  text-align: center;
  justify-content: center;
`

const Container = styled.article`
  background: #f6f6f6;
  width: 100%;
`

const FlexContainer = styled.div`
  width: 174px;
  margin: 0 auto;
`
const Card = styled.div`
  margin: 10px;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
  transition: 0.3s;
  border-radius: 5px;
  background-color: #ffffff;
`

const CardContainer = styled.div`
  padding: 12px 16px;
`
const Text = styled.div`
  margin: 5px 5px 5px 5px;
  display: inline-block;
  -webkit-text-decoration: none;
  text-decoration: none;
  color: #668ad8;
  border-radius: 5px;
  border: solid 2px #668ad8;
  text-align: center;
  overflow: hidden;
  font-weight: bold;
  -webkit-transition: 0.4s;
  transition: 0.4s;
  padding: 10px 10px 10px 10px;
`

const Tittle = styled.div`
  margin: 5px 5px 5px 5px;
  text-decoration: none;
  text-align: center;
  overflow: hidden;
  font-weight: bold;
  -webkit-transition: 0.4s;
  transition: 0.4s;
  padding: 10px 10px 10px 10px;
  width: 135px;
`

const Center = styled.div`
  margin: 0 auto;
`

const Name = styled.text`
  font-size: 1.5rem;
  padding: 2rem;
  text-align: center;
  color: #353b41;
`

const Form = styled.form`
  margin-left: 20px;
`

const name = 'Yukio Orita'

export default () => {
  const handleOnChange = e => {
    setRoomTitle(e.target.value)
  }

  const handleSubmit = e => {
    axios.post('http://localhost:8080/rooms', {
      headers: {
        Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTY5MDM0MzQ4LCJpYXQiOiIyMDE5LTA5LTIwVDExOjUyOjI4LjMzODQ4OSswOTowMCIsIm5hbWUiOiJ5dWtpbyIsInN1YiI6ImJtMjN0a2kzcTU2Mm1nMzJudmZnIn0.ZwaqrkynGbck5h6wJSX20yI95PQ28gaCGE2iXksNUq0
    `
      },
      data: {
        title: 'sss',
        uid: 'fregrhjtfghfjgh'
      }
    })
  }

  const [roomTitle, setRoomTitle] = useState('')

  return (
    <>
      <Header>Profile</Header>
      <Container>
        <Card>
          <CardContainer>
            <Name>
              <b>{name}</b>
            </Name>
            {/* <TextBody>
            {/* {name.map(name => {
              return <Text>{name}</Text>
            })} */}
            {/* </TextBody>  */}
          </CardContainer>
        </Card>
        <FlexContainer>
          <Center>
            <Tittle>トークルームのタイトルを入力してください。</Tittle>
            <Form onSubmit={handleSubmit}>
              <input
                onChange={e => {
                  handleOnChange(e)
                }}
                value={roomTitle}
              />
              <Link href='/chat'>
                <Button type='submit'>会話する</Button>
              </Link>
            </Form>
          </Center>
        </FlexContainer>
      </Container>
    </>
  )
}
