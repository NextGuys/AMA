import React from 'react'
import Header from '../components/Header'
import Container from '../components/Container'
import Layout from '../components/MyLayout.js'
import styled from 'styled-components'
import Link from 'next/link'

// const Index = () => (
//   <Layout>
//     <p>Hello Next.js</p>
//     {/* aタグを<Link>コンポーネントで囲い、hrefなどのリンクの情報は<Link>に渡します。 */}
//     <Link href="/second">
//       <button style={{ fontSize: 20 }}>Go to second Page</button>
//     </Link>
//   </Layout>
// )
// export default Index

const Text = styled.div`
  font-size: 80px;
`

// const TextLine = () => (
//     アイウエおりたあ
// )

export default () => (
  <Container>
    <Header>Next-Guys</Header>
    <h1>React合宿用</h1>
    <Text>アイウエおりたあ</Text>
  </Container>
)
