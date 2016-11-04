import React, { PropTypes } from 'react'
import { translate } from 'react-i18next'
import {NavDropdown, MenuItem} from 'react-bootstrap'

const Widget = ({t}) => (
  <NavDropdown eventKey={3} title={t('languages.switch')} id="switch-lang-bar">
    <MenuItem eventKey={3.1} href="/?locale=en-US">
      {t('languages.english')}
    </MenuItem>
    <MenuItem eventKey={3.2} href="/?locale=zh-CN">
      {t('languages.simplified-chinese')}
    </MenuItem>
    <MenuItem eventKey={3.3} href="/?locale=zh-TW">
      {t('languages.traditional_chinese')}
    </MenuItem>
  </NavDropdown>
)

Widget.propTypes = {
  t: PropTypes.func.isRequired
}

export default translate()(Widget)
