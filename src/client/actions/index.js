import { showDatasetTable, toggleModal } from 'kepler.gl/dist/actions'

export function showDataTable (datasetId) {
  return (dispatch) => {
    dispatch(showDatasetTable(datasetId))
    dispatch(toggleModal('dataTable'))
  }
}

export * from './query'
export * from './file'
export * from './report'
export * from './env'
export * from './message'
export * from './clipboard'
export * from './version'
export * from './dataset'
