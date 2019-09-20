import React from 'react'
import Header from '../components/Header'
import styled from 'styled-components'
import Link from 'next/link'

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

const Center = styled.div`
  margin: 0 auto;
`

const Name = styled.text`
  font-size: 1.5rem;
  padding: 2rem;
  text-align: center;
  color: #353b41;
`

const name = 'Yukio Orita'

export default () => (
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
          <Link href='/chat'>
            <Button>会話する</Button>
          </Link>
        </Center>
      </FlexContainer>
    </Container>
  </>
)
