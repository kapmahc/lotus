import React, { PropTypes } from 'react'

import Header from './Header'
import Footer from './Footer'

const Widget = ({children}) => (
  <div>
    <Header/>
    layout
    {children}
    <Footer/>
  </div>
)

Widget.propTypes = {
  children: PropTypes.node.isRequired
}

export default Widget
