import React, { PropTypes } from 'react'

import Header from './Header'
import Footer from './Footer'

const Widget = ({children}) => (
  <div>
    <Header/>
    <div className="container">
      <br/>
      {children}
      <Footer/>
    </div>
  </div>
)

Widget.propTypes = {
  children: PropTypes.node.isRequired
}

export default Widget
