import React, { PropTypes } from 'react'
import { translate } from 'react-i18next'

const Widget = ({t}) => (
  <div className="row">
    <hr/>
    <footer>
      <p className="pull-right">
        {t('this-site-in-other-languages')}
        <a href="/?locale=en-US"> {t('languages.english')} </a>
        <a href="/?locale=zh-CN"> {t('languages.simplified-chinese')} </a>
        <a href="/?locale=zh-TW"> {t('languages.traditional_chinese')} </a>
      </p>
      <p>
        &copy; 2016 Company, Inc.
        &middot;
        <a href="/?locale=en-US" target='_blank'> {t('languages.english')} </a>
        &middot;
        <a href="#">Terms</a>
      </p>
    </footer>
  </div>
)

Widget.propTypes = {
  t: PropTypes.func.isRequired
}

export default translate()(Widget)
