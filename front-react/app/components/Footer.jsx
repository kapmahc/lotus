import React, { PropTypes } from 'react'
import { translate } from 'react-i18next'
import { connect } from 'react-redux'

import {get} from '../ajax'
import {refresh} from '../engines/auth/actions'

const Widget = React.createClass({
  componentDidMount () {
    const {onRefresh} = this.props
    onRefresh()
  },
  render () {
    const {t} = this.props
    return (
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
  }
})

Widget.propTypes = {
  t: PropTypes.func.isRequired,
  onRefresh: PropTypes.func.isRequired
}

const Model = connect(
  state => ({info: state.siteInfo}),
  dispatch => ({
    onRefresh: function () {
      console.log('refresh')
      get('/siteInfo', null, function (ifo) {
        dispatch(refresh(ifo))
        document.documentElement.lang = ifo.lang
        document.title = ifo.title
      })
    }
  })
)(Widget)

export default translate()(Model)
