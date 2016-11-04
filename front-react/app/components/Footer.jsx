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
    const {t, info} = this.props
    return (
      <div className="row">
        <hr/>
        <footer>
          <p className="pull-right">
            {t('lang-bar.other')}
            <a href="/?locale=en-US"> {t('languages.english')} </a>
            <a href="/?locale=zh-CN"> {t('languages.simplified-chinese')} </a>
            <a href="/?locale=zh-TW"> {t('languages.traditional-chinese')} </a>
          </p>
          <p>
            &copy; {info.copyright}
            {info.bottomLinks.map(function (l, i) {
              return (
                <span key={i}>
                  &middot;
                  <a href={l.href}> {l.label} </a>
                </span>
              )
            })}
          </p>
        </footer>
      </div>
    )
  }
})

Widget.propTypes = {
  t: PropTypes.func.isRequired,
  info: PropTypes.object.isRequired,
  onRefresh: PropTypes.func.isRequired
}

const Model = connect(
  state => ({info: state.siteInfo}),
  dispatch => ({
    onRefresh: function () {
      get('/site-info', function (ifo) {
        dispatch(refresh(ifo))
        document.documentElement.lang = ifo.lang
        document.title = ifo.title
      })
    }
  })
)(Widget)

export default translate()(Model)
