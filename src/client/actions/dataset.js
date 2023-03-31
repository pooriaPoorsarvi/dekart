import { CreateDatasetRequest, RemoveDatasetRequest } from '../../proto/dekart_pb'
import { Dekart } from '../../proto/dekart_pb_service'
import { unary } from '../lib/grpc'
import { downloading, error, finishDownloading, success } from './message'
import {
  addDataToMap,
  toggleSidePanel,
  removeDataset as keplerRemove,
  receiveMapConfig
} from 'kepler.gl/dist/actions'
import { processCsvData, processGeojson } from 'kepler.gl/dist/processors'
import { get } from '../lib/api'
import { KeplerGlSchema } from 'kepler.gl/dist/schemas'


export function createDataset (reportId) {
  return (dispatch) => {
    dispatch({ type: createDataset.name })
    const request = new CreateDatasetRequest()
    request.setReportId(reportId)
    unary(Dekart.CreateDataset, request).catch(err => dispatch(error(err)))
  }
}

export function setActiveDataset (datasetId) {
  return (dispatch, getState) => {
    const { datasets } = getState()
    const dataset = datasets.find(d => d.id === datasetId) || datasets[0]
    if (dataset) {
      dispatch({ type: setActiveDataset.name, dataset })
    }
  }
}

export function removeDataset (datasetId) {
  return async (dispatch, getState) => {
    const { datasets, activeDataset } = getState()
    if (activeDataset.id === datasetId) {
      // removed active query
      const datasetsLeft = datasets.filter(q => q.id !== datasetId)
      if (datasetsLeft.length === 0) {
        dispatch(error(new Error('Cannot remove last dataset')))
        return
      }
      dispatch(setActiveDataset(datasetsLeft.id))
    }
    dispatch({ type: removeDataset.name, datasetId })

    const request = new RemoveDatasetRequest()
    request.setDatasetId(datasetId)
    try {
      await unary(Dekart.RemoveDataset, request)
      dispatch(success('Dataset removed'))
    } catch (err) {
      dispatch(error(err))
    }
  }
}

export function downloadDataset (dataset, sourceId, extension, label) {
  return async (dispatch, getState) => {
    dispatch({ type: downloadDataset.name, dataset })
    dispatch(downloading(dataset))
    let data
    try {
      const res = await get(`/dataset-source/${sourceId}.${extension}`)
      if (extension === 'csv') {
        const csv = await res.text()
        data = processCsvData(csv)
      } else {
        const json = await res.json()
        data = processGeojson(json)
      }
    } catch (err) {
      dispatch(error(err))
      return
    }
    const { datasets } = getState()
    const i = datasets.findIndex(d => d.id === dataset.id)
    if (i < 0) {
      return
    }
    try {

      // TODO : remove the following retrieval of current stage on the fixing of kepler bug
      const { keplerGl } = getState()
      const visState = keplerGl['kepler'].visState;
      const datasetLayer = visState.layers.find(layer => layer.config.dataId === dataset.id);


      dispatch(keplerRemove(dataset.id))


      // TODO : remove the following retrieval of current stage on the fixing of kepler bug
      if (datasetLayer){
        // Changes required for keeping the in memory changes of kepler
        let cK = localStorage.getItem("keplerConfig");
        if( cK && cK !="undefined" ){
          const parsedConfig = KeplerGlSchema.parseSavedConfig(JSON.parse(cK));
          // console.log("configs are:");
          // console.log(parsedConfig.visState);
          let new_config =  KeplerGlSchema.parseSavedConfig(JSON.parse(JSON.stringify(KeplerGlSchema.getConfigToSave(keplerGl.kepler))));
          // console.log(new_config);
          parsedConfig.visState.layers = parsedConfig.visState.layers.filter(layer => layer.config.dataId == dataset.id);
          dispatch(receiveMapConfig(parsedConfig, {keepExistingConfig: true}));
        }
      }


      await dispatch(addDataToMap({
        datasets: {
          info: {
            label,
            id: dataset.id
          },
          data
        }
      }))
    } catch (err) {
      console.log(err)
      dispatch(error(
        new Error(`Failed to add data to map: ${err.message}`),
        false
      ))
      return
    }
    dispatch(finishDownloading(dataset))
    const { reportStatus } = getState()
    if (reportStatus.edit) {
      dispatch(toggleSidePanel('layer'))
    }
  }
}
