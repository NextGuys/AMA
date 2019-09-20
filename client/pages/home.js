import React from 'react'
import Header from '../components/Header'
import {
  RightContainer,
  LeftContainer,
  Container
} from '../components/Container'
import styled from 'styled-components'
import Link from 'next/link'

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

const SkillText = styled.div`
  margin: 5px 5px 5px 5px;
  display: inline-block;
  -webkit-text-decoration: none;
  text-decoration: none;
  color: #ff69a3;
  border-radius: 5px;
  border: solid 2px #ff69a3;
  text-align: center;
  overflow: hidden;
  font-weight: bold;
  -webkit-transition: 0.4s;
  transition: 0.4s;
  padding: 10px 10px 10px 10px;
`

const TextBody = styled.div`
  display: flex;
`

const Form = styled.form`
  margin: 20px;
  text-align: center;
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

const name = ['Yukio Orita', 'Shogo Yasuda']
const titles = ['React Hooksについて', '【質問】Atomic Design']
const users = [
  'Yukio Orita',
  'Shogo Yasuda',
  'Katsuhito Karube',
  'Nobuo Orita',
  'Maruko Orita'
]
const skills = ['React', 'Go']

export default () => (
  <>
    <Header>AMA</Header>
    <Container>
      <LeftContainer>
        <Form>
          <label>検索</label>
          <div>
            <input type='text' name='name' />
            <input type='submit' value='Submit' />
          </div>
        </Form>

        {titles.map(title => {
          return (
            <Link href='/chat'>
              <Card>
                <CardContainer>
                  <h4>
                    <b>{title}</b>
                  </h4>
                  <TextBody>
                    {name.map(name => {
                      return <Text>{name}</Text>
                    })}
                  </TextBody>
                </CardContainer>
              </Card>
            </Link>
          )
        })}
      </LeftContainer>

      <RightContainer>
        {users.map(users => {
          return (
            <Link href='/profile'>
              <Card>
                <CardContainer>
                  <h4>
                    <b>{users}</b>
                  </h4>
                  <TextBody>
                    {skills.map(skills => {
                      return <SkillText>{skills}</SkillText>
                    })}
                  </TextBody>
                </CardContainer>
              </Card>
            </Link>
          )
        })}
      </RightContainer>
    </Container>
  </>
)
