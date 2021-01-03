import { useHistory, useParams } from 'react-router-dom'
import Button from 'antd/es/button'
import Input from 'antd/es/input'
import { useEffect, useState } from 'react'
import { KeplerGl } from 'kepler.gl/components'
import styles from './ReportPage.module.css'
import { AutoSizer } from 'react-virtualized'
import { useDispatch, useSelector } from 'react-redux'
import { closeReport, openReport, createQuery, saveMap, reportTitleChange, getTokens } from './actions'
import Query from './Query'
import { SaveOutlined, PlaySquareOutlined, EditOutlined } from '@ant-design/icons'
import debounce from 'lodash.debounce'
import { KeplerGlSchema } from 'kepler.gl/schemas'
import classnames from 'classnames'
import DekartMenu from './DekartMenu'
import { Header } from './Header'

function ReportQuery ({ reportId }) {
  const queries = useSelector(state => state.queries)
  const report = useSelector(state => state.report)
  const dispatch = useDispatch()
  useEffect(() => {
    if (report && !(queries && queries.length)) {
      dispatch(createQuery(reportId))
    }
  }, [reportId, report, queries, dispatch])
  if (queries && queries.length) {
    const queriesSections = queries.map(query => <Query query={query} key={query.id} />)
    return (
      <div className={styles.querySection}>{queriesSections}</div>
    )
  } else {
    return null
  }
}

const checkMapConfig = debounce((kepler, mapConfig, setMapChanged) => {
  if (kepler) {
    const configToSave = JSON.stringify(KeplerGlSchema.getConfigToSave(kepler))
    setMapChanged(configToSave !== mapConfig)
  }
}, 500)

function HeaderButtons ({ edit, changed, canSave, reportId }) {
  const dispatch = useDispatch()
  const history = useHistory()
  if (edit) {
    return (
      <div className={styles.headerButtons}>
        <Button
          type='primary'
          icon={<SaveOutlined />}
          disabled={!canSave}
          onClick={() => dispatch(saveMap())}
        >Save Map{changed ? '*' : ''}
        </Button>
        <Button
          icon={<PlaySquareOutlined />}
          disabled={changed}
          onClick={() => history.replace(`/reports/${reportId}`)}
        >Present
        </Button>
      </div>
    )
  }
  return (
    <div className={styles.headerButtons}>
      <Button
        type='primary'
        icon={<EditOutlined />}
        onClick={() => history.replace(`/reports/${reportId}/edit`)}
      >Edit
      </Button>
    </div>
  )
}

function Title () {
  const reportStatus = useSelector(state => state.reportStatus)
  const [edit, setEdit] = useState(false)
  const dispatch = useDispatch()
  if (reportStatus.edit && edit) {
    return (
      <div className={styles.title}>
        <Input
          className={styles.titleInput}
          value={reportStatus.title}
          onChange={(e) => dispatch(reportTitleChange(e.target.value))}
          onBlur={() => setEdit(false)}
          placeholder='Untitled'
          // size='large'
          disabled={!reportStatus.edit}
          // bordered={false}
        />
      </div>
    )
  } else {
    return (
      <div className={styles.title}>
        <span
          className={classnames(
            styles.titleText,
            reportStatus.edit && styles.titleTextEdit
          )}
          onClick={() => reportStatus.edit && setEdit(true)}
        >{reportStatus.title} <EditOutlined className={styles.titleEditIcon} />
        </span>
      </div>
    )
  }
}

function Kepler ({ mapboxApiAccessToken }) {
  // const dispatch = useDispatch()
  // useEffect(() => {
  //   dispatch(getTokens())
  // }, [dispatch])
  // const mapboxApiAccessToken = useSelector(state => state.tokens.mapbox)
  if (!mapboxApiAccessToken) {
    return (
      <div className={styles.keplerFlex}>
        <div className={styles.keplerBlock} />
      </div>
    )
  }
  return (
    <div className={styles.keplerFlex}>
      <div className={styles.keplerBlock}>
        <AutoSizer>
          {({ height, width }) => (
            <KeplerGl
              id='kepler'
              mapboxApiAccessToken={mapboxApiAccessToken}
              width={width}
              height={height}
            />
          )}
        </AutoSizer>
      </div>
    </div>

  )
}

export default function ReportPage ({ edit }) {
  const { id } = useParams()
  const history = useHistory()

  const kepler = useSelector(state => state.keplerGl.kepler)
  const report = useSelector(state => state.report)
  const { mapConfig, title } = report || {}
  const reportStatus = useSelector(state => state.reportStatus)
  const mapboxApiAccessToken = useSelector(state => state.tokens.mapbox)

  const dispatch = useDispatch()

  const [mapChanged, setMapChanged] = useState(false)

  useEffect(() => {
    if (mapboxApiAccessToken) {
      dispatch(openReport(id, edit, history))
      return () => dispatch(closeReport(id))
    } else {
      dispatch(getTokens())
    }
  }, [id, dispatch, edit, history, mapboxApiAccessToken])

  useEffect(() => checkMapConfig(kepler, mapConfig, setMapChanged), [kepler, mapConfig, setMapChanged])
  const titleChanged = reportStatus.title && title && reportStatus.title !== title

  if (!report) {
    return null
  }

  return (
    <div className={styles.report}>
      <Header>
        <Title />
        <DekartMenu />
        <HeaderButtons
          reportId={id}
          changed={mapChanged || titleChanged}
          canSave={reportStatus.canSave}
          edit={edit}
        />
      </Header>
      <div className={styles.body}>
        <Kepler mapboxApiAccessToken={mapboxApiAccessToken} />
        {edit ? <ReportQuery reportId={id} /> : null}
      </div>
    </div>
  )
}
