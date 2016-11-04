import React, { PropTypes } from 'react'

const Widget = ({children}) => (
  <div>
    <div className="container">
      <br/>
      {children}
    </div>
  </div>
)

Widget.propTypes = {
  children: PropTypes.node.isRequired
}

export default Widget
