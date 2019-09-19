import React from 'react'
import styled from 'styled-components'

export default function Card (props) {
  const { title, children, message } = props
  return <div {...props}>{children}</div>
}
